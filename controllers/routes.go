package controllers

func (c *Controllers) Routing() error {

	// Welcome page
	//c.Gin.GET("/", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "index.html", gin.H{
	//		"title": "Main website",
	//	})
	//})

	// Main Page
	HomePage := c.Gin.Group("/main")
	HomePage.GET("/posts", c.MainPagePosts)
	HomePage.GET("/updates", c.MainPageNuktanazarUpdates)
	HomePage.GET("/election", c.MainPageElectionPosts)
	HomePage.GET("/mobile-posts", c.MainPageMobilePosts)

	return c.Gin.Run(":" + c.Config.Port)
}
