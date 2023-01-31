package main

import (
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

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// Router Golang
	router := gin.Default()

	// api verisoning digunakan agar saat kita ada perubahan mayor atau update besar yang versi sebelumnya tidak rusak
	// caranya dengan meng group route nya
	v1 := router.Group("/v1")
	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/hello", bookHandler.HelloHandler)
	v1.GET("/books/:id/:title", bookHandler.BooksHandler)
	v1.GET("/query", bookHandler.QueryHandler)
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)
	router.Run()

	// Find All with repository
	// books, err := bookRepository.FindAll()
	// for _, book := range books {
	// 	fmt.Println("title:", book.Title)
	// }

	// Find By Id 
	// bookById, err := bookRepository.FindByID(1)
	// fmt.Println("title:", bookById.Title)

	// Create Book langsung repository
	// book := book.Book{
	// 	Title: "ini kisah nyut",
	// 	Description: "ini buku bagus",
	// 	Price: 89000,
	// 	Discount: 0,
	// 	Rating: 5,
	// }
	// bookRepository.Create(book)

	// Create book menggunakan service layer
	// bookRequest := book.BookRequest{
	// 	Title: "ini kisah Haped",
	// 	Price: "90000",
	// }
	// bookService.Create(bookRequest)

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
	// var books []book.Book
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
	// err = db.Where("rating = ?", 5).Find(&books).Error
	// if err != nil {
	// 	fmt.Println("======================")
	// 	fmt.Println("Gagal Menampilkan Buku")
	// 	fmt.Println("======================")
	// }

	// perulangan untuk menampilkan semua data pada table
	// for _, b := range books{
	// 	fmt.Println("title:", b.Title)
	// 	fmt.Println("book object %v", b)
	// }

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
	// var book book.Book
	// err = db.Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("======================")
	// 	fmt.Println("Gagal Menampilkan Buku")
	// 	fmt.Println("======================")
	// }

	// Menghapus data pada tale
	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("======================")
	// 	fmt.Println("Gagal Mengahapus Buku")
	// 	fmt.Println("======================")
	// }
}





