package main

import (
	"html/template"
	"log"
	"net/http"
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
	fileServer := http.FileServer(http.Dir("./Static"))
	http.Handle("/resources/", http.StripPrefix("/resources", fileServer))
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(Host+":"+Port, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server :", err)
		return
	}
}
