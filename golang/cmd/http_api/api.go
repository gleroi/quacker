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
				return
			}
			tr := memory.NewTransaction(store)
			defer tr.Commit()

			quacker.Quack(tr, cmd.AuthorID, cmd.Content)

			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
			})
		})

		api.POST("/follow", func(c *gin.Context) {
			var cmd FollowUser
			err := c.BindJSON(&cmd)
			if err != nil {
				return
			}
			tr := memory.NewTransaction(store)
			defer tr.Commit()

			quacker.Follow(tr, cmd.Follower, cmd.Followee)

			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
			})
		})

		api.POST("/unfollow", func(c *gin.Context) {

		})

		api.GET("/timeline/:user", func(c *gin.Context) {
			userID := quacker.UserID(c.Param("user"))
			followees := quacker.GetFolloweeList(store, userID)
			tl := quacker.GetTimeline(store, userID, followees)

			c.JSON(http.StatusOK, tl)
		})

	}

	root.StaticFile("/", "www/index.html")

	root.Run()
}
