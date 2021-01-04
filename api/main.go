package main

import (
	"SimpleApp/api/handler"
	"SimpleApp/api/platform/newsfeed"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}

	r.Use(cors.New(config))
	api := r.Group("/api")
	{
		api.GET("/ping", handler.PingGet())
		api.GET("/newsfeed", handler.NewsfeedGet(feed))
		api.POST("/newsfeed", handler.NewsfeedPost(feed))
		api.GET("/loc", handler.LocationPost())
	}

	_ = r.Run(":3000")
}
