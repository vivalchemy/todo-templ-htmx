package handlers

import (
	"net/http"
)

func HandleRoutes() {
	// the ./assets/ folder is mapped to the /static/ route
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))

	http.HandleFunc("/", Home)
	http.HandleFunc("/info", Info)
	http.HandleFunc("/add", Add)
	http.HandleFunc("/delete/", Delete)
	http.HandleFunc("/completed", Completed)
}
