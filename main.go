package main

import "html/template"
import "path/filepath"
import "log"
import "net/http"
import "fmt"

var templates map[string]*template.Template

func main() {
	initialize()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		tmplName := req.URL.Path[1:] + ".tmpl"
		err := renderTemplate(w, tmplName, nil)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}
	})

	http.ListenAndServe(":8080", nil)

}

func initialize() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	const templatesDir = "templates/" //move to config
	templateNames, err := filepath.Glob(templatesDir + "pages/*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	commonTemplates, err := filepath.Glob(templatesDir + "includes/*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	for _, templateName := range templateNames {
		files := append(commonTemplates, templateName)
		templates[filepath.Base(templateName)] = template.Must(template.ParseFiles(files...))
	}

}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("the template %s does not exist", name)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl.ExecuteTemplate(w, "base.tmpl", data)
}
