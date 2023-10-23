package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ivanrafli14/API-BTPN/controllers"
	"github.com/ivanrafli14/API-BTPN/database"
	"github.com/ivanrafli14/API-BTPN/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	database.ConnectDatabase()

	r.POST("/users/register", controllers.Register)
	r.POST("/users/login", controllers.Login)
	r.POST("/users/logout", controllers.Logout)

	authorized := r.Group("/")
	authorized.Use(middleware.CheckAuth())
	{
		authorized.GET("/users/:id", controllers.GetUserbyID)
		authorized.POST("/users/:id", controllers.UpdateUser)
		authorized.DELETE("/users/:id", controllers.DeleteUser)

		authorized.GET("/photos", controllers.GetAllPhotos)
		authorized.POST("/photos", controllers.CreatePhoto)
		authorized.PUT("/photos/:id", controllers.UpdatePhoto)
		authorized.DELETE("/photos/:id", controllers.DeletePhoto)

	}

	 return r


}