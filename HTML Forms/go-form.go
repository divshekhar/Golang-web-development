package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("Template/index.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "Email : ", r.PostForm.Get("email"))
	fmt.Fprintln(w, "Password : ", r.PostForm.Get("password"))
	fmt.Fprintln(w, "Remember Me : ", r.PostForm.Get("remember_check"))
}

func main() {
	http.HandleFunc("/", renderTemplate)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server : ", err)
		return
	}
}
