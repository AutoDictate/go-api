package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct {
	ID       	string 	`json:"id"`
	Title    	string 	`json:"title"`
	Author   	string 	`json:"author"`
	Quantity 	int    	`json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, books)

}

func bookById(c *gin.Context) {

	id := c.Param("id")
	book, err := getBooksById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBooksById(id string) (*book, error) {
	
	for i, val := range books {
		if val.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book Not found")
}

func checkOutbook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Missing id query parameter"})
		return
	}

	book, err := getBooksById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not Found"})
		return
	}

	if book.Quantity <=0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not available."})
		return
	}

	book.Quantity -=1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Missing id query parameter"})
		return
	}

	book, err := getBooksById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not Found"})
		return
	}

	book.Quantity +=1
	c.IndentedJSON(http.StatusOK, book)
}

func createBooks(c *gin.Context) {

	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)

	c.IndentedJSON(http.StatusCreated, books)
}

func main() {

	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/addBooks", createBooks)
	router.GET("/getBook/:id", bookById)
	router.PATCH("/checkout", checkOutbook)
	router.PATCH("/returnBook", returnBook)
	router.Run("localhost:8180")
}
