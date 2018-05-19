package controllers

import (
	"encoding/json"
)

func (c *Controllers) Routing() error {
	c.Gin.GET("/main/posts", c.MainPagePosts)
	c.Gin.GET("/main/updates", c.NuktanazarUpdates)
	return c.Gin.Run(":" + c.Config.Port)
}

func (c *Controllers) GetParams(params string) map[string]interface{} {
	var model interface{}
	if params != "" {
		in := []byte(params)
		json.Unmarshal(in, &model)
	}
	param, _ := model.(map[string]interface{})
	return param
}
