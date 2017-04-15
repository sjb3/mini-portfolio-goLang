package main

import (
	"html/template"
	"net/http"
	_"github.com/julienschmidt/httprouter"
	"log"
)

var tpl *template.Template

func init() {
	// tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
	tpl = template.Must(template.ParseGlob("templates/*"))

}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
	// mux := httprouter.New()
	// mux.GET("/", index)
	// http.ListenAndServe(":8080", mux)
}

func contact(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func index(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	// tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
