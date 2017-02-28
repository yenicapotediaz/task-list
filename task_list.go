package main

import (
	"fmt"
	"html/template"
	"log"
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
	http.HandleFunc("/additem", addItemHandler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	captionFormValue := r.PostFormValue("caption")
	if captionFormValue != "" {
		fmt.Println("caption:", captionFormValue)
		newId := len(TodoItemsSlice) + 1
		p := &TodoItem{Id: newId, Caption: captionFormValue, IsFinished: false}
		TodoItemsSlice = append(TodoItemsSlice, *p)
	}

	t, err := template.ParseFiles("./views/main.gtpl")
	if err != nil {
		log.Fatal("can not parse views/main.gtpl " + err.Error())
	}
	t.Execute(w, TodoItemsSlice)
}

func addItemHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("addItemHandler")
	t, _ := template.ParseFiles("./views/additem.gtpl")
	t.Execute(w, t)
}
