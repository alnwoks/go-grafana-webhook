package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.SetHTMLTemplate(html)
	r.Static("/assets", "./assets")

	r.GET("/", Default())

	r.POST("/grafana", PostBody())

	// Listen and Server in https://127.0.0.1:8080
	// r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
	r.Run(":8080")
}
