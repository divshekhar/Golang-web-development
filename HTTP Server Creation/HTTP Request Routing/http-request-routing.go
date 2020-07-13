package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	// Host name of the HTTP Server
	Host = "localhost"
	// Port of the HTTP Server
	Port = "8080"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HOME Page")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ABOUT Page")
}

func services(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SERVICES Page")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/services", services)
	err := http.ListenAndServe(Host+":"+Port, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server : ", err)
		return
	}

}
