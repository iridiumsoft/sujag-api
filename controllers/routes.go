package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controllers) Routing() error {

	// Welcome page
	c.Gin.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	// Home Page
	HomePage := c.Gin.Group("/home")
	HomePage.GET("/posts", c.MainPagePosts)
	HomePage.GET("/updates", c.MainPageNuktanazarUpdates)
	HomePage.GET("/election", c.MainPageElectionPosts)
	HomePage.GET("/mobile-posts", c.MainPageMobilePosts)

	// Archive
	c.Gin.GET("/archive", c.getArchivePosts)

	// Single Post
	c.Gin.GET("/single/:slug", c.getSinglePost)

	return c.Gin.Run(":" + c.Config.Port)
}
