package controllers

import (
	"fmt"
	"net/http"
	"quickstart/config"
	"quickstart/models"

	"github.com/gin-gonic/gin"
)

func GetAllAUthors(c *gin.Context){
	var authors []models.Author

	//get al authors
	config.DB.Find(&authors)
	c.JSON(http.StatusOK,authors)
}

//get a single author by id
func GetAuthor(c *gin.Context){
	id :=c.Param("id")
	var author models.Author

	//get author with thier books(preload relationship)
	if err:=config.DB.Preload("Books").First(&author,id).Error;err !=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"Author not found"})
		return
	}

	  // Debug: Print to see what's in the books
    for _, book := range author.Books {
        fmt.Printf("Book: ID=%d, Title=%s, Year=%d, AuthorID=%d\n", 
            book.ID, book.Title, book.Year, book.AuthorID)
    }
	//conver books to booksimple to avaoid circular reference
	bookSimples := make([]models.BookSimple,len(author.Books))
	for i,book:=range author.Books{
		bookSimples[i] = models.BookSimple{
			ID:book.ID,
			CreatedAt: book.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
            UpdatedAt: book.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			Title: book.Title,
			Year: book.Year,
			AuthorID: book.AuthorID,
		}

		
	}

	//create clean response struct
	response := models.AuthorResponse{
		ID : author.ID,
		CreatedAt: author.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
        UpdatedAt: author.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		Name: author.Name,
		Country: author.Country,
		Books:bookSimples, //books withour nested author
	}
	c.JSON(http.StatusOK,response)
}


//create a new author
func CreateAuthor(c *gin.Context){
	var newAuthor models.Author

	if err:=c.BindJSON(&newAuthor); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	if err:=config.DB.Create(&newAuthor).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{	"error":"failed to create author"})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"message": "Author created successfully",
		"author": newAuthor,
	})
}


//update author -update an author
func UpdateAuthor(c *gin.Context){
	id := c.Param("id")
	var author models.Author

	if err := config.DB.First(&author,id).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"Author not found"})
		return
	}

	if err := c.BindJSON(&author);err !=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
		return
	}
	config.DB.Save(&author)

	c.JSON(http.StatusOK,gin.H{
		"message": "Author updated successfully",
		"author": author,
	})


}

// DeleteAuthor - Delete an author (and all their books due to CASCADE)

func DeleteAuthor(c *gin.Context){
id:=c.Param("id")
var author models.Author

if err:= config.DB.First(&author,id).Error;err!=nil{
	c.JSON(http.StatusNotFound,gin.H{"error":"Author not found"})
	return
}
if err:=config.DB.Select("Books").Delete(&author).Error;err!=nil{
	c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to delete author"})
	return
}

c.JSON(http.StatusOK,gin.H{
	"message":"Author deleted successfully (all books deleted to)"})
}