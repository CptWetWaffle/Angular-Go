package handler

import (
	"Angular-Go/api/server/platform/newsfeed"
	"github.com/gin-gonic/gin"
	"net/http"
)

type newsfeedPost struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsfeedPost{}
		_ = c.Bind(&requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}

		feed.Add(item)
		c.Status(http.StatusOK)
	}
}
