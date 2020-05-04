package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

var (
	indexTmpl = template.Must(
		template.ParseFiles(filepath.Join("templates", "index.html")),
	)
)

func main() {
	log.Println("Startup")

	http.HandleFunc("/", indexHandler)

	// Serve static files out of the static directory.
	// By configuring a static handler in app.yaml, App Engine serves all the
	// static content itself. As a result, the following two lines are in
	// effect for development only.
	static := http.StripPrefix("/static", http.FileServer(http.Dir("static")))
	http.Handle("/static/", static)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("Terminated")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	type indexData struct {
		Style       string
		RequestTime string
	}
	data := indexData{
		Style:       "/static/style.css",
		RequestTime: time.Now().Format(time.RFC822),
	}
	if err := indexTmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
