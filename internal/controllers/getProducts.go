package controllers

import "github.com/gin-gonic/gin"

func GetProducts(c *gin.Context) {
	c.String(200, "GetProducts")
}
