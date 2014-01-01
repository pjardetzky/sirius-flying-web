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

type Menu struct {
	Home      template.HTMLAttr
	Discovery template.HTMLAttr
	Training  template.HTMLAttr
	Rental    template.HTMLAttr
	About     template.HTMLAttr
	Contact   template.HTMLAttr
}

func index(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/site.html",
		"templates/threecolumn.html",
		"templates/index.html",
	))
	menu := Menu{Home:template.HTMLAttr("class='active'"),
		Discovery:"",
		Training:"",
		Rental:"",	
		About:"",
		Contact:""}
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
	menu := Menu{Home:"",
		Discovery:template.HTMLAttr("class='active'"),
		Training:"",
		Rental:"",	
		About:"",
		Contact:""}
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
	menu := Menu{Home:"",
		Discovery:"",
		Training:template.HTMLAttr("class='active'"),
		Rental:"",	
		About:"",
		Contact:""}
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
	menu := Menu{Home:"",
		Discovery:"",
		Training:"",
		Rental:template.HTMLAttr("class='active'"),
		About:"",
		Contact:""}
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
	menu := Menu{Home:"",
		Discovery:"",
		Training:"",
		Rental:"",	
		About:template.HTMLAttr("class='active'"),
		Contact:""}
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
	menu := Menu{Home:"",
		Discovery:"",
		Training:"",
		Rental:"",	
		About:"",
		Contact:template.HTMLAttr("class='active'"),}
	err := t.Execute(w, menu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
