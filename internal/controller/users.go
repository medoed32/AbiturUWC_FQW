package controller

import (
	"AbiturUWC/internal/database"
	"AbiturUWC/internal/model"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idNum, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}
	_, err = database.DBmoderator.Exec("UPDATE users SET delete_check = ? WHERE id = ?", 1, idNum)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/users", 301)
}

func EditPage(w http.ResponseWriter, r *http.Request, message string) {
	vars := mux.Vars(r)
	id := vars["id"]

	row := database.DBmoderator.QueryRow("SELECT id, login, password FROM users WHERE id = ?", id)
	user := model.User{}
	err := row.Scan(&user.Id, &user.Login, &user.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		type answer struct {
			Data    model.User
			Message string
		}
		data := answer{user, message}
		tmpl, _ := template.ParseFiles("/Users/nikitaepancin/Developer/AbiturUWC/web/html/edit.html")
		err = tmpl.ExecuteTemplate(w, "editMessage", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	user := model.User{}
	user.Id, err = strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println(err)
	}
	user.Login = r.FormValue("login")
	user.Password = r.FormValue("password")
	if user.Login == "" || user.Password == "" {
		EditPage(w, r, "Поля не должны быть пустыми")
	} else {
		_, err = database.DBmoderator.Exec("UPDATE users SET login = ?, password = ? WHERE id = ?",
			user.Login, user.Password, user.Id)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/users", 301)
	}

}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		login := r.FormValue("login")
		password := r.FormValue("password")

		_, err = database.DBmoderator.Exec("INSERT INTO users (login, password, role_id) values (?, ?, 1)",
			login, password)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/users", 301)
	} else {
		http.ServeFile(w, r, "/Users/nikitaepancin/Developer/AbiturUWC/web/html/create.html")
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DBmoderator.Query("SELECT id, login, password FROM users WHERE delete_check = ?", 0)
	if err != nil {
		log.Println(err)
	}

	user := []model.User{}

	for rows.Next() {
		u := model.User{}
		err := rows.Scan(&u.Id, &u.Login, &u.Password)
		if err != nil {
			fmt.Println(err)
			continue
		}
		user = append(user, u)
	}

	tmpl, _ := template.ParseFiles("/Users/nikitaepancin/Developer/AbiturUWC/web/html/index.html")
	err = tmpl.ExecuteTemplate(w, "users", user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
