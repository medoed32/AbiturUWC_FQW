package main

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

func main() {

	//Conect mySQL DB
	err := database.InitDBmoderator("webtest:pas@/abitur_uwc")
	if err != nil {
		log.Fatal(err)
	}
	defer database.DBmoderator.Close()

	r := mux.NewRouter()
	routers(r)
	http.Handle("/", r)

	fmt.Println("Stars WebInfo web is listening...")
	http.ListenAndServe(":8182", nil)

}

var chekRes bool = false
var resultTestDrive = ""

func routers(r *mux.Router) {
	r.HandleFunc("/testdrive", func(w http.ResponseWriter, r *http.Request) {
		testDirivePage(w, "")
	}).Methods("GET")
	r.HandleFunc("/testdrive", func(w http.ResponseWriter, r *http.Request) {
		testDiriveHandler(w, r)
	}).Methods("POST")
	r.HandleFunc("/testdrive/result", chekResult(func(w http.ResponseWriter, r *http.Request) {
		ResHandle(w, resultTestDrive)
	})).Methods("GET")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/Users/nikitaepancin/Developer/AbiturUWC/web/html/home.html")
	})
	r.HandleFunc("/faqs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/Users/nikitaepancin/Developer/AbiturUWC/web/html/faqs.html")
	})
	r.HandleFunc("/information", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/Users/nikitaepancin/Developer/AbiturUWC/web/html/information.html")
	})

	//get static file
	fileServer := http.FileServer(http.Dir("/Users/nikitaepancin/Developer/AbiturUWC/web/static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

}
func testDirivePage(w http.ResponseWriter, message string) {
	type answer struct {
		Message string
	}
	data := answer{message}
	tmpl, _ := template.ParseFiles("/Users/nikitaepancin/Developer/AbiturUWC/web/html/testdrive.html")
	err := tmpl.ExecuteTemplate(w, "testDriveMessage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func testDiriveHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	userData := model.UserData{}
	userData.Patronymic = r.FormValue("patronymic")
	if userData.Patronymic == "" && r.FormValue("patronymicCheck1") != "on" {
		testDirivePage(w, "Укажите все поля. Если у вас нет отчества, то поставьте отметку")
		return
	}
	userData.FirstName = r.FormValue("firstName")
	userData.LastName = r.FormValue("lastName")
	userData.Phone = r.FormValue("phone")
	userData.City = r.FormValue("city")
	userData.Email = r.FormValue("email")
	if userData.FirstName == "" || userData.LastName == "" || userData.Phone == "" || userData.City == "" || userData.Email == "" {
		testDirivePage(w, "Укажите все поля. Если у вас нет отчества, то поставьте отметку")
		return
	}

	resTest := map[string][][]int{
		"Политология":   {{1, 5, 9, 17, 24, 25}, {0, 1}},
		"Психология":    {{4, 6, 8, 11, 17}, {0, 2}},
		"Юриспруденция": {{}, {0, 3}},
		"Реклама и связи с общественностью": {{}, {0, 4}},
		"Зарубежное регионоведение":         {{}, {0, 5}},
		"Международные отношения":           {{}, {0, 6}},
		"Экономика":  {{}, {0, 7}},
		"Менеджмент": {{}, {0, 8}},
		"Государственное и муниципальное управление": {{}, {0, 9}},
		"Бизнес-информатика":                         {{}, {0, 10}},
		"Дизайн":                                     {{}, {0, 11}},
		"Лингвистика":                                {{}, {0, 12}},
		"Журналистика":                               {{}, {0, 13}},
	}
	for key, value := range resTest {
		for _, v := range value[0] {
			questionRadio := r.FormValue("questionRadio" + strconv.Itoa(v))
			if questionRadio == "yes" {
				resTest[key][1][0] += 1
			} else {

			}
		}
	}
	kMax, vMax, idMax := "", 0, 1
	for k, v := range resTest {
		if v[1][0] > vMax {
			kMax, vMax, idMax = k, v[1][0], v[1][1]
		}
	}

	_, err = database.DBmoderator.Exec("INSERT INTO users_data (first_name, last_name, patronymic, phone, city, email, direction_id) values (?, ?, ?, ?, ?, ?, ?)",
		userData.FirstName, userData.LastName, userData.Patronymic, userData.Phone, userData.City, userData.Email, idMax)
	if err != nil {
		log.Println(err)
	}
	resultTestDrive = kMax
	chekRes = true
	http.Redirect(w, r, "/testdrive/result", 301)
}

func ResHandle(w http.ResponseWriter, result string) {
	type resultTest struct {
		ResultDirectionName string
	}
	data := resultTest{result}
	tmpl, _ := template.ParseFiles("/Users/nikitaepancin/Developer/AbiturUWC/web/html/results.html")
	err := tmpl.ExecuteTemplate(w, "testDriveResult", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func chekResult(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !chekRes {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}
