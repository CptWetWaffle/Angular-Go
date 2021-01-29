package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type location struct {
	Id int `json:"id"`
	Coordinate map[string]string `json:"title"`
	Address  map[string]string `json:"address"`
	UserId  string `json:"userId"`
	Type  string `json:"type"`
	AppKey  string `json:"appKey"`
	Time  time.Time `json:"time"`
	CreatedAt time.Time `json:"createdAt"`
}

func LocationGet(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := db.Exec("SELECT * FROM location"); err != nil {
			c.JSON(http.StatusOK, map[string]string{
				"hello": "found get",
			})
		}
	}
}

func LocationPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := location{}
		_ = c.Bind(&requestBody)

		c.JSON(
			http.StatusCreated,
			location{
			Id: 0,
			Coordinate: map[string]string{ "latitude": "test" },
			Address: map[string]string{ "address": "test" },
			UserId: "54643564355",
			Type: "myapp",
			AppKey: "v1@3423423434",
			Time: time.Now(),
			CreatedAt: time.Now(),
		},
		)
	}
}

func LocationDelete(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "found delete",
		})
	}
}
