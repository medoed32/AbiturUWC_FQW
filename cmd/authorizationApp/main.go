package main

import (
	"log"
	"net/http"

	"github.com/medoed32/AbiturUWC_FQW/internal/app/authorizationApp/controller"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	routes(r)

	err := http.ListenAndServe("localhost:4444", r)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router) {
	r.ServeFiles("/web/*filepath", http.Dir("public"))
	r.GET("/", controller.StartPage)
	r.GET("/users", controller.GetUsers)
}
