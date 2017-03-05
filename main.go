package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Content-Type", "text/html")

		body, _ := template.New("body").Parse(doc)
		header, _ := template.New("header").Parse(header)
		footer, _ := template.New("footer").Parse(footer)
		templates := template.New("main")
		templates.
		context := Context{
			[3]string{"Lemon", "Orange", "Apple"},
			"the title",
		}
		err := templates.Lookup("test").Execute(resp, context)
		if err != nil {
			resp.WriteHeader(404)
			resp.Write([]byte(err.Error()))
		}
	})

	http.ListenAndServe(":8080", nil)
}

const doc = `
{{template "header" .Title}}
  <body>
    <h1>List of Fruit</h1>
    <ul>
      {{range .Fruit}}
      	<li>{{.}}</li>
      {{end}}
    </ul>
  </body>
{{template "footer"}}
`

const header = `
<!DOCTYPE html>
<html>
  <head><title>{{.}}</title></head>
`

const footer = `
</html>
`

type Context struct {
	Fruit [3]string
	Title string
}
