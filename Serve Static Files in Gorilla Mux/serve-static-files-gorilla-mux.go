package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	Host = "localhost"
	Port = "8080"
)

type Student struct {
	Name       string
	College    string
	RollNumber int
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	student := Student{
		Name:       "GB",
		College:    "GolangBlogs",
		RollNumber: 1,
	}
	parsedTemplate, _ := template.ParseFiles("Template/index.html")
	err := parsedTemplate.Execute(w, student)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", renderTemplate).Methods("GET")
	fileServer := http.FileServer(http.Dir("./Static"))
	router.PathPrefix("/").Handler(http.StripPrefix("/resources", fileServer))
	err := http.ListenAndServe(Host+":"+Port, router)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server :", err)
		return
	}
}
