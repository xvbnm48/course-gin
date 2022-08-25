package controllers

import "github.com/gin-gonic/gin"

func PostProduct(c *gin.Context) {
	c.String(200, "PostProduct")
}
