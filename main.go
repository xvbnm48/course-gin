package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	address := ":3000"

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	log.Fatal(router.Run(address))

}
