package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

func projectsPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/projects.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

// Simple health check, handy for Kubernetes probes later.
func healthPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte("ok"))
}

// "/" shows the home page; any other unknown path is a real 404.
func rootPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	homePage(w, r)
}

func main() {
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/projects", projectsPage)
	http.HandleFunc("/contact", contactPage)
	http.HandleFunc("/healthz", healthPage)

	// Serves static/style.css and static/images/*
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
