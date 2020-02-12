package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hoge", hoge)
	r.Run(":80")
}
func hoge(c *gin.Context) {
	c.JSON(200, "hgoe")
}
