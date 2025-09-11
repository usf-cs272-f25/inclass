package main

import (
	"fmt"
	"net/http"
)

func startServer() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.Handle("/top10", http.FileServer(http.Dir("./static/top10")))
	http.Handle("/search", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		fmt.Fprintln(w, "search term: ", q)
	}))
	http.ListenAndServe(":8080", nil)
}

func main() {
	startServer()
}
