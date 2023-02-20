package model

import (
	"github.com/medoed32/AbiturUWC_FQW/internal/app/authorizationApp/server"
)

type UserData struct {
	Id           int
	First_name   string
	Last_name    string
	Patronymic   string
	Phone        string
	City         string
	Direction_id int
}

func GetAllUserData() (userData []UserData, err error) {
	query := `SELECT * FROM users_data`
	err = server.Db.Select(&userData, query)
	return
}

func NewUserData(firstName, lastName, patronymic, phone, city string, directionId int) *UserData {
	return &UserData{First_name: firstName, Last_name: lastName, Patronymic: patronymic, Phone: phone, City: city, Direction_id: directionId}
}

func GetUserDataById(userDataId string) (u UserData, err error) {
	query := `SELECT * FROM users_data WHERE id = ?`
	err = server.Db.Get(&u, query, userDataId)
	return
}

func (u *UserData) AddUserData() (err error) {
	query := `INSERT INTO users_data(first_name, last_name, patronymic, phone, city, Direction_id) VALUES (?, ?, ?, ?, ?, ?)`
	_, err = server.Db.Exec(query, u.First_name, u.Last_name, u.Patronymic, u.Phone, u.City, u.Direction_id)
	return
}
