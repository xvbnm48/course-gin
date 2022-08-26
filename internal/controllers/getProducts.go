package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func GetProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "GetProducts")
	}
}
