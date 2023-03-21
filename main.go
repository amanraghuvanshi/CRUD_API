package main

import (
	"assignTele/helper"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	routes := httprouter.New()
	routes.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home!")
	})

	server := http.Server{Addr: "localhost:3333", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
