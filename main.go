package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "errors"
)

type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Qty int `json:"qty"`
}

var books = []book{
	{ID: "1", Title: "Golang", Author: "Google", Qty: 10},
	{ID: "2", Title: "Java", Author: "Oracle", Qty: 20},
	{ID: "3", Title: "Python", Author: "Python Software Foundation", Qty: 30},
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func main() {
	r := gin.Default()

	r.GET("/books", getBooks)
	// r.GET("/books/:id", getBook)
	// r.POST("/books", createBook)
	// r.PUT("/books/:id", updateBook)
	// r.DELETE("/books/:id", deleteBook)

	r.Run("localhost:8080")
}