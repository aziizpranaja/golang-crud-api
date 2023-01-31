package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"golang-api/book"
)

type bookHandler struct{
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler{
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Aziiz Pranaja",
		"desc": "website development enthusiast",
	})
}

func (h *bookHandler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "Hello World",
		"subtitle": "Belajar golang kidzztoo",
	})
}

func (h *bookHandler) BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func (h *bookHandler) QueryHandler(c *gin.Context) {
	title := c.Query("title")
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{"name": name, "title": title})
}

func (h *bookHandler) CreateBook(c *gin.Context){
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil{
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error in field %s condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertBookResponse(book),
	})
}

// Get all books
func (h *bookHandler) GetBooks(c *gin.Context){
		books, err := h.bookService.FindAll()
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		var booksResponse []book.BookResponse

		for _, b := range books{
			bookResponse := convertBookResponse(b)

			booksResponse = append(booksResponse, bookResponse)
		}
		
		c.JSON(http.StatusOK, gin.H{
			"data": booksResponse,
		})
}

// Get Book By Id
func (h *bookHandler) GetBook(c *gin.Context){
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	bookById, err := h.bookService.FindByID(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	bookResponse := convertBookResponse(bookById)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) UpdateBook(c *gin.Context){
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil{
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error in field %s condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Update(id, bookRequest)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertBookResponse(book),
	})
}

// Delete Book 
func (h *bookHandler) DeleteBook(c *gin.Context){
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	bookById, err := h.bookService.Delete(int(id))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	bookResponse := convertBookResponse(bookById)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

// Convert Response 
func convertBookResponse(b book.Book) book.BookResponse{
	return book.BookResponse{
		ID: b.ID,
		Title: b.Title,
		Description: b.Description,
		Price: b.Price,
		Rating: b.Rating,
		Discount: b.Discount,
	}
}