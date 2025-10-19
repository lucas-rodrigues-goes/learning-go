package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var bookList = []Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 1},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 5},
}

func GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bookList)
}

func Get(c *gin.Context) {
	id := c.Param("id")

	for i := range bookList {
		if bookList[i].ID == id {
			c.IndentedJSON(http.StatusOK, bookList[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
}

func Create(c *gin.Context) {
	var newBook Book

	err := c.BindJSON(&newBook)
	if err != nil {
		return
	}

	bookList = append(bookList, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func Update(c *gin.Context) {
	id := c.Param("id")
	var patchData struct {
		Title    *string `json:"title"`
		Author   *string `json:"author"`
		Quantity *int    `json:"quantity"`
	}

	err := c.BindJSON(&patchData)
	if err != nil {
		return
	}

	for i := range bookList {
		if bookList[i].ID == id {
			if patchData.Title != nil {
				bookList[i].Title = *patchData.Title
			}
			if patchData.Author != nil {
				bookList[i].Author = *patchData.Author
			}
			if patchData.Quantity != nil {
				bookList[i].Quantity = *patchData.Quantity
			}
			c.IndentedJSON(http.StatusOK, bookList[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	for i := range bookList {
		if bookList[i].ID == id {
			bookList = append(bookList[:i], bookList[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
}
