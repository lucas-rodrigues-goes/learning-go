package main

import (
	"library-api/books"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Routes
	router.GET("/books", books.GetAll)
	router.GET("/books/:id", books.Get)
	router.POST("/books", books.Create)
	router.PATCH("/books/:id", books.Update)
	router.DELETE("/books/:id", books.Delete)

	// Run
	router.Run("localhost:8080")
}
