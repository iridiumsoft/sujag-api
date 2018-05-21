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
		"title":        1,
		"content":      1,
		"thumbnail":    1,
		"type":         1,
		"published_on": 1,
		"district":     1,
		"liked":        1,
		"author":       1,
	}).One(&Post)

	c.App.DB.C("authors").Find(bson.M{"username": Post.Author}).Select(bson.M{"info": 1, "name": 1, "username": 1, "dp_lg": 1}).One(&Author)

	ctx.JSON(http.StatusOK, bson.M{
		"post":   Post,
		"author": Author,
	})

}
