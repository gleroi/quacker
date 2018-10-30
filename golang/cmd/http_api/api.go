package main

import (
	"quacker"
	"quacker/memory"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	root := gin.Default()

	store := memory.NewStore()

	api := root.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		api.POST("/quack", func(c *gin.Context) {
			var cmd QuackMessage
			err := c.BindJSON(&cmd)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, err)
				return
			}
			tr := memory.NewTransaction(store)
			defer tr.Commit()

			quacker.Quack(tr, cmd.AuthorID, cmd.Content)

			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
			})
		})

		api.GET("/timeline/:author", func(c *gin.Context) {
			authorID := c.Param("author")
			tl := quacker.GetTimeline(store, quacker.UserID(authorID))

			c.JSON(http.StatusOK, tl)
		})

	}

	root.StaticFile("/", "www/index.html")

	root.Run()
}
