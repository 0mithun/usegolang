package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Parsing template %v failed.", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return Template{}, fmt.Errorf("Parsing template %v", err)
	}

	return Template{
		HTMLTpl: tpl,
	}, nil
}

type Template struct {
	HTMLTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.HTMLTpl.Execute(w, data)
	if err != nil {
		log.Printf("Execute template %v failed.", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}
