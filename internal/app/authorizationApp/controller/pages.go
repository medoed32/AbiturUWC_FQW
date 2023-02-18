package controller

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func StartPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	text := "Приветсвую тебя на стартовой странице этого сайта"
	fmt.Fprint(w, text)
}
