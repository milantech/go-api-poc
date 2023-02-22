package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-api-poc/controllers"
	"github.com/go-api-poc/models"
)

func Routes() {
	route := gin.Default()
	//	gin.SetMode(gin.ReleaseMode)
	models.ConnectDatabase()
	//route.POST("/todo", controllers.CreateTodo)
	//route.POST("/createroles", controllers.CreateRoles)
	route.POST("/createnewroles", controllers.CreateNewRoles)
	route.PUT("/updateroles/:id", controllers.UpdateNewRoles)
	//route.POST("/createpost", controllers.CreatePost) //CreateNewRoles
	//route.POST("/posts", controllers.CreatePost)
	//	route.GET("/todo", controllers.GetAllTodos)
	route.PUT("/todo/:idTodo", controllers.UpdateTodo)
	//route.DELETE("/todo/:idTodo", controllers.DeleteTodo)

	route.Run(":9999")
}
