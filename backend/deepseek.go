package main

import (
	"context"
	"sort"

	p "github.com/bestcb2333/gin-gorm-preloader/preloader"
	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type ExtractReq struct {
	Message string `json:"message"`
}

func RegDeepSeekhandler(r *gin.Engine, bc *p.BaseConfig, cfg *Config) {

	client := openai.NewClient(
		option.WithAPIKey(cfg.DeepSeekAPI),
		option.WithBaseURL("https://api.deepseek.com"),
	)

	r.POST("/deepseek", p.Preload(
		&p.Config[ExtractReq]{
			Base:       bc,
			Bind:       &p.BindConfig{JSON: true},
			Permission: p.Login,
			Tables:     []string{"Messages"},
		},
		func(c *gin.Context, u *User, r *ExtractReq) {

			var openaiMessages []openai.ChatCompletionMessageParamUnion
			openaiMessages = append(
				openaiMessages,
				openai.SystemMessage("用户可能会提问与城市内涝隐患相关的问题，你需要解答用户的问题。回答50字即可。"),
			)

			for _, message := range u.Messages {
				if message.Role == "user" {
					openaiMessages = append(openaiMessages, openai.UserMessage(message.Content))
				} else {
					openaiMessages = append(openaiMessages, openai.AssistantMessage(message.Content))
				}
			}

			openaiMessages = append(openaiMessages, openai.UserMessage(r.Message))

			completion, err := client.Chat.Completions.New(
				context.Background(),
				openai.ChatCompletionNewParams{
					Model:    "deepseek-chat",
					Messages: openaiMessages,
				},
			)
			if err != nil {
				c.JSON(500, Resp("请求失败", err, nil))
				return
			}

			result := completion.Choices[0].Message.Content

			if err := bc.DB.Create(&Message{
				UserID: u.ID,
				MessageDTO: MessageDTO{
					Role:    "user",
					Content: r.Message,
				},
			}).Error; err != nil {
				c.JSON(500, bc.Resp("存储用户消息到数据库失败", err, nil))
				return
			}

			if err := bc.DB.Create(&Message{
				UserID: u.ID,
				MessageDTO: MessageDTO{
					Role:    "assistant",
					Content: result,
				},
			}).Error; err != nil {
				c.JSON(500, bc.Resp("存储结果到数据库失败", err, nil))
				return
			}

			var messages []Message
			if err := bc.DB.Order("created_at desc").Limit(10).Find(&messages).Error; err != nil {
				c.JSON(500, bc.Resp("获取消息列表失败", err, nil))
				return
			}

			sort.Slice(messages, func(i, j int) bool {
				return messages[i].ID < messages[j].ID
			})

			c.JSON(200, bc.Resp("请求成功", nil, messages))
		},
	))

	r.GET("/deepseek", p.Preload(
		&p.Config[struct{}]{
			Base:       bc,
			Permission: p.Login,
			Tables:     []string{"Messages"},
		},
		func(c *gin.Context, u *User, r *struct{}) {

			sort.Slice(u.Messages, func(i, j int) bool {
				return u.Messages[i].ID < u.Messages[j].ID
			})

			c.JSON(200, bc.Resp("获取历史消息成功", nil, u.Messages))
		},
	))

	r.DELETE("/deepseek", p.Preload(
		&p.Config[struct{}]{
			Base:       bc,
			Permission: p.Login,
		},
		func(c *gin.Context, u *User, r *struct{}) {

			if err := bc.DB.Where("user_id = ?", u.ID).Delete(new(Message)).Error; err != nil {
				c.JSON(500, bc.Resp("删除失败", err, nil))
				return
			}

			c.JSON(200, bc.Resp("删除成功", nil, nil))
		},
	))

}
