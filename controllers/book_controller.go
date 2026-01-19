package controllers

import (
	"quickstart/database"
	"quickstart/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllBooks handles GET /books
func GetAllBooks(c *gin.Context) {
    c.JSON(http.StatusOK, database.Books)
}

// GetBook handles GET /books/:id
func GetBook(c *gin.Context) {
    id := c.Param("id")
    
    for _, book := range database.Books {
        if book.ID == id {
            c.JSON(http.StatusOK, book)
            return
        }
    }
    
    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

// CreateBook handles POST /books
func CreateBook(c *gin.Context) {
    var newBook models.Book
    
    if err := c.BindJSON(&newBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // Check if ID already exists
    for _, book := range database.Books {
        if book.ID == newBook.ID {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Book with this ID already exists",
            })
            return
        }
    }
    
    database.Books = append(database.Books, newBook)
    c.JSON(http.StatusCreated, gin.H{
        "message": "Book created",
        "book": newBook,
    })
}

// UpdateBook handles PUT /books/:id
func UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var updatedBook models.Book
    
    if err := c.BindJSON(&updatedBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    for i, book := range database.Books {
        if book.ID == id {
            database.Books[i] = updatedBook
            database.Books[i].ID = id
            
            c.JSON(http.StatusOK, gin.H{
                "message": "Book updated",
                "book": database.Books[i],
            })
            return
        }
    }
    
    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

// DeleteBook handles DELETE /books/:id
func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    
    for i, book := range database.Books {
        if book.ID == id {
            database.Books = append(database.Books[:i], database.Books[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
            return
        }
    }
    
    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}