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



	  //public routes
	 //health check route
      r.GET("/health",controllers.HealthCheck)
	  r.GET("/books",controllers.GetAllBooks)
	  r.GET("/books/:id",controllers.GetBook)
	  //sats route
	  r.GET("/stats",controllers.GetBookStats)
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
		protected.POST("/books",controllers.CreateBook)
		protected.PUT("/books/:id",controllers.UpdateBook)
		protected.DELETE("/books/:id",controllers.DeleteBook)
	}
}