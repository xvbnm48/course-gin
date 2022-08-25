package controllers

import "github.com/gin-gonic/gin"

func GetProduct(c *gin.Context) {
	c.String(200, "GetProduct")
}
