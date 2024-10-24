package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/johnull/todo-golang/internal/database"
	"github.com/johnull/todo-golang/internal/models"
	"github.com/joho/godotenv"
)

var (
	db = database.ConnectDB()
	view *template.Template
)

func init() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal(err)
	}
	viewPath, ok := os.LookupEnv("VIEW_PATH")
	if !ok {
		log.Fatal("VIEW_PATH not set")
	}
	
	view = template.Must(template.ParseFiles(viewPath))
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM TodoList`)
	if err != nil {
		log.Printf("database get items error: %v", err)
		return
	}
	defer rows.Close()

	var todos []models.Todo
	
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.Id, &todo.Item, &todo.Completed); err != nil {
			log.Printf("database reading rows error: %v", err)
			return
		}
		todos = append(todos, todo)
	}
	
	data := models.View{Todos: todos}
	
	if err := view.Execute(w, data); err != nil {
		log.Printf("template execute error: %v", err)
		return
	}
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("current-item")
	if item == "" {
		http.Error(w, "Empty value", http.StatusBadRequest)
		return
	}
	if _, err := db.Exec(`INSERT INTO TodoList (item) VALUE (?)`, item); err != nil {
		log.Printf("database add item error: %v", err)
		return
	}
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func CompleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if _, err := db.Exec(`UPDATE TodoList SET completed = 1 WHERE id = (?)`, id); err != nil {
		log.Printf("database update item error: %v", err)
		return
	}
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeleteCompleted(w http.ResponseWriter, r *http.Request) {
	if _, err := db.Exec(`DELETE FROM TodoList WHERE completed = 1`); err != nil {
		log.Printf("database delete item error: %v", err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
