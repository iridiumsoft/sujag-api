package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"github.com/user/sujag/models"
	"net/http"
	"github.com/user/sujag/util"
)

func (c *Controllers) MainPagePosts(ctx *gin.Context) {

	selectFields := bson.M{"title": 1, "thumbnail": 1, "excerpt": 1, "slug": 1, "published_on": 1}

	var videos []models.Post
	var features []models.Post
	var nuktanazars []models.Post
	var nuktanazarsTops []models.Post
	var post models.Post

	Select := bson.M{"title": 1, "thumbnail": 1, "excerpt": 1, "slug": 1, "district": 1, "published_on": 1}
	// TODO:: Get published on date, its not working for some reasons
	c.App.DB.C("posts").Find(bson.M{"type": "feature", "category": "election", "status": 1}).Limit(9).Select(Select).Sort("-published_on").All(&features)

	// List of categories we need
	categories := [5]string{"nuktanazar", "nuktanazar", "baylag", "baylag", "terrorism-1"}

	NuktanazarFetched := false
	BayLagFetched := false
	var NuktanazarFetchedSlugs []string

	// Get BayLag
	c.App.DB.C("posts").Find(bson.M{"type": "nuktanazar", "category": "election-update"}).Select(Select).Sort("-published_on").One(&post)
	nuktanazarsTops = append(nuktanazarsTops, post)
	NuktanazarFetchedSlugs = append(NuktanazarFetchedSlugs, post.Slug)

	// Get Nuktanazar
	c.App.DB.C("posts").Find(bson.M{"type": "nuktanazar", "category": "nuktanazar"}).Select(Select).Sort("-published_on").One(&post)
	nuktanazarsTops = append(nuktanazarsTops, post)
	NuktanazarFetchedSlugs = append(NuktanazarFetchedSlugs, post.Slug)

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

		c.App.DB.C("posts").Find(bson.M{"type": "nuktanazar", "category": category, "status": 1, "slug": bson.M{"$nin": NuktanazarFetchedSlugs}}).Skip(offset).Sort("-published_on").Select(selectFields).One(&post)

		if post.Title != "" {
			nuktanazars = append(nuktanazars, post)
		}

	}

	// get Videos
	c.App.DB.C("posts").Find(bson.M{"type": "video", "status": 1}).Limit(7).Select(bson.M{"title": 1, "thumbnail": 1, "excerpt": 1, "content": 1, "slug": 1}).Sort("-published_on").All(&videos)

	ctx.JSON(http.StatusOK, gin.H{
		"features":       features,
		"nuktanazar":     nuktanazars,
		"nuktanazar-top": nuktanazarsTops,
		"videos":         videos,
	})

}

func (c *Controllers) MainPageNuktanazarUpdates(ctx *gin.Context) {

	selectFields := bson.M{"title": 1, "thumbnail": 1, "excerpt": 1, "slug": 1}
	limit := 10
	params := util.GetParams(ctx)
	Page := params["page"].(float64)
	Skip := limit * int(Page)
	Skip = Skip + 1
	var posts []models.Post

	c.App.DB.C("posts").Find(bson.M{"type": "nuktanazar", "category": "election-update", "status": 1}).Skip(Skip).Limit(limit).Select(selectFields).Sort("-published_on").All(&posts)

	ctx.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

}

func (c *Controllers) MainFooterContent(ctx *gin.Context) {

	var kahani models.Kahani
	var sports models.Post
	var magazine models.Magazine
	c.App.DB.C("kahani").Find(bson.M{}).Select(bson.M{"name": 1, "thumbnail": 1, "audio": 1, "title": 1, "excerpt": 1}).Sort("-created_on").One(&kahani)
	c.App.DB.C("posts").Find(bson.M{"type": "sports", "status": 1}).Select(bson.M{"title": 1, "thumbnail": 1, "excerpt": 1, "slug": 1}).Sort("-published_on").One(&sports)
	c.App.DB.C("magazine").Find(bson.M{}).Select(bson.M{"title": 1, "is": 1, "img": 1, "created_on": 1}).Sort("-created_on").One(&magazine)

	ctx.JSON(http.StatusOK, gin.H{
		"kahani":   kahani,
		"sports":   sports,
		"magazine": magazine,
	})
}

func (c *Controllers) MainPageElectionPosts(ctx *gin.Context) {

}

func (c *Controllers) MainPageMobilePosts(ctx *gin.Context) {

}
