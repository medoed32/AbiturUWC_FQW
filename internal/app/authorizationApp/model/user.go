package model

import "net/http"

type User struct {
	Id      int
	Name    string
	Surname string
}

func GetAllUsers() (users []User, err, error) {
	users = []User{
		{1, "Джон", "До"},
		{2, "Jora", "Tohan"},
	}
	return
}
