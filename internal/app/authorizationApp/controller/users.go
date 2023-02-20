package controller

import (
	"encoding/json"
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

func AddUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем значение из параметра login, переданного в форме запроса
	login := r.FormValue("login")
	//получаем значение из параметра password, переданного в форме запроса
	password := r.FormValue("password")

	//проверяем на пустые значения
	if login == "" || password == "" {
		http.Error(rw, "Поля не могут быть пустыми", 400)
		return
	}
	//создаем новый объект
	user := model.NewUser(login, password, 3)
	//записываем нового пользователя в таблицу БД
	err := user.AddUser()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//возвращаем текстовое подтверждение об успешном выполнении операции
	err = json.NewEncoder(rw).Encode("Пользователь успешно добавлен!")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
