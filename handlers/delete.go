package handlers

import (
	"log"
	"net/http"
	"regexp"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if notValid := regexp.MustCompile(`^/delete/todo-(\d+)$`).FindStringSubmatch(r.URL.Path); notValid == nil {
		log.Println("Received a delete req")
		// 404 not found
		http.NotFound(w, r)
	}
	// returns nothing
	w.Write([]byte(""))
}
