package routes

import (
	"quickstart/controllers"
	"quickstart/middleware"

	"github.com/gin-gonic/gin"
)




func SetupRoutes(r *gin.Engine){

	  //apply global middleware
	  r.Use(middleware.Logger())
	  r.Use(middleware.CORS())

//author routes pulic
r.GET("/authors",controllers.GetAllAUthors)
r.GET("/author/:id",controllers.GetAuthor)
r.GET("/authors/:authorId/books",controllers.GetBooksByAuthor)


r.GET("/categories", controllers.GetAllCategories)
r.GET("/categories/:id", controllers.GetCategory)

	  //public routes
	 //health check route
      r.GET("/health",controllers.HealthCheck)
	  r.GET("/books",controllers.GetAllBooks)
	  r.GET("/books/:id",controllers.GetBook)
	  r.GET("/book",controllers.GetFirst10)
	  //sats route
	//   r.GET("/stats",controllers.GetBookStats)
	//book routes
	// bookRoutes :=r.Group("/books")
	// {
	// 	bookRoutes.GET("",controllers.GetAllBooks)
	// 	bookRoutes.GET("/:id",controllers.GetBook)
	// 	bookRoutes.POST("",controllers.CreateBook)
	// 	bookRoutes.PUT("/:id",controllers.UpdateBook)
	// 	bookRoutes.DELETE("/:id",controllers.DeleteBook)
	// }

	//protected routes
	protected :=r.Group("/")
	protected.Use(middleware.AuthRequired())
	{
		//books management
		protected.POST("/books",controllers.CreateBook)
		protected.PUT("/books/:id",controllers.UpdateBook)
		protected.DELETE("/books/:id",controllers.DeleteBook)
	
	//author management
	protected.POST("/authors",controllers.CreateAuthor)
	protected.PUT("/authors/:id",controllers.UpdateAuthor)
	protected.DELETE("/authors/:id",controllers.DeleteAuthor)


	// Category management
        protected.POST("/categories", controllers.CreateCategory)
        protected.POST("/categories/:categoryid/books/:bookId", controllers.AddBookToCategory)
        protected.DELETE("/categories/:categoryId/books/:bookId", controllers.RemoveBookFromCategory)
	}
}