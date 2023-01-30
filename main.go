package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"golang-api/book"
	"golang-api/handler"
)

func main() {
	// Connect Golang to Database
	dsn := "root:@tcp(127.0.0.1:3306)/golang_api?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if error connect
	if err != nil{
		log.Fatal("Database connected error")
	}

	// Auto Migration table database
	db.AutoMigrate(&book.Book{})

	// Create
	// masukkan data dalam construct
	// book := book.Book{}
	// book.Title = "Haped sang developer"
	// book.Price = 120000
	// book.Discount = 10
	// book.Rating = 5
	// book.Description = "Ini adalah buku yang menceritakan kisah perjalanan seorang anak dari madiun"

	//save to database
	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("=====================")
	// 	fmt.Println("Gagal Memasukkan Buku")
	// 	fmt.Println("=====================")
	// }

	// Read 
	var books []book.Book
	// var book book.Book
	// Read First Data in table
	// err = db.Debug().First(&book).Error

	// Read Last Data in table
	// err = db.Debug().Last(&book).Error

	// Read Data By Id in table
	// err = db.Debug().First(&book, 1).Error

	// Read All Data in table
	// err = db.Debug().Find(&books).Error

	// Read Where Data in table (title)
	// err = db.Where("title = ?", "Kisah hidup Nyut").Find(&books).Error

	// Read Where Data in table (Rating)
	err = db.Where("rating = ?", 5).Find(&books).Error
	if err != nil {
		fmt.Println("======================")
		fmt.Println("Gagal Menampilkan Buku")
		fmt.Println("======================")
	}

	// perulangan untuk menampilkan semua data pada table
	for _, b := range books{
		fmt.Println("title:", b.Title)
		fmt.Println("book object %v", b)
	}

	//print 1 data
	// fmt.Println("title:", book.Title)
	// fmt.Println("book object %v", book)

	//Update
	// var book book.Book

	// mencari data yang ingin di update
	// err = db.Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("======================")
	// 	fmt.Println("Gagal Menemukan Buku")
	// 	fmt.Println("======================")
	// }

	// Memasukkan perubhaan data
	// book.Title = "Kisah perjalanan Hidup Nyut part 2"
	// simpan perubahan
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("======================")
	// 	fmt.Println("Gagal Mengupdate Buku")
	// 	fmt.Println("======================")
	// }

	// Delete 
	var book book.Book
	err = db.Where("id = ?", 1).First(&book).Error
	if err != nil {
		fmt.Println("======================")
		fmt.Println("Gagal Menampilkan Buku")
		fmt.Println("======================")
	}

	// Menghapus data pada tale
	err = db.Delete(&book).Error
	if err != nil {
		fmt.Println("======================")
		fmt.Println("Gagal Mengahapus Buku")
		fmt.Println("======================")
	}
	// Router Golang
	router := gin.Default()

	// api verisoning digunakan agar saat kita ada perubahan mayor atau update besar yang versi sebelumnya tidak rusak
	// caranya dengan meng group route nya

	v1 := router.Group("/v1")
	
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)
	router.Run()
}





