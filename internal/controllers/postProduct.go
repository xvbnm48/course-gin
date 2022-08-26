package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xvbnm48/course-gin/internal"
)

type postProduct struct {
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Description string  `json:"description" binding:"omitempty,max=250"`
}

type Product struct {
	GUID        string  `json:"guid"`
	Name        string  `json:"name" `
	Price       float64 `json:"price" `
	Description string  `json:"description" `
	CreatedAt   string  `json:"createdAt" `
}

func PostProduct(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload postProduct
		var c = ctx.Request.Context()
		if e := ctx.ShouldBindJSON(&payload); e != nil {
			var res = internal.NewHttpResponse(400, e)
			ctx.JSON(400, res)
			return
		}
		guid := uuid.New().String()
		createdAt := time.Now().Format(time.RFC3339)
		if _, e := db.ExecContext(c, "INSERT INTO products (guid, name, price, description, createdAt) VALUES (?, ?, ?,?, ?)", guid, payload.Name, payload.Price, payload.Description, createdAt); e != nil {
			res := internal.NewHttpResponse(500, e)
			ctx.JSON(500, res)
			return
		}
		var product Product
		var row = db.QueryRow("SELECT guid,name,price,description,createdAt FROM products WHERE guid=?", guid)

		if err := row.Scan(&product.GUID, &product.Name, &product.Price, &product.Description, &product.CreatedAt); err != nil {
			fmt.Println(err)
			res := internal.NewHttpResponse(500, err)
			ctx.JSON(500, res)
			return
		}

		fmt.Println(payload)
		res := internal.NewHttpResponse(http.StatusCreated, product)
		ctx.Writer.Header().Add("Location", fmt.Sprintf("/products/%s", guid))
		ctx.JSON(200, res)
	}
}
