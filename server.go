package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Body  string
}

func handler(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
	}
	tpl, _ := Asset("data/layout.tpl")
	t := template.Must(template.New("hello").Funcs(funcMap).Parse(string(tpl)))
	page := Page{Title: "hello, world", Body: "<b>hello!!</b>"}
	err := t.Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
