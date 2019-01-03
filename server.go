package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	err := t.Execute(w, nil)
	if err != nil {
		return
	}
}

func RunWebServer() {
	routes := mux.NewRouter()
	routes.HandleFunc("/index", indexHandler)

	http.Handle("/", routes)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe(":8593", nil)
	if err != nil {
		fmt.Println("It's Server.RunWebServer() we got problem")
	}
}
