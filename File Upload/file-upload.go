package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	// Host of the HTTP Server
	Host = "localhost"
	// Port of the HTTP Server
	Port = "8080"
)

var parsedTemplate *template.Template

func init() {
	parsedTemplate, _ = template.ParseFiles("template/index.html")
}

func index(w http.ResponseWriter, r *http.Request) {
	parsedTemplate.Execute(w, nil)
}

func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Println("Error Getting File", err)
		return
	}
	defer file.Close()

	out, pathError := ioutil.TempFile("temp-images", "upload-*.png")
	if pathError != nil {
		log.Println("Error Creating a file for writing", pathError)
		return
	}
	defer out.Close()

	_, copyError := io.Copy(out, file)
	if copyError != nil {
		log.Println("Error copying", copyError)
	}
	fmt.Fprintln(w, "File Uploaded Successfully! ")
	fmt.Fprintln(w, "Name of the File: ", header.Filename)
	fmt.Fprintln(w, "Size of the File: ", header.Size)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", fileUploadHandler)
	err := http.ListenAndServe(Host+":"+Port, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server : ", err)
	}
}
