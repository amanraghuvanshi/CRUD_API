package routes

import (
	"assignTele/controllers"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRoutes(BlogControllers *controllers.BlogControllers) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome")
	})

	router.GET("/blogs", BlogControllers.Find)
	router.GET("/blogs/:id", BlogControllers.FindbyID)
	router.DELETE("blogs/:id", BlogControllers.Delete)
	router.POST("/blogs", BlogControllers.Create)
	router.PATCH("/blogs/:id", BlogControllers.Update)

	return router
}
