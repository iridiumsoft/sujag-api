package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"github.com/user/sujag/models"
)

func (c *Controllers) getMainPosts(context *gin.Context) {
	var features []models.Post
	var nuktanazars []models.Post
	c.App.DB.C("posts").Find(bson.M{"type": "feature",}).Limit(11).Select(bson.M{"title": 1, "thumbnail": 1, "excerpt": 1, "slug": 1}).All(&features)
	var categories [5]string
	categories[0] = "nuktanazar"
	categories[1] = "nuktanazar"
	categories[2] = "baylag"
	categories[3] = "baylag"
	categories[4] = "election-diary"
	for _, category := range categories {
		var post models.Post
		c.App.DB.C("posts").Find(bson.M{"type": "nuktanazar", "category": category}).Select(bson.M{"title": 1, "thumbnail": 1, "excerpt": 1, "slug": 1}).One(&post)
		if post.Title != "" {
			nuktanazars = append(nuktanazars, post)
		}
	}
	context.JSON(200, gin.H{
		"features":    features,
		"nuktanazars": nuktanazars,
	})
}
