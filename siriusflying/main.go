package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/discovery", discovery)
	mux.HandleFunc("/training", training)
	mux.HandleFunc("/rental", rental)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contact", contact)

	css := http.FileServer(http.Dir("./css/"))
	fonts := http.FileServer(http.Dir("./fonts/"))
	js := http.FileServer(http.Dir("./js/"))
	img := http.FileServer(http.Dir("./img/"))

	mux.Handle("/css/", http.StripPrefix("/css", css))
	mux.Handle("/fonts/", http.StripPrefix("/fonts", fonts))
	mux.Handle("/js/", http.StripPrefix("/js", js))
	mux.Handle("/img/", http.StripPrefix("/img", img))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
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

func discovery(w http.ResponseWriter, r *http.Request) {
	menu := menuMap("Discovery")
	t := goTemplate("discovery", "three")
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
