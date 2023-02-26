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

func DeleteHandlerUsersData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idNum, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}
	_, err = database.DBmoderator.Exec("UPDATE users_data SET delete_check = ? WHERE id = ?", 1, idNum)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/usersdata", 301)
}

func EditPageUsersData(w http.ResponseWriter, r *http.Request, message string) {
	vars := mux.Vars(r)
	id := vars["id"]

	row := database.DBmoderator.QueryRow("SELECT id, first_name, last_name, patronymic, phone, city, email, direction_id  FROM users_data WHERE id = ?", id)
	userData := model.UserData{}
	err := row.Scan(&userData.Id, &userData.FirstName, &userData.LastName, &userData.Patronymic, &userData.Phone, &userData.City, &userData.Email, &userData.DirectionId)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		type answer struct {
			Data           model.UserData
			Message        string
			DirectionName  [][]string
			DirectionIdStr string
		}
		rowsDirections, err := database.DBmoderator.Query("SELECT id, direction FROM directions")
		if err != nil {
			log.Println(err)
		}
		dirs := make([][]string, 0)
		for rowsDirections.Next() {
			dir := make([]string, 2, 2)
			err = rowsDirections.Scan(&dir[0], &dir[1])
			if err != nil {
				log.Println(err)
				continue
			}
			dirs = append(dirs, dir)

		}
		strid := strconv.Itoa(userData.DirectionId)
		data := answer{userData, message, dirs, strid}
		tmpl, _ := template.ParseFiles("/Users/nikitaepancin/Developer/AbiturUWC/web/html/editUsersData.html")
		err = tmpl.ExecuteTemplate(w, "editUserDataMessage", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func EditHandlerUsersData(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	userData := model.UserData{}
	userData.Id, err = strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println(err)
	}
	userData.FirstName = r.FormValue("firstName")
	userData.LastName = r.FormValue("lastName")
	userData.Patronymic = r.FormValue("patronymic")
	userData.Phone = r.FormValue("phone")
	userData.City = r.FormValue("city")
	userData.Email = r.FormValue("email")
	userData.DirectionId, err = strconv.Atoi(r.FormValue("contact"))

	if userData.FirstName == "" || userData.LastName == "" || userData.Patronymic == "" || userData.Phone == "" || userData.City == "" || userData.Email == "" {
		EditPageUsersData(w, r, "Поля не должны быть пустыми")
	} else {
		_, err = database.DBmoderator.Exec("UPDATE users_data SET first_name = ?, last_name = ?, patronymic = ?, phone = ?, city = ?, email = ?, direction_id = ? WHERE id = ?",
			userData.FirstName, userData.LastName, userData.Patronymic, userData.Phone, userData.City, userData.Email, userData.DirectionId, userData.Id)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/usersdata", 301)
	}

}

func CreateHandlerUsersData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		userData := model.UserData{}
		userData.FirstName = r.FormValue("firstName")
		userData.LastName = r.FormValue("lastName")
		userData.Patronymic = r.FormValue("patronymic")
		userData.Phone = r.FormValue("phone")
		userData.City = r.FormValue("city")
		userData.Email = r.FormValue("email")
		userData.DirectionId, err = strconv.Atoi(r.FormValue("directionId"))
		_, err = database.DBmoderator.Exec("INSERT INTO users_data (first_name, last_name, patronymic, phone, city, email, direction_id) values (?, ?, ?, ?, ?, ?, ?)",
			userData.FirstName, userData.LastName, userData.Patronymic, userData.Phone, userData.City, userData.Email, r.FormValue("directionId"))

		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/usersdata", 301)
	} else {
		http.ServeFile(w, r, "/Users/nikitaepancin/Developer/AbiturUWC/web/html/create.html")
	}
}

func IndexHandlerUsersData(w http.ResponseWriter, r *http.Request) {
	rowsUserData, err := database.DBmoderator.Query("SELECT id, first_name, last_name, patronymic, phone, city, email, direction_id FROM users_data WHERE delete_check = ?", 0)
	if err != nil {
		log.Println(err)
	}

	type UserDataDirection struct {
		Data          model.UserData
		DirectionName string
	}
	userData := []UserDataDirection{}
	for rowsUserData.Next() {
		uD := UserDataDirection{}
		err := rowsUserData.Scan(&uD.Data.Id, &uD.Data.FirstName, &uD.Data.LastName, &uD.Data.Patronymic, &uD.Data.Phone, &uD.Data.City, &uD.Data.Email, &uD.Data.DirectionId)
		if err != nil {
			fmt.Println(err)
			continue
		}

		rowsDirection := database.DBmoderator.QueryRow("SELECT direction FROM directions WHERE id = ?", uD.Data.DirectionId)
		err = rowsDirection.Scan(&uD.DirectionName)
		if err != nil {
			fmt.Println(err)
			continue
		}

		userData = append(userData, uD)
	}

	tmpl, _ := template.ParseFiles("/Users/nikitaepancin/Developer/AbiturUWC/web/html/indexUsersData.html")
	err = tmpl.ExecuteTemplate(w, "usersData", userData)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
