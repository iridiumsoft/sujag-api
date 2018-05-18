package controllers

// Routing initialize endpoints and launch http server.
func (c *Controllers) Routing() error {

	c.Gin.GET("/main/posts", c.getMainPosts)

	return c.Gin.Run(":" + c.Config.Port)
}
