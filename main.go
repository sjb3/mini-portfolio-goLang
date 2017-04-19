package main

import (
	"html/template"
	"log"
	"net/http"

	_ "github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// mux := httprouter.New()
	// mux.GET("/", index)
	// mux.GET("/contact", contact)
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/resume", resume)
	http.ListenAndServe(":8080", nil)
}

func resume(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	// http.Redirect(w, req, "/resume", http.StatusMovedPermanently)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
