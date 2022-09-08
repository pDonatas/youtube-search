package main

import (
	"fmt"
	ytsearch "github.com/AnjanaMadu/YTSearch"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func search(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	title := r.URL.Query().Get("title")

	id := getVideoId(title)

	// create json response
	var json = fmt.Sprintf(`{"id": "%s"}`, id)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err := fmt.Fprintf(w, json)
	if err != nil {
		return
	}
}

func main() {
	// Make http server
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", search)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func getVideoId(title string) string {
	results, err := ytsearch.Search(title)
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		if result.VideoId != "" {
			return result.VideoId
		}
	}

	return ""
}
