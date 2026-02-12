package controllers

import (
	"net/http"
	"quickstart/config"
	"quickstart/models"

	"github.com/gin-gonic/gin"
)


func GetAllCategories(c *gin.Context){
	var categories []models.Category

	config.DB.Find(&categories)
	c.JSON(http.StatusOK,categories)
}


//getcategory - get a category with all its books 
func GetCategory(c *gin.Context){
	id := c.Param("id")
var category models.Category


//preload books (and each book with author)
if err:= config.DB.Preload("Books.Author").First(&category,id).Error;err !=nil{
	c.JSON(http.StatusNotFound,gin.H{"error":"Category not found"})
return 
}



  books:= make([]models.BookResponse,len(category.Books))
  for i ,book:= range category.Books{
	books[i] = models.BookResponse{
		ID: book.ID,
		CreatedAt: book.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: book.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		Title: book.Title,
		Year: book.Year,
		AuthorID: book.AuthorID,
		Author: &book.Author, // Include author details in the book response
	
  }
}
response:=models.CategoryResponse{
	ID: category.ID,
	CreatedAt: category.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	UpdatedAt: category.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	Name: category.Name,
	Books: books, // Include the books in the response
}

c.JSON(http.StatusOK,response)

}


//create category - create  a new category 
func CreateCategory(c *gin.Context){
	var newCategory models.Category

	//bindjson
	if err:= c.BindJSON(&newCategory); err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	if err:= config.DB.Create(&newCategory).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"failed to create category"})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"message": "Category created successfully",
		"category": newCategory,
	})

	
}


//addbokk to a category 
	func AddBookToCategory(c *gin.Context){
categoryID := c.Param("categoryid")
bookID := c.Param("bookId")


var category models.Category
var book models.Book

//find a category
if err:= config.DB.First(&category,categoryID).Error; err !=nil{
	c.JSON(http.StatusNotFound,gin.H{"error":"Category not found"})
	return
}


//find a book
if err:= config.DB.First(&book,bookID).Error;err!=nil{
	c.JSON(http.StatusNotFound,gin.H{"error":"Book not found"})
	return
}

//add book to category 
config.DB.Model(&category).Association("Books").Append(&book)

c.JSON(http.StatusOK,gin.H{"message":"Book added to category successfullt"})

	}

	//remove book from category 
	func RemoveBookFromCategory(c *gin.Context){
		categoryID := c.Param("categoryId")
		bookID := c.Param("bookId")

		var category models.Category
		var book models.Book
		if err:= config.DB.First(&category,categoryID).Error;err !=nil{
			c.JSON(http.StatusNotFound,gin.H{"error":"Category not found"})
			return
		}

		if err := config.DB.First(&book,bookID).Error;err!=nil{
			c.JSON(http.StatusNotFound,gin.H{"error":"Book not found"})
			return
		}

		//remove book from category (deletes entry from join table)
		config.DB.Model(&category).Association("Books").Delete(&book)

		c.JSON(http.StatusOK,gin.H{"message":"Book removed from category repository"})
	}