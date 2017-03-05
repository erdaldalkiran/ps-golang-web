package main

import (
	"bufio"
	"net/http"
	"os"
	"path"
)

func main() {
	http.Handle("/", new(myHandler))
	http.ListenAndServe(":8080", nil)
}

type myHandler struct {
	http.Handler
}

func (handler *myHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	filePath := "public" + req.URL.Path

	contentType := getContentType(filePath)
	resp.Header().Add("Content-Type", contentType)

	f, err := os.Open(filePath)
	if err != nil {
		resp.WriteHeader(404)
		resp.Write([]byte("404 - " + err.Error()))
		return
	}

	bufferedReader := bufio.NewReader(f)
	bufferedReader.WriteTo(resp)
}

func getContentType(filePath string) string {
	var contentType string
	ext := path.Ext(filePath)[1:]
	switch ext {
	case "html":
		contentType = "text/html"
	case "css":
		contentType = "text/css"
	case "png":
		contentType = "image/png"
	case "js":
		contentType = "application/javascript"
	case "mp4":
		contentType = "video/mp4"
	default:
		contentType = "text/plain"
	}

	return contentType
}
