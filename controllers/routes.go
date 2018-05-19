package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func (c *Controllers) Routing() error {
	c.Gin.GET("/main/posts", c.MainPagePosts)
	c.Gin.GET("/main/updates", c.NuktanazarUpdates)
	return c.Gin.Run(":" + c.Config.Port)
}

func (c *Controllers) GetParams(ctx *gin.Context) map[string]interface{} {
	params := ctx.Query("params")
	var model interface{}
	if params != "" {
		in := []byte(params)
		json.Unmarshal(in, &model)
	}
	param, _ := model.(map[string]interface{})
	return param
}
