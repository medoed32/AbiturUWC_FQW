package controller

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/medoed32/AbiturUWC_FQW/internal/app/authorizationApp/model"

	"github.com/julienschmidt/httprouter"
)

func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем список всех пользователей
	users, err := model.GetAllUsers()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	//указываем путь к файлу с шаблоном
	main := filepath.Join("/Users/nikitaepancin/Developer/AbiturUWC_FQW/web/account_users.html")
	//создаем html-шаблон
	tmpl, err := template.ParseFiles(main)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	//исполняем именованный шаблон "users", передавая туда массив со списком пользователей
	err = tmpl.ExecuteTemplate(rw, "users", users)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
