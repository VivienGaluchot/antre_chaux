package main

import (
	"log"
	"net/http"
	"regexp"
)

// url handling

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(ViewHandler))
	http.HandleFunc("/edit/", makeHandler(EditHandler))
	http.HandleFunc("/save/", makeHandler(SaveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
