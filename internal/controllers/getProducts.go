package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/course-gin/internal"
)

func GetProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var rows *sql.Rows
		var e error
		if rows, e = db.Query("SELECT guid,name,price,description,createdAt FROM products"); e != nil {
			var res = internal.NewHttpResponse(500, e)
			c.JSON(500, res)
			return
		}
		defer rows.Close()
		var products []Product
		for rows.Next() {
			var product Product

			if e := rows.Scan(&product.GUID, &product.Name, &product.Price, &product.Description, &product.CreatedAt); e != nil {
				var res = internal.NewHttpResponse(500, e)
				c.JSON(500, res)
				return
			}

			products = append(products, product)
		}

		if len(products) == 0 {
			var res = internal.NewHttpResponse(404, sql.ErrNoRows)
			c.JSON(404, res)
			return
		}

		var res = internal.NewHttpResponse(200, products)
		c.JSON(200, res)
	}
}
