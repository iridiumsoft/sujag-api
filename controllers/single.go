package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/user/sujag/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func (c *Controllers) getSinglePost(ctx *gin.Context) {

	var Post models.Post
	var Author models.Author

	Slug := ctx.Param("slug")

	c.App.DB.C("posts").Find(bson.M{"slug": Slug}).Select(bson.M{
		"title":        true,
		"content":      true,
		"thumbnail":    true,
		"type":         true,
		"published_on": true,
		"author":       true,
	}).One(&Post)

	c.App.DB.C("authors").Find(bson.M{"username": Post.Author}).Select(bson.M{"info": true, "name": true, "username": true}).One(&Author)

	ctx.JSON(http.StatusOK, bson.M{
		"post":   Post,
		"author": Author,
	})

}
