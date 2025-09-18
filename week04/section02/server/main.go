package main

import (
	"fmt"
	"net/http"
	"time"
)

func startServer(idx Index) {
	// Use http.Dir to serve the contents of ./static for GET requests
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		hits := idx.Search(q)
		fmt.Fprintln(w, "search term: ", q)
	})
	go http.ListenAndServe(":8080", nil)
}

func main() {
	idx := NewIndex()
	startServer(idx)
	crawl(idx)
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
