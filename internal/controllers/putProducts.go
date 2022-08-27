package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/xvbnm48/course-gin/internal"
)

type putProducts struct {
	Name        string  `json:"name" binding:"required_without_all=Price Description"`
	Price       float64 `json:"price" binding:"omitempty,gt=0"`
	Description string  `json:"description" binding:"omitempty,max=250"`
}

func PutProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var binding guidBinding
		var payload putProducts
		var ctx = c.Request.Context()

		if e := c.ShouldBindUri(&binding); e != nil {
			var res = internal.NewHttpResponse(http.StatusBadRequest, e)
			c.JSON(res.Status, res)
			return
		}

		if e := c.ShouldBindJSON(&payload); e != nil {
			var res = internal.NewHttpResponse(http.StatusBadRequest, e)
			c.JSON(res.Status, res)
			return
		}

		var row = db.QueryRowContext(ctx, "SELECT name, price, description FROM products WHERE guid=?", binding.GUID)
		var currentProduct Product

		if e := row.Scan(&currentProduct.Name, &currentProduct.Price, &currentProduct.Description); e != nil {
			if e == sql.ErrNoRows {
				var res = internal.NewHttpResponse(http.StatusNotFound, e)
				c.JSON(res.Status, res)
				return
			}

			var res = internal.NewHttpResponse(http.StatusInternalServerError, e)
			c.JSON(res.Status, res)
		}

		var option = copier.Option{
			IgnoreEmpty: true,
			DeepCopy:    true,
		}

		if e := copier.CopyWithOption(&currentProduct, &payload, option); e != nil {
			var res = internal.NewHttpResponse(http.StatusInternalServerError, e)
			c.JSON(res.Status, res)
			return
		}
		if _, e := db.ExecContext(ctx, "UPDATE products SET name=?, price=?, description=? WHERE guid=?", currentProduct.Name, currentProduct.Price, currentProduct.Description, binding.GUID); e != nil {
			var res = internal.NewHttpResponse(http.StatusInternalServerError, e)
			c.JSON(res.Status, res)
			return
		}

		var updateRow = db.QueryRowContext(ctx, "SELECT guid, name, price, description, createdAt FROM products WHERE guid=?", binding.GUID)
		var product Product
		if e := updateRow.Scan(&product.GUID, &product.Name, &product.Price, &product.Description, &product.CreatedAt); e != nil {
			var res = internal.NewHttpResponse(http.StatusInternalServerError, e)
			c.JSON(res.Status, res)
			return
		}

		var res = internal.NewHttpResponse(http.StatusOK, product)
		c.JSON(res.Status, res)
	}
}
