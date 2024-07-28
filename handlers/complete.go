package handlers

import (
	"github.com/vivalchemy/ecom/components"
	"github.com/vivalchemy/ecom/utils"
	"log"
	"math/rand"
	"net/http"
)

func Completed(w http.ResponseWriter, r *http.Request) {
	log.Println("Logging value")
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	CompletedPost(w, r)
}

func CompletedPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("Logging value of form", r.Form)
	for k, v := range r.Form {
		log.Println(k, v)
	}
	log.Println("Logged value of form")
	components.Todo(utils.TodoProps{Id: rand.Int(), Title: r.FormValue("add-todo"), IsCompleted: true}).Render(r.Context(), w)
}
