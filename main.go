package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/course-gin/internal/controllers"
)

func main() {
	router := gin.Default()
	address := ":3000"
	var db *sql.DB
	var e error
	if db, e = sql.Open("sqlite3", "./data.db"); e != nil {
		log.Fatalf("error: %v", e)
	}
	defer db.Close()

	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:guid", controllers.GetProduct)
	router.POST("/products", controllers.PostProduct)
	router.DELETE("/products/:guid", controllers.DeleteProduct)
	router.PUT("/products/:guid", controllers.PutProducts)

	log.Fatal(router.Run(address))

}
