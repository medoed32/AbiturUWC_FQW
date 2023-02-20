package main

import (
	"log"
	"net/http"

	"github.com/medoed32/AbiturUWC_FQW/internal/app/authorizationApp/controller"
	"github.com/medoed32/AbiturUWC_FQW/internal/app/authorizationApp/server"

	"github.com/julienschmidt/httprouter"
)

func main() {
	//инициализируем подключение к базе данных
	err := server.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	//создаем и запускаем в работу роутер для обслуживания запросов
	r := httprouter.New()
	routes(r)
	//прикрепляемся хосту и порту для приема и обслуживания входящих запросов
	//вторым параметром передается роутер, который будет работать с запросами
	err = http.ListenAndServe("localhost:4444", r)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router) {
	r.ServeFiles("/web/*filepath", http.Dir("public"))
	r.GET("/", controller.StartPage)
	r.GET("/users", controller.GetUsers)
	r.POST("/user/add", controller.AddUsers)
}
