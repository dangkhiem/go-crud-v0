package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	// users := router.Group("/users")
	// {
	// 	users.GET("/", controllers.ListUsers)
	// users.POST("/", controllers.CreateUser)
	// users.PUT("/:id", controllers.UpdateUser)
	// users.DELETE("/:id", controllers.DeleteUser)
	// }

	return router
}
