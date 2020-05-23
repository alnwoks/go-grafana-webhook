package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Default displays the static template
func Default() gin.HandlerFunc {
	return func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// use pusher.Push() to do server push
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, schema, gin.H{
			"status": "success",
		})
	}
}

//PostBody is used to handle the body of a post action
func PostBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body GrafanaBody
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
		} else {
			res := AddNew(&body)
			c.JSON(http.StatusAccepted, gin.H{
				"id": res,
			})

		}
	}
}

//DeleteBody is used to delete all documents in the collection
func DeleteBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		Delete()
		c.JSON(http.StatusOK, gin.H{
			"status": "deleted",
		})

	}
}
