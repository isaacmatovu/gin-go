package controllers

import (
	"math"
	"quickstart/config"
	"quickstart/models"
	"strconv"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

type pagination struct {
    NextPage int
    PreviousPage int
    CuurentPage int
    TotalPages int
}

// GetAllBooks handles GET /books
func GetAllBooks(c *gin.Context) {
    var books []models.Book

    page:=1
    perPage:=10
 
pageStr := c.Param("page")

if pageStr != "" {
page, _ = strconv.Atoi(pageStr)
    
}
    
    offset := (page - 1) * perPage
    //total records
    var totalCount int64
    config.DB.Model(&models.Book{}).Count(&totalCount)

    //number of pages
    totalPages := math.Ceil(float64(totalCount / int64(perPage)))


    //select * from books where deleted_at is null
    config.DB.Limit(perPage).Offset(offset).Preload("Author").Find(&books)

    
    
    c.JSON(http.StatusOK, gin.H{
        "message": "Books retrieved successfully",
        "books": books,
        "pagination": pagination{
            NextPage: page + 1,
            PreviousPage: page - 1,
            CuurentPage: page,
            TotalPages:int(totalPages),  
        },
    })
}

// GetBook handles GET /books/:id
func GetBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
   
    
  //select * from books where id =? and deleted_at IS NULL
  if err := config.DB.Preload("Author").Preload("Categories").First(&book,id).Error;err!=nil{
    c.JSON(http.StatusNotFound,gin.H{"error":"Book not found"})
    return
  }


 var authorPtr *models.Author

    if book.AuthorID !=0{
        authorPtr = &book.Author
    } else{
        authorPtr = nil
    }

  categories:= make([]models.CategorySimple,len(book.Categories))
  for i,category := range book.Categories{
    categories[i] = models.CategorySimple{
       ID : category.ID,
       CreatedAt: category.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
       UpdatedAt: category.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
        Name: category.Name,
    }
  }

reponse:= models.BookCategory{
    ID :book.ID,
    CreatedAt: book.CreatedAt.Format(time.RFC3339),
    UpdatedAt: book.UpdatedAt.Format(time.RFC3339),
    Title:book.Title,
    Year:book.Year,
    AuthorID: book.AuthorID,
    Author: authorPtr,
    Categories: categories,

}
c.JSON(http.StatusOK,reponse)

}

// CreateBook handles POST /books
func CreateBook(c *gin.Context) {
    var newBook models.Book
    
    if err := c.BindJSON(&newBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    //check if author exists
    var author models.Author
    if err:=config.DB.First(&author,newBook.AuthorID).Error;err!=nil{
        c.JSON(http.StatusBadRequest,gin.H{"error":"Author not found"})
    }
    //insert into books(title,author,year,created_at,updatedat)
    if err:= config.DB.Create(&newBook).Error;err !=nil{
        c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to create book"})
        return
    }

    config.DB.Preload("Author").First(&newBook,newBook.ID)
    c.JSON(http.StatusCreated,gin.H{"message":"book created successfully","book":newBook})
}

// UpdateBook handles PUT /books/:id
func UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    
 //find the book
 if err := config.DB.First(&book,id).Error; err !=nil {
    c.JSON(http.StatusNotFound,gin.H{"error":"Book not found"})
    return
 }

 //bind updated book
 if err := c.BindJSON(&book); err !=nil {
    c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
 return
}

//update books set tittle =?,author=?,year=?,updated_at=? where id =?
config.DB.Save(&book)

//reload with author information
config.DB.Preload("Author").First(&book,book.ID)
    
c.JSON(http.StatusOK,gin.H{
    "message":"book updated successfully",
    "book": book,
})
}

// DeleteBook handles DELETE /books/:id
func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
 //find the book
 if err := config.DB.First(&book,id).Error; err!=nil{
    c.JSON(http.StatusNotFound,gin.H{"error":"Book not found"})
    return
 }

 //update books set deleted_at=? where id=?
 //soft delete -doesnt actually remove from the database
 config.DB.Delete(&book)
    
c.JSON(http.StatusOK,gin.H{"message":"Book deleted successfully"})
}


//get the first 10 books
func GetFirst10(c *gin.Context){
    var book []models.Book
    //get the first10
    if err := config.DB.Limit(10).Find(&book).Error; err!=nil{
        c.JSON(http.StatusNotFound,gin.H{"error":"no books found"})
        return
    }
    c.JSON(http.StatusOK,gin.H{"success":"books found","books":book})

}

//getbooksby author- get all books of a specific author
func GetBooksByAuthor(c *gin.Context){
    authorID := c.Param("authorId")
    var books []models.Book

    //find books where author_id matches
    config.DB.Where("author_id = ?",authorID).Preload("Author").Find(&books)


 booksimple:= make([]models.BookSimple,len(books))
 for i,book:=range books{
    booksimple[i] = models.BookSimple{
        ID: book.ID,
        CreatedAt: book.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
        UpdatedAt: book.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
        Title: book.Title,
        Year: book.Year,
        AuthorID: book.AuthorID,
    }
 }


  c.JSON(http.StatusOK,gin.H{"books":booksimple})
}