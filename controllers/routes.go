package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controllers) Routing() error {

	// Welcome page
	c.Gin.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	// Main Page
	HomePage := c.Gin.Group("/home")
	HomePage.GET("/posts", c.MainPagePosts)
	HomePage.GET("/updates", c.MainPageNuktanazarUpdates)
	HomePage.GET("/election", c.MainPageElectionPosts)
	HomePage.GET("/mobile-posts", c.MainPageMobilePosts)

	// Archive
	Archive := c.Gin.Group("/archive")
	Archive.GET("/", c.getArchivePosts)

	return c.Gin.Run(":" + c.Config.Port)
}
