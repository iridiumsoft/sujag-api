package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/user/sujag/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"log"
)

func (c *Controllers) getSinglePost(ctx *gin.Context) {

	var Post models.Post
	var Author models.Author

	Slug := ctx.Param("slug")

	// TODO:: its not including published on, liked etc
	err := c.App.DB.C("posts").Find(bson.M{"slug": Slug}).Select(bson.M{
		"title":        1,
		"content":      1,
		"thumbnail":    1,
		"type":         1,
		"published_on": 1,
		"district":     1,
		"liked":        1,
		"author":       1,
	}).One(&Post)

	if err != nil {
		log.Println("Error while getting single post", err.Error())
	}

	// TODO:: Author is not sending name, because there is urdu name
	err = c.App.DB.C("authors").Find(bson.M{"username": Post.Author}).Select(bson.M{"info": 1, "name": 1, "username": 1, "dp_lg": 1}).One(&Author)

	if err != nil {
		log.Println("Error while getting author ", err.Error())
	}

	ctx.JSON(http.StatusOK, bson.M{
		"post":   Post,
		"author": Author,
	})

}
