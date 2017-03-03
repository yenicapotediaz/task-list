package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//TodoItem struct
type TodoItem struct {
	ID         int
	Caption    string
	IsFinished bool
}

//TodoItemsSlice todo object
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
	captionFormValue := r.PostFormValue("caption")
	if captionFormValue != "" {
		fmt.Println("caption:", captionFormValue)
		newID := len(TodoItemsSlice) + 1
		p := &TodoItem{ID: newID, Caption: captionFormValue, IsFinished: false}
		TodoItemsSlice = append(TodoItemsSlice, *p)
	}

	t, err := template.ParseFiles("/templates/index.html")
	if err != nil {
		log.Fatal("can not parse templates/index.html: " + err.Error())
	}
	t.Execute(w, TodoItemsSlice)
}
