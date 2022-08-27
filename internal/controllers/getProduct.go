package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/course-gin/internal"
)

type guidBinding struct {
	GUID string `uri:"guid" binding:"required,uuid4"`
}

func GetProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var binding guidBinding
		var ctx = c.Request.Context()
		if e := c.ShouldBindUri(&binding); e != nil {
			res := internal.NewHttpResponse(http.StatusBadRequest, e)
			c.JSON(res.Status, res)
			return
		}

		row := db.QueryRowContext(ctx, "SELECT guid,name,price,description,createdAt FROM products WHERE guid=?", binding.GUID)
		var product Product
		if e := row.Scan(&product.GUID, &product.Name, &product.Price, &product.Description, &product.CreatedAt); e != nil {
			if e == sql.ErrNoRows {
				res := internal.NewHttpResponse(http.StatusNotFound, e)
				c.JSON(res.Status, res)
				return
			}

			res := internal.NewHttpResponse(http.StatusInternalServerError, e)
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		var res = internal.NewHttpResponse(http.StatusOK, product)
		c.JSON(http.StatusOK, res)

	}
}
