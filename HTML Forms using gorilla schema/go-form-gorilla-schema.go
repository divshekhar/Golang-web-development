package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

// User Struct
type User struct {
	Email         string
	Password      string
	RememberCheck string
}

func readForm(r *http.Request) *User {
	r.ParseForm()
	user := new(User)
	decoder := schema.NewDecoder()
	decodeErr := decoder.Decode(user, r.PostForm)
	if decodeErr != nil {
		log.Println("Error mapping parsed form data to struct : ", decodeErr)
	}
	return user

}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parsedTemplate, _ := template.ParseFiles("Template/index.html")
		err := parsedTemplate.Execute(w, nil)
		if err != nil {
			log.Println("Error executing template :", err)
			return
		}
	} else {
		user := readForm(r)
		fmt.Fprintln(w, "Email : ", user.Email)
		fmt.Fprintln(w, "Password : ", user.Password)
		fmt.Fprintln(w, "Remember Me : ", user.RememberCheck)
	}

}

func main() {
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server : ", err)
		return
	}
}
