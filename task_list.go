package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type TodoItem struct {
	Id         int
	Caption    string
	IsFinished bool
}

var TodoItemsSlice = []TodoItem{}

func main() {
	fmt.Println("initializing...")
	TodoItemsSlice = make([]TodoItem, 0)
	runServer()
}

func runServer() {
	fmt.Println("Running server...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./views/main.gtpl")
	t.Execute(w, TodoItemsSlice)
}
