package controllers

import "github.com/gin-gonic/gin"

func DeleteProduct(c *gin.Context) {
	c.String(200, "DeleteProduct")
}
