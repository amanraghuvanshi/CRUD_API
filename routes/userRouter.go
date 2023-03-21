package routes

import (
	"blogAPI/controllers"
	controller "blogAPI/controllers"
	"blogAPI/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/blogs", controller.GetPost())
	incomingRoutes.POST("/blogs", controller.CreatePost())
	incomingRoutes.GET("/blogs/:id", controllers.GetOneBlog())
	incomingRoutes.DELETE("/blogs/:id", controllers.DeleteOne())
	incomingRoutes.PATCH("/blogs:id", controller.UpdateBlogs())
}
