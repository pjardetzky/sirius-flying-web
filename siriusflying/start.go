package start

import (
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/training", training)
	http.HandleFunc("/rental", rental)
	http.HandleFunc("/commercial", commercial)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
}

func menuMap(active string) map[string]template.HTMLAttr {
	menu := map[string]template.HTMLAttr{
		"Home":       "",
		"Training":   "",
		"Rental":     "",
		"Commercial": "",
		"About":      "",
		"Contact":    ""}
	menu[active] = template.HTMLAttr("class='active'")
	return menu
}

func goTemplate(page, cols string) *template.Template {
	return template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/site.html",
		"templates/"+cols+"column.html",
		"templates/"+page+".html",
	))
}

func index(w http.ResponseWriter, r *http.Request) {
	menu := menuMap("Home")
	t := goTemplate("index", "three")
	err := t.Execute(w, menu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func training(w http.ResponseWriter, r *http.Request) {
	menu := menuMap("Training")
	t := goTemplate("training", "three")
	err := t.Execute(w, menu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func rental(w http.ResponseWriter, r *http.Request) {
	menu := menuMap("Rental")
	t := goTemplate("rental", "three")
	err := t.Execute(w, menu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func commercial(w http.ResponseWriter, r *http.Request) {
	menu := menuMap("Commercial")
	t := goTemplate("commercial", "three")
	err := t.Execute(w, menu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	menu := menuMap("About")
	t := goTemplate("about", "one")
	err := t.Execute(w, menu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	menu := menuMap("Contact")
	t := goTemplate("contact", "three")
	err := t.Execute(w, menu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
