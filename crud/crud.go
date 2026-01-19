package crud

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type Book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Year int `json:"year"`
}


var books =[]Book{
	{ID:"1",Title:"The Great Gatsby",Author:"F. Scott Fitzgerald",Year:1925},
	{ID:"2",Title:"To Kill a Mockingbird",Author:"Harper Lee",Year:1960},
}
	func main(){
		r:=gin.Default()


		//get all books
		r.GET("/books",func(c *gin.Context){
			c.JSON(http.StatusOK,books)
		})

		//get a single book by id
		r.GET("/bookk/:id",func(c *gin.Context) {
           id :=c.Param("id")

		   for _,book:=range books{
			if book.ID ==id{
				c.JSON(http.StatusOK,book)
				return
			}
		   }

		   //book not found 
		   c.JSON(http.StatusNotFound,gin.H{
			"error":"Book not found",
		   })
		})

		//post -create a book
		r.POST("/books",func(c *gin.Context) {
            var newBook Book


			if err:=c.BindJSON(&newBook);err!=nil{
				c.JSON(http.StatusBadRequest,gin.H{
					"error":"invalid JSON",
				})
				return
			}

			//add book to database
			books =append(books,newBook)

			c.JSON(http.StatusCreated,gin.H{
				"message":"Book created",
				"book": newBook,
			})
 		})

		//put update a book
		r.PUT("/books/:id",func(c *gin.Context) {
           id :=c.Param("id")
		   var updatedBook Book

		   if err:=c.BindJSON(&updatedBook);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"invalid JSON",
			})
			return
		   }

		   //find and update book
		   for i,book:=range books{
			if book.ID ==id{
				books[i] = updatedBook
				books[i].ID =id //keep id the same

				c.JSON(http.StatusOK,gin.H{
					"message":"Book updated",
					"book":books[i],
				})
				return
			}
		   }

		   c.JSON(http.StatusNotFound,gin.H{
			"error":"Book not found",
		   })
		})

		//delete a book
		r.DELETE("/books/:id",func(c *gin.Context){
			id :=c.Param("id")

			//find and delete a book
			for i,book:= range books{
				if book.ID ==id{
					books = append(books[:i],books[i+1:]...)

					c.JSON(http.StatusOK,gin.H{
						"message":"Book deleted",
					})
					return
				}
			}

			c.JSON(http.StatusNotFound,gin.H{
				"error":"Book not found",
			})
		})
		r.Run(":8080")

}