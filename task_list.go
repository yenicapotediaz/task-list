package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//TodoItem struct
type TodoItem struct {
	ID      int
	Caption string
	Done    bool
}

//todo
var todo []TodoItem

//assign global db variable
var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root@tcp(a9d7d3790003711e790670a9b77d070e-319954018.us-west-2.elb.amazonaws.com:3306)/list")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
}

//output
func main() {
	fmt.Println("initializing...")
	todo = make([]TodoItem, 0)

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var task = TodoItem{}
		var done int
		err := rows.Scan(&task.ID, &task.Caption, &done)
		if done == 0 {
			task.Done = false
		} else {
			task.Done = true
		}

		if err != nil {
			log.Fatal(err)
		}
		todo = append(todo, task)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	for _, task := range todo {
		log.Println("here's caption:", task.Caption)
	}

	runServer()
}

func runServer() {
	fmt.Println("Running server...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	newItem := r.URL.Query().Get("caption")
	if newItem != "" {
		fmt.Println("caption:", newItem)
		newTask := TodoItem{Caption: newItem, Done: false}

		//add newItem todo into DB
		_, err := db.Exec("INSERT INTO todos (ID, Caption, Done) VALUES(?,?,?)", 0, newItem, 0)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		log.Println("caption:", newItem, "added")
		todo = append(todo, newTask)
	}

	t, err := template.ParseFiles("/templates/index.html")
	if err != nil {
		log.Fatal("can not parse /templates/index.html: " + err.Error())
	}
	t.Execute(w, todo)
}
