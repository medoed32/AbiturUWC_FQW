package main

import (
	"AbiturUWC/internal/controller"
	"AbiturUWC/internal/database"
	"AbiturUWC/internal/model"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	//Conect mySQL DB
	err := database.InitDBmoderator("webinfo:pas@/abitur_uwc")
	if err != nil {
		log.Fatal(err)
	}
	defer database.DBmoderator.Close()

	r := mux.NewRouter()
	routers(r)
	http.Handle("/", r)

	fmt.Println("Stars moderator web is listening...")
	http.ListenAndServe(":8181", nil)
}

var flag bool = false

func routers(r *mux.Router) {
	//Login tamplate
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		LoginPage(w, "")
	}).Methods("GET")
	r.HandleFunc("/login", LoginAutorizate).Methods("POST")

	//Templates user table
	r.HandleFunc("/users", authorized(controller.IndexHandler))
	r.HandleFunc("/users/create", authorized((controller.CreateHandler)))
	r.HandleFunc("/users/edit/{id:[0-9]+}", authorized((func(w http.ResponseWriter, r *http.Request) {
		controller.EditPage(w, r, "")
	}))).Methods("GET")
	r.HandleFunc("/users/edit/{id:[0-9]+}", authorized((controller.EditHandler))).Methods("POST")
	r.HandleFunc("/users/delete/{id:[0-9]+}", authorized((controller.DeleteHandler)))

	//Tamplates user data table
	r.HandleFunc("/usersdata", authorized(controller.IndexHandlerUsersData))
	r.HandleFunc("/usersdata/edit/{id:[0-9]+}", authorized((func(w http.ResponseWriter, r *http.Request) {
		controller.EditPageUsersData(w, r, "")
	}))).Methods("GET")
	r.HandleFunc("/usersdata/edit/{id:[0-9]+}", authorized((controller.EditHandlerUsersData))).Methods("POST")
	r.HandleFunc("/usersdata/delete/{id:[0-9]+}", authorized((controller.DeleteHandlerUsersData)))

	//get static file
	fileServer := http.FileServer(http.Dir("/Users/nikitaepancin/Developer/AbiturUWC/web/static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
}

func LoginPage(rw http.ResponseWriter, message string) {
	flag = false
	tmpl, err := template.ParseFiles("/Users/nikitaepancin/Developer/AbiturUWC/web/html/login.html")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	type answer struct {
		Message string
	}
	data := answer{message}
	err = tmpl.ExecuteTemplate(rw, "login", data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
}

func LoginAutorizate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	user := model.User{}
	user.Login = r.FormValue("login")
	user.Password = r.FormValue("password")
	if user.Login == "" || user.Password == "" {
		LoginPage(w, "Поля не должны быть пустыми")

	} else {
		row := database.DBmoderator.QueryRow("SELECT password FROM users WHERE login = ?", user.Login)
		userdb := model.User{}
		err = row.Scan(&userdb.Password)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(404), http.StatusNotFound)
		} else if user.Password == userdb.Password {
			flag = true
			http.Redirect(w, r, "/users", http.StatusSeeOther)
		} else {
			LoginPage(w, "Не верный логин или пароль")
		}
	}

}

// func return on login template if user not autorized
func authorized(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !flag {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}
