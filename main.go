package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
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
	if e = db.Ping(); e != nil {
		log.Fatalf("error: %v", e)
	}

	router.GET("/products", controllers.GetProducts(db))
	router.GET("/products/:guid", controllers.GetProduct(db))
	router.POST("/products", controllers.PostProduct(db))
	router.DELETE("/products/:guid", controllers.DeleteProduct(db))
	router.PUT("/products/:guid", controllers.PutProducts)

	log.Fatal(router.Run(address))

}
