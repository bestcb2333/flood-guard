package main

import "github.com/gin-gonic/gin"

func AddNotice(c *gin.Context, u *User, r struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}) {

	notice := Notice{
		UserID:  &u.ID,
		Title:   r.Title,
		Content: r.Content,
	}

	if err := DB.Create(&notice).Error; err != nil {
		c.JSON(500, Resp("公告添加失败", err, nil))
		return
	}

	c.JSON(200, Resp("公告添加成功", nil, &notice))
}
