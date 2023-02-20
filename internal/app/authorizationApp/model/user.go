package model

import (
	"github.com/medoed32/AbiturUWC_FQW/internal/app/authorizationApp/server"
)

type User struct {
	Id           int
	Login        string
	Password     string
	Role_id      int
	User_data_id int
}

func GetAllUsers() (users []User, err error) {
	rows, err := server.Db.Queryx(`SELECT id, login, password FROM users`)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	user := User{}
	for rows.Next() {
		err = rows.StructScan(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, err
}

func NewUser(login, password string, roleId int) *User {
	return &User{Login: login, Password: password, Role_id: roleId}
}

func GetUserById(userId string) (u User, err error) {
	query := `SELECT * FROM users WHERE id = ?`
	err = server.Db.Get(&u, query, userId)
	return
}

func (u *User) AddUser() (err error) {
	query := `INSERT INTO users(login, password, role_id) VALUES (?, ?, 3)`
	_, err = server.Db.Exec(query, u.Login, u.Password)
	return
}

func (u *User) DeleteUser() (err error) {
	query := `DELETE FROM users WHERE id = ?`
	_, err = server.Db.Exec(query, u.Id)
	return
}
