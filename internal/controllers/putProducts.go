package controllers

import "github.com/gin-gonic/gin"

func PutProducts(c *gin.Context) {
	c.String(200, "PutProducts")
}
