package controllers

import (
	"fmt"
	"strings"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	// Local Files
	"github.com/user/sujag/app"
	"github.com/user/sujag/conf"
	"encoding/json"
)

// Controller implemented router of project
type Controllers struct {
	Config conf.Main
	App    *app.App
	Gin    *gin.Engine
}

func maxRequestsAtOnce(n int) gin.HandlerFunc {
	s := make(chan struct{}, n)
	return func(c *gin.Context) {
		s <- struct{}{}
		defer func() { <-s }()
		c.Next()
	}
}

var allowedMethods = "GET, PUT, POST, DELETE"

func initGin() *gin.Engine {
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(maxRequestsAtOnce(50))
	g.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        allowedMethods,
		RequestHeaders: "token, Origin, Authorization, Content-Type, Access-Control-Allow-Origin",
		ExposedHeaders: "",
		MaxAge:         50 * time.Second,
	}))
	return g
}

func security(ctx *gin.Context) {
	if !strings.Contains(allowedMethods, ctx.Request.Method) {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Method %s is not allowed", ctx.Request.Method),
		})
		ctx.Abort()
	}
	ctx.Next()
}

func New(config conf.Main, application *app.App) *Controllers {
	return &Controllers{
		Config: config,
		Gin:    initGin(),
		App:    application,
	}
}

func (c *Controllers) Print(data interface{}) {
	json_data, _ := json.Marshal(data)
	fmt.Printf("%s\n", json_data)
	return;
}
