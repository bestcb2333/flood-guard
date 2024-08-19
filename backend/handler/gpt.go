package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
	ai "github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

type GptMessage struct {
	MessageID string
	Role      string
	Time      string
	Content   string
}

var cli *ai.Client

func Gpt(c *gin.Context) {

	switch c.Request.Method {
	case http.MethodPost:
		var request struct {
			ThreadID string
			Message  string
		}
		if err := util.ParseJSON(c, &request); err != nil {
			util.Error(c, 400, "你的请求体格式不正确", err)
			return
		}

		// 从JWT里获取用户ID
		user, err := VerifyJwt(c, "Thread")
		if err != nil {
			util.Error(c, 400, "无法读取你的用户身份", err)
			return
		}

		asst := ai.RunRequest{AssistantID: viper.GetString("GPT_ASST_ID")}

		if request.ThreadID == "" { // 创建新的会话

			ctx, canc := context.WithTimeout(context.Background(), 5*time.Second)
			defer canc()

			// 创建会话
			thread, err := cli.CreateThread(ctx, ai.ThreadRequest{
				Messages: []ai.ThreadMessage{{
					Role:    ai.ThreadMessageRole("user"),
					Content: request.Message,
				}}})
			if err != nil {
				util.Error(c, 500, "会话创建失败", err)
				return
			}

			request.ThreadID = thread.ID

			if err := DB.Create(&database.Thread{
				ThreadID:   thread.ID,
				ThreadName: time.Now().Format("2006-01-02 15:04:05"),
				UserID:     user.ID,
			}).Error; err != nil {
				util.HandleQueryErr(c, "无法将你的会话信息存储到数据库", err)
				return
			}

		} else { // 使用已有的会话

			if err := DB.First(
				&database.Thread{},
				"user_id = ? AND thread_id = ?",
				user.ID, request.ThreadID,
			).Error; err != nil {
				util.Error(c, 400, "你没有这个会话", err)
				return
			}

			ctx, canc := context.WithTimeout(
				context.Background(),
				5*time.Second,
			)
			defer canc()

			if _, err := cli.CreateMessage(
				ctx, request.ThreadID,
				ai.MessageRequest{
					Role:    "user",
					Content: request.Message,
				},
			); err != nil {
				util.Error(c, 500, "无法将你的消息添加到会话", err)
				return
			}

		}

		ctx, canc := context.WithTimeout(
			context.Background(),
			5*time.Second,
		)
		defer canc()

		run, err := cli.CreateRun(ctx, request.ThreadID, asst)
		if err != nil {
			util.Error(c, 500, "无法生成回答", err)
			return
		}

		mes, err := PollRunStatus(cli, request.ThreadID, run.ID)
		if err != nil {
			util.Error(c, 500, "获取消息失败", err)
			return
		}

		util.Info(c, 200, "执行成功", gin.H{
			"userId":   user.ID,
			"username": user.Username,
			"threadId": request.ThreadID,
			"message":  mes,
		})

	case http.MethodGet:
		threadID := c.Query("session_id")
		threadName := c.Query("session_name")
		action := c.Query("action")

		user, err := VerifyJwt(c, "Thread")
		if err != nil {
			util.Error(c, 500, "读取用户JWT信息失败", err)
			return
		}

		if threadID == "" { // 如果会话ID为空，就返回这个用户的所有会话

			var results []map[string]any
			for _, thread := range user.Thread {
				results = append(results, map[string]any{
					"session_id":   thread.ThreadID,
					"session_name": thread.ThreadName,
					"time":         thread.UpdatedAt.In(util.Loc).Format("2006/01/02 15:04"),
				})
			}
			util.Info(c, 200, "会话信息查找完成", results)

		} else { // 如果会话ID不为空，进行删除会话或修改会话名或列出消息

			var tmp database.Thread
			if err = DB.First(
				&tmp, "thread_id = ? AND user_id = ?", threadID, user.ID,
			).Error; err != nil {
				util.HandleQueryErr(c, "你没有这个会话", err)
				return
			}

			if action == "delete" { // 如果行为等于删除，删除会话
				if err := DB.Delete(
					&tmp, "thread_id = ?", threadID,
				).Error; err != nil {
					util.Error(c, 500, "无法删除这个会话", err)
					return
				}
				util.Info(c, 200, "会话删除成功", nil)

			} else if threadName != "" { // 如果会话名称不为空，修改会话名称
				if err := DB.Model(&tmp).Update(
					"thread_name", threadName,
				).Error; err != nil {
					util.Error(c, 500, "无法修改这个会话的名称", err)
					return
				}
				util.Info(c, 200, "会话名称修改成功", nil)

			} else { // 否则，查看会话内容
				messages, err := FetchMessages(cli, threadID)
				if err != nil {
					util.Error(c, 500, "无法获取会话消息", err)
					return
				}
				util.Info(c, 200, "查询成功", messages)
			}
		}
	}
}

// 对run对象进行轮询检测判断执行状态
func PollRunStatus(cli *ai.Client, threadID, runID string) (mes []GptMessage, err error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			run, err := cli.RetrieveRun(context.Background(), threadID, runID)
			if err != nil {
				return nil, fmt.Errorf("failed to retrieve run: %w", err)
			}

			switch run.Status {
			case ai.RunStatusCompleted:
				return FetchMessages(cli, threadID)
			case ai.RunStatusQueued, ai.RunStatusInProgress:
				continue
			default:
				return nil, fmt.Errorf("run failed with status: %s", run.Status)
			}
		}
	}
}

func FetchMessages(cli *ai.Client, threadID string) (destination []GptMessage, err error) {

	// 创建一个等待会话超时的上下文
	ctx, canc := context.WithTimeout(
		context.Background(),
		30*time.Minute,
	)
	defer canc()

	messages, err := cli.ListMessage(ctx, threadID, nil, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("无法列出消息：%w", err)
	}
	for _, value := range messages.Messages {
		destination = append(destination, GptMessage{
			MessageID: value.ID,
			Role:      value.Role,
			Time:      time.Now().Format("2006-01-02 15:04:05"),
			Content:   value.Content[0].Text.Value,
		})
	}
	return
}
