package util

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"fmt"
)

func GetParams(ctx *gin.Context) map[string]interface{} {
	params := ctx.Query("params")
	var model interface{}
	if params != "" {
		in := []byte(params)
		json.Unmarshal(in, &model)
	}
	param, _ := model.(map[string]interface{})
	return param
}

// its like p_rr function to print in json format
func Print(data interface{}) {
	jsonData, _ := json.Marshal(data)
	fmt.Printf("%s\n", jsonData)
}
