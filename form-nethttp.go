package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// import (
// 	"github.com/Basillica/golang-guide/handler/api"
// )

// const Version = "v1"

// func main() {

// 	api.New(Version)
// }

type FormData struct {
	Name    string
	Email   string
	Message string
}

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func WriteToDb(data FormData) {
	fmt.Println(data, "the data has been written to the database")
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	req := FormData{
		Name:    name,
		Email:   email,
		Message: message,
	}
	WriteToDb(req)

	fmt.Fprintf(w, "Form submitted successfully!")
}

func main2() {
	http.HandleFunc("/", RenderTemplate)
	http.HandleFunc("/submit", SubmitHandler)
	fmt.Println("server is being served on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
