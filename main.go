package main

import (
	"fmt"
	"log"
	"pustaka/book"
	"pustaka/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:12345678@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}
	db.AutoMigrate(&book.Book{})

	book := book.Book{}
	book.Title = "itikaf"
	book.Price = 100000
	book.Rating = 4
	book.Description = "jamaah masjid "

	err = db.Create(&book).Error
	if err != nil{ 
		fmt.Println("=====================")
		fmt.Println("error")
		fmt.Println("=====================")

	}

	router := gin.Default()

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"name":    "fahmi illahi",
	// 		"Bio":     "Pengangguran Taat",
	// 		"Tinggal": "Dimasjid",
	// 	})
	// })
	v1 := router.Group("/v1")

	v1.GET("/hello", handler.GetQuery)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.BookInputData)

	router.Run()
}
