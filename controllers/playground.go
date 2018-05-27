package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"gopkg.in/mgo.v2/bson"
)

func (c *Controllers) PlayGround(ctx *gin.Context) {
	ch := make(chan string)
	Test(50000, ch)
	Test(50000, ch)
	str1, str2 := <-ch, <-ch
	log.Println("completed", str1, str2)
	str := "RESPONSE = " + str1 + ":" + str2
	ctx.JSON(http.StatusOK, bson.M{
		"response": str,
	})
}

func Test(times int, ch chan string) {
	for v := 1; v < times; v++ {
		log.Println(v)
	}
	response := "PAKISTAN " + string(times)
	ch <- response
}
