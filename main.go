// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Book struct {
// 	ID string `json:"id"`
// 	Title string `json:"title"`
// 	Author string `json:"author"`
// 	Year int `json:"year"`
// }

// func main(){
// 	//create a default router
// 	r:=gin.Default()
// 	//define a simple get point retrieve all data at this endpoint
// 	r.GET("/books",func(c *gin.Context){
// 		//return json response
// 		c.JSON(http.StatusOK,gin.H{
// 			"message":"pong",
// 		})
// 	})

// 	//post create new record
// 	r.POST("/books",func(c *gin.Context) {
//        c.JSON(http.StatusOK,gin.H{
// 		"message":"Create a book",
// 	   })
// 	})

// 	//put update a new book
// 	r.PUT("/books",func(c *gin.Context){
// 		c.JSON(http.StatusOK,gin.H{
// 			"message":"update a book",
// 		})
// 	})

// 	//delete a book
// 	r.DELETE("/books",func(c *gin.Context){
// 		c.JSON(http.StatusOK,gin.H{
// 			"message":"Delete a book",
// 		})
// 	})
// 	//

// //user parameters exaple /user/123
// r.GET("/user/:id",func(c *gin.Context){
// 	id:=c.Param("id")//gets 123

// 	c.JSON(http.StatusOK,gin.H{
//          "user_id":id,
// 	})
// })

// // Multiple parameters: /user/123/book/456
// // r.GET("/user/:userId/book/:bookId",func(c *gin.Context){
// // 	userId:=c.Param("userId") //gets 123
// // 	bookId:= c.Param("bookId") //gets 456

// // 	c.JSON(http.StatusOK,gin.H{
// // 		"user_id":userId,
// // 		"book_id":bookId,
// // 	})
// // })
// // Example: /search?keyword=golang&page=2 come after ?
// r.GET("/search",func(c *gin.Context){
// keyword :=c.Query("keyword")
// page :=c.Query("page")

//  // With default value if parameter is missing
//  limit :=c.DefaultQuery("limit","10")

//  c.JSON(http.StatusOK,gin.H{
// 	"keyword":keyword,
// 	"page":page,
// 	"limit":limit,
//  })

// //get all users (with optional filtering)
// //url:/users?role=admin
// r.GET("/users?role=admin",func(c *gin.Context){
//   role :=c.DefaultQuery("role","all")

//   c.JSON(http.StatusOK,gin.H{
// 	"message":"getting all users",
// 	"filter": role,
//   })
// })
// })//

// //get specific user by id
// //url: /users/42
// // r.GET("/users",func(c *gin.Context){
// // 	id:=c.Param("id")

// // 	c.JSON(http.StatusOK,gin.H{
// // 		"message":"getting user",
// // 		"user_id":id,
// // 	})

// // })

// //create a new user
// r.POST("/users",func(c *gin.Context){
// 	c.JSON(http.StatusOK,gin.H{
// 		"message":"user created successfully",
// 	})
// })

// //update a user
// //url: /users/42
// r.PUT("/users/:id",func(c *gin.Context) {

// id :=c.Param("id")

// c.JSON(http.StatusOK,gin.H{
// 	"message":"user updated",
// 	"user_id":id,
// })
// })

// //delete a user
// //url /user/42
// r.DELETE("/users/:id",func(c *gin.Context){
// 	id :=c.Param("id")

// 	c.JSON(http.StatusOK,gin.H{
// 		"message":"user deleted",
// 		"user_id":id,
// 	})
// })

// //query admins
// r.GET("/admin",func(c *gin.Context) {
// 	author:=c.Query("author")
// 	year :=c.Query("year")
//  // Method 2: Get with a default value
//         limit := c.DefaultQuery("limit", "10")  // default is "10" if not provided
// response :=gin.H{
//          "message":"getting all admins",
// }
// 	  // Check if author filter was provided
//         if author != "" {
//             response["filtered_by_author"] = author
//         }

//         // Check if year filter was provided
//         if year != "" {
//             response["filtered_by_year"] = year
//         }
// 		  response["limit"] = limit

// c.JSON(http.StatusOK,response)

