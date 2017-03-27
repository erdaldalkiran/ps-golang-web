package controllers

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/oxtoacart/bpool"
)

var templates map[string]*template.Template
var bufpool *bpool.BufferPool

func Initialize() {
	bufpool = bpool.NewBufferPool(32)
	populateTemplates()
}

func Register() {
	registerController()

	http.HandleFunc("/img/", func(w http.ResponseWriter, req *http.Request) {
		err := serverResource(w, "public"+req.URL.Path, "image/png")
		if err != nil {
			handleError(w, err)
		}
	})

	http.HandleFunc("/css/", func(w http.ResponseWriter, req *http.Request) {
		err := serverResource(w, "public"+req.URL.Path, "text/css")
		if err != nil {
			handleError(w, err)
		}
	})
}

func registerController() {
	hc := new(homeController)
	hc.register()

	cc := new(categoriesController)
	cc.register()

	psc := new(productsController)
	psc.register()

	pc := new(productController)
	pc.register()
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}

func serverResource(w http.ResponseWriter, fileName string, contentType string) error { //use struct instead of seperate parameters
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	w.Header().Set("Content-Type", contentType)
	reader.WriteTo(w)

	return nil
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("the template %s does not exist", name)
	}
	buff := bufpool.Get()
	defer bufpool.Put(buff)

	err := tmpl.ExecuteTemplate(buff, "base.tmpl", data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buff.WriteTo(w)
	return nil
}

func populateTemplates() map[string]*template.Template {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	const templatesDir = "templates/" //move to config

	commonTemplates, err := filepath.Glob(templatesDir + "common/*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	pagesDir := templatesDir + "pages/"
	pageDirs, err := ioutil.ReadDir(pagesDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, pageDir := range pageDirs {
		pageDirPath := filepath.Join(pagesDir, pageDir.Name())

		pageTemplates, err := filepath.Glob(pageDirPath + "/*.tmpl")
		if err != nil {
			log.Fatal(err)
		}

		files := append(commonTemplates, pageTemplates...)
		pageName := filepath.Base(pageDir.Name())

		templates[pageName] = template.Must(template.ParseFiles(files...))
	}

	return templates
}
