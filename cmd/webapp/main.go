package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/johnull/todo-golang/controllers"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env")
	}

	/* ROUTES */
	router := http.NewServeMux()
	router.HandleFunc("GET /", controllers.GetItems)
	router.HandleFunc("POST /add", controllers.AddItem)
	router.HandleFunc("POST /complete/{id}", controllers.CompleteItem)
	router.HandleFunc("POST /delete/all", controllers.DeleteCompleted)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	} 
}
