package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/course-gin/internal"
)

func DeleteProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var binding guidBinding
		var ctx = c.Request.Context()

		if e := c.ShouldBindUri(&binding); e != nil {
			res := internal.NewHttpResponse(http.StatusBadRequest, e)
			c.JSON(res.Status, res)
			return
		}

		var result sql.Result
		var e error
		if result, e = db.ExecContext(ctx, "DELETE FROM products WHERE guid=?", binding.GUID); e != nil {
			res := internal.NewHttpResponse(http.StatusInternalServerError, e)
			c.JSON(res.Status, res)
			return
		}

		if nProducts, _ := result.RowsAffected(); nProducts == 0 {
			res := internal.NewHttpResponse(http.StatusNotFound, sql.ErrNoRows)
			c.JSON(res.Status, res)
			return
		}

		c.JSON(http.StatusNoContent, "deleted")
	}
}
