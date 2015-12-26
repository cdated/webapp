package main

import (
	"fmt"
	"github.com/cdated/webapp/data"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, r)

		if err != nil {
			// Session error, use public layout
			generateHTML(w, threads, "layout", "public.navbar", "index")

		} else {
			// Session valid, use private layout
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func main() {
	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/tmp/webby/"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// index
	mux.HandleFunc("/", index)

	// error
	// 	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	// 	mux.HandleFunc("/login", login)
	// 	mux.HandleFunc("/logout", logout)
	// 	mux.HandleFunc("/signup", signup)
	// 	mux.HandleFunc("/signup_account", signupAccount)
	// 	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	// 	mux.HandleFunc("/thread/new", newThread)
	// 	mux.HandleFunc("/thread/create", createThread)
	// 	mux.HandleFunc("/thread/post", postThread)
	// 	mux.HandleFunc("/thread/read", readThread)

	// starting up the server
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
