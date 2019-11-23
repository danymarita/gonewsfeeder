package main

import (
	"gonewsfeeder/httpd/handler"
	"gonewsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("Goooo..")
	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	r.Run(":5000")
}
