package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	address := ":3000"
	v1 := router.Group("/api/v1")

	v1.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	v2 := router.Group("/api/v2")
	v2.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "get")
	})

	log.Fatal(router.Run(address))

}
