package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/vivalchemy/ecom/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("PORT"))

	log.Println("Starting server...")

	// serving the static assets
	handlers.HandleRoutes()

	// assigning the port
	if PORT, ok := os.LookupEnv("PORT"); !ok {
		log.Fatal(http.ListenAndServe(":3000", nil))
	} else {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
	}

}
