package main

import (
	"assignTele/config"
	"assignTele/controllers"
	"assignTele/helper"
	"assignTele/routes"
	"assignTele/service"
	"assignTele/utility"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("Starting Server!")

	db := config.DatabaseConnection()
	// database
	blogRepo := utility.NewBlogsRepo(db)
	//service
	blogService := service.NewBlogServiceImpl(blogRepo)
	// controller
	BlogControllers := controllers.NewBlogsControl(blogService)

	// routes
	routes := routes.NewRoutes(BlogControllers)
	routes.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home!")
	})

	server := http.Server{Addr: "localhost:3333", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
