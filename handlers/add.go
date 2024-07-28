package handlers

import (
	"github.com/vivalchemy/ecom/components"
	"github.com/vivalchemy/ecom/utils"
	"math/rand"
	"net/http"
)

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	AddPost(w, r)
}

func todoIsCompleted(isCompleted string) bool {
	if isCompleted == "checked" {
		return true
	}
	return false
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	components.Todo(utils.TodoProps{Id: rand.Int(), Title: r.FormValue("add-todo"), IsCompleted: todoIsCompleted(r.FormValue("is-completed"))}).Render(r.Context(), w)
}
