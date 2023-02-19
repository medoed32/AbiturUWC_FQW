package model

type User struct {
	Id         int
	Login      string
	FirstName  string
	LastName   string
	Patronymic string
	Phone      string
	City       string
	Email      string
	Role       string
}

func GetAllUsers() (users []User, err error) {
	users = []User{
		{1, "test", "Джон", "До", "Konstantinovich", "79998764643", "Moscow", "Email", "Admin"},
		{2, "test2", "Tera", "dida", "Konstantinovich", "79458763343", "SPB", "test@test.ru", "Moderator"},
	}
	return
}
