package main

import (
	"fmt"
	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	fmt.Println(err)
	if err == nil {

		_, err := session(w, r)
		public_tmpl_files := []string{"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html",
		}
		private_tmpl_files := []string{"templates/layout.html",
			"templates/private.navbar.html",
			"templates/index.html",
		}

		var templates *template.Template

		if err != nil {
			templates = template.Must(template.ParseFiles(private_tmpl_files...))

		} else {
			templates = template.Must(template.ParseFiles(public_tmpl_files...))

		}

		templates.ExecuteTemplate(w, "layout", threads)
	} else {
		fmt.Println("Error with threads")
	}
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
