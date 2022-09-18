package handlers

import (
	"html/template"
	"log"
	"net/http"
)

//Registering Login Page
func Login(w http.ResponseWriter, _ *http.Request) {
	tpl, err := template.ParseFiles("view/login.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(w, nil)
}

//Rendering registration page
func Register(w http.ResponseWriter, _ *http.Request) {
	tpl, err := template.ParseFiles("view/register.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(w, nil)
}
