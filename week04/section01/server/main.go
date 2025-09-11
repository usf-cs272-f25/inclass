package main

import (
	"fmt"
	"net/http"
	"time"
)

func startServer(index Index) {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		fmt.Fprintln(w, "search term: ", query)
		// results := index.Search(query)
	})
	go http.ListenAndServe(":8080", nil)
}

func main() {
	index := NewIndex()
	crawler.Crawl(index)
	startServer(index)

	for {
		time.Sleep(100 * time.Millisecond)
	}
}