// })

// //pens
// r.POST("/pens",func(c *gin.Context) {
// 	var newBook Book

// 	//bind the json in the request body to our struct
// 	if err:=c.BindJSON(&newBook); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "invalid JSON DATA",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "pen created successfully",
// 		"book":    newBook,
// 	})
// })

// 			//start the server on port 8080
// r.Run(":8080")
// }

package main

import (
	"quickstart/config"
	"quickstart/routes"

	"github.com/gin-gonic/gin"
)

//first middleware

	func main(){
		//connect to database
		config.ConnectDatabase()
		r:=gin.New()//default middleware logger and recovery


		//setup all routes
		routes.SetupRoutes(r)

		// //get all books
		// r.GET("/books",func(c *gin.Context){
		// 	c.JSON(http.StatusOK,books)
		// })

		// //get a single book by id
		// r.GET("/books/:id",func(c *gin.Context) {
        //    id :=c.Param("id")

		//    for _,book:=range books{
		// 	if book.ID ==id{
		// 		c.JSON(http.StatusOK,book)
		// 		return
		// 	}
		//    }

		//    //book not found 
		//    c.JSON(http.StatusNotFound,gin.H{
		// 	"error":"Book not found",
		//    })
		// })

		// //post -create a book
		// r.POST("/books",func(c *gin.Context) {
        //     var newBook Book


		// 	if err:=c.BindJSON(&newBook);err!=nil{
		// 		c.JSON(http.StatusBadRequest,gin.H{
		// 			"error":err.Error(),
		// 		})
		// 		return
		// 	}


		// 	 //custom errors
		// 	 if strings.Contains(strings.ToLower(newBook.Title),"badword"){
		// 		c.JSON(http.StatusBadRequest,gin.H{
		// 			"error":"Title containing inappropriate content",
		// 		})
		// 		return
		// 	 }

		// 	 //check if the authors name is all uppercase
		// 	 if newBook.Author == strings.ToUpper(newBook.Author){
		// 		c.JSON(http.StatusBadRequest,gin.H{
		// 			"error":"Author name looks suspicious",
		// 		})
		// 		return
		// 	 }

		// 	 //check if year is in future
		// 	 currentYear := 2026
		// 	 if newBook.Year > currentYear{
		// 		c.JSON(http.StatusBadRequest,gin.H{
		// 			"error":"Year cannot be in the future",
		// 		})
		// 		return
		// 	 }

			
		// 	//add book to database
		// 	books =append(books,newBook)

		// 	c.JSON(http.StatusCreated,gin.H{
		// 		"message":"Book created",
		// 		"book": newBook,
		// 	})
 		// })

		// //put update a book
		// r.PUT("/books/:id",func(c *gin.Context) {
        //    id :=c.Param("id")
		//    var updatedBook Book

		//    if err:=c.BindJSON(&updatedBook);err!=nil{
		// 	c.JSON(http.StatusBadRequest,gin.H{
		// 		"error":err.Error(),
		// 	})
		// 	return
		//    }

		//    //find and update book
		//    for i,book:=range books{
		// 	if book.ID ==id{
		// 		books[i] = updatedBook
		// 		books[i].ID =id //keep id the same

		// 		c.JSON(http.StatusOK,gin.H{
		// 			"message":"Book updated",
		// 			"book":books[i],
		// 		})
		// 		return
		// 	}
		//    }

		//    c.JSON(http.StatusNotFound,gin.H{
		// 	"error":"Book not found",
		//    })
		// })

		// //delete a book
		// r.DELETE("/books/:id",func(c *gin.Context){
		// 	id :=c.Param("id")

		// 	//find and delete a book
		// 	for i,book:= range books{
		// 		if book.ID ==id{
		// 			books = append(books[:i],books[i+1:]...)

		// 			c.JSON(http.StatusOK,gin.H{
		// 				"message":"Book deleted",
		// 			})
		// 			return
		// 		}
		// 	}

		// 	c.JSON(http.StatusNotFound,gin.H{
		// 		"error":"Book not found",
		// 	})
		// })
		r.Run(":8080")

}