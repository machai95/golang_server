package main

import (
	"io"
	"log"
	"net/http"
)

const (
	link = `<link rel="stylesheet" href="/path/to/style.css">`
)

func main() {
	fs := http.FileServer(http.Dir("/upload"))
	var handler http.HandlerFunc
	handler = func(w http.ResponseWriter, r *http.Request) {
		var (
			url   = r.URL.Path
			isDir = url[len(url)-1] == '/'
		)
		fs.ServeHTTP(w, r)
		if isDir {
			io.WriteString(w, link)
		}
	}
	log.Fatal(http.ListenAndServe(":8080", handler))
}
