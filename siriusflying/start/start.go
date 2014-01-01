package start

import (
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/discovery", discovery)
	http.HandleFunc("/training", training)
	http.HandleFunc("/rental", rental)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
}


func index(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/site.html",
		"templates/threecolumn.html",
		"templates/index.html",
	))
	err := t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func discovery(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/site.html",
		"templates/threecolumn.html",
		"templates/discovery.html",
	))
	err := t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func training(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/site.html",
		"templates/threecolumn.html",
		"templates/training.html",
	))
	err := t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func rental(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/site.html",
		"templates/threecolumn.html",
		"templates/rental.html",
	))
	err := t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/site.html",
		"templates/onecolumn.html",
		"templates/about.html",
	))
	err := t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/site.html",
		"templates/threecolumn.html",
		"templates/contact.html",
	))
	err := t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
