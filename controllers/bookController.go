package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Book struct {
	BookID int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas []Book

func GetAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Books retrieved successfully",
		"data":    BookDatas,
	})
}

func CreateBook(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	book.BookID = len(BookDatas) + 1
	BookDatas = append(BookDatas, book)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Book created successfully",
		"data":    book,
	})
}

func GetBookByID(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	var book Book
	condition := false
	for _, b := range BookDatas {
		if bookID == b.BookID {
			condition = true
			book = b
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Book not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Book retrieved successfully",
		"data":    book,
	})
}

func UpdateBook(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	var book Book
	condition := false
	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, b := range BookDatas {
		if bookID == b.BookID {
			condition = true
			book.BookID = b.BookID
			BookDatas[i] = book
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Book not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Book updated successfully",
		"data":    book,
	})
}

func DeleteBook(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	condition := false
	for i, b := range BookDatas {
		if bookID == b.BookID {
			condition = true
			BookDatas = append(BookDatas[:i], BookDatas[i+1:]...)
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Book not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Book deleted successfully",
	})
}
