package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/johnull/todo-golang/internal/controllers"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env")
	}

	/* ROUTES */
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.GetItems)
	router.HandleFunc("/add", controllers.AddItem).Methods("POST")
	router.HandleFunc("/complete/{id}", controllers.CompleteItem).Methods("POST")
	router.HandleFunc("/delete/all", controllers.DeleteCompleted).Methods("POST")

	// http server config
	server := &http.Server{
		Addr: ":8080",
		Handler: router,
	}
	
	// starts the server in a goroutine to avoid blocking execution
	go func() {

		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("Server stopped")
	}()

	// channel to catch SIGINT & SIGTERM from the OS
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// waits interrupt signals 
	<-sig

	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// wait for the graceful shutdown of the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP close error: %v", err)
	}
	log.Println("Server shutdown.")
}
