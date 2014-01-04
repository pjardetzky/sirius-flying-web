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

func menuMap(active string) map[string]template.HTMLAttr {
	menu := map[string]template.HTMLAttr{
		"Home":      "",
		"Discovery": "",
		"Training":  "",
		"Rental":    "",
		"About":     "",
		"Contact":   ""}
	menu[active] = template.HTMLAttr("class='active'")
	return menu
}

func index(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/site.html",
		"templates/threecolumn.html",
		"templates/index.html",
	))
	menu := menuMap("Home")
	err := t.Execute(w, menu)
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
	menu := menuMap("Discovery")
	err := t.Execute(w, menu)
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
	menu := menuMap("Training")
	err := t.Execute(w, menu)
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
	menu := menuMap("Rental")
	err := t.Execute(w, menu)
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
	menu := menuMap("About")
	err := t.Execute(w, menu)
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
	menu := menuMap("Contact")
	err := t.Execute(w, menu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
