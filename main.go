package main

import (
	"fmt"
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

	router.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "The id is %s", id)
	})

	router.POST("/path", func(c *gin.Context) {
		id := c.DefaultQuery("id", "11")
		fmt.Println("id:", id)
		c.String(http.StatusOK, "The id is %s", id)
	})

	router.POST("/products/:id", func(c *gin.Context) {
		id := c.GetHeader("id")
		fmt.Println("id:", id)
		c.String(http.StatusOK, "The id is %s", id)
	})
	log.Fatal(router.Run(address))

}
