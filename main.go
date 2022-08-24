package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type urlBinding struct {
	ID string `uri:"id"`
}

type Products struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProductsShouldBind struct {
	ID   string `form:"id"`
	Name string `form:"name"`
}

type headerBinding struct {
	RequestId string `header:"x-request-id"`
}

func main() {
	router := gin.Default()
	address := ":3000"

	router.GET("/products/:id", func(ctx *gin.Context) {
		var binding urlBinding
		if e := ctx.ShouldBindUri(&binding); e != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
			return
		}

		fmt.Println("the id is ", binding.ID)
		ctx.String(http.StatusOK, "the id is %s", binding.ID)
	})
	// should bind json
	router.POST("/products/", func(ctx *gin.Context) {
		var bindJSon Products

		if err := ctx.ShouldBindJSON(&bindJSon); err != nil {
			ctx.String(http.StatusBadRequest, "the data is not json")
			return
		}

		fmt.Println("Product:", bindJSon)
		ctx.String(http.StatusOK, "the id is %s", bindJSon.ID)
	})
	// shoud bind
	router.POST("/products2", func(ctx *gin.Context) {
		var products2 ProductsShouldBind
		if err := ctx.ShouldBind(&products2); err != nil {
			ctx.String(http.StatusBadRequest, "the data is not json")
			return
		}

		fmt.Println("Product:", products2)
		ctx.String(http.StatusOK, "the id is %s", products2.ID)
	})

	router.POST("/header", func(ctx *gin.Context) {
		binding := headerBinding{}
		if err := ctx.ShouldBindHeader(&binding); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}
		fmt.Println("the request id is ", binding.RequestId)
		ctx.String(http.StatusOK, "the request id is %s", binding.RequestId)
	})

	log.Fatal(router.Run(address))

}
