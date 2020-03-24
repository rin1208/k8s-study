package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Names struct {
	Name string
}

func main() {
	r := gin.Default()
	r.GET("/", huga)
	r.GET("/hoge", hoge)
	r.Run(":80")
}
func hoge(c *gin.Context) {
	client, err := gorm.Open("mysql", "root:root@tcp([mysql]:3306)/hoge")

	client.AutoMigrate(&Names{})

	if err != nil {
		log.Fatal(err.Error())
	}
	var hunga Names
	hunga.Name = "hoge"
	client.Create(&hunga)
	c.JSON(200, "hoge")
}

func huga(c *gin.Context) {
	client, err := gorm.Open("mysql", "root:root@tcp([mysql]:3306)/hoge")
	client.AutoMigrate(&Names{})
	if err != nil {
		log.Fatal(err.Error())
	}
	var hunga Names
	client.Find(&hunga)
	c.JSON(200, hunga)
}
