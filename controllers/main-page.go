package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"github.com/user/sujag/models"
	"net/http"
)

func (c *Controllers) MainPagePosts(ctx *gin.Context) {

	selectFields := bson.M{"title": 1, "thumbnail": 1, "excerpt": 1, "slug": 1}

	var videos []models.Post
	var features []models.Post
	var nuktanazars []models.Post

	c.App.DB.C("posts").Find(bson.M{"type": "feature", "category": "election"}).Limit(11).Select(selectFields).Sort("-published_on").All(&features)

	// List of categories we need
	categories := [5]string{"nuktanazar", "nuktanazar", "baylag", "baylag", "terrorism-1"}

	NuktanazarFetched := false
	BayLagFetched := false
	for _, category := range categories {
		var post models.Post
		offset := 0
		if category == "nuktanazar" {
			if NuktanazarFetched {
				offset = 1
			}
			NuktanazarFetched = true
		} else if category == "baylag" {
			if BayLagFetched {
				offset = 1
			}
			BayLagFetched = true
		}
		c.App.DB.C("posts").Find(bson.M{"type": "nuktanazar", "category": category}).Skip(offset).Sort("-published_on").Select(selectFields).One(&post)
		if post.Title != "" {
			nuktanazars = append(nuktanazars, post)
		}
	}

	// get Videos
	c.App.DB.C("posts").Find(bson.M{"type": "video"}).Limit(7).Select(bson.M{"title": 1, "thumbnail": 1, "excerpt": 1, "content": 1, "slug": 1}).Sort("-published_on").All(&videos)

	ctx.JSON(http.StatusOK, gin.H{
		"features":   features,
		"nuktanazar": nuktanazars,
		"videos":     videos,
	})

}

func (c *Controllers) NuktanazarUpdates(ctx *gin.Context) {

	selectFields := bson.M{"title": 1, "thumbnail": 1, "excerpt": 1, "slug": 1}
	limit := 10
	params := c.GetParams(ctx)
	Page := params["page"].(float64)
	Skip := limit * int(Page)
	var posts []models.Post

	c.App.DB.C("posts").Find(bson.M{"type": "nuktanazar", "category": "election-update"}).Skip(Skip).Limit(limit).Select(selectFields).Sort("-published_on").All(&posts)

	ctx.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

}
