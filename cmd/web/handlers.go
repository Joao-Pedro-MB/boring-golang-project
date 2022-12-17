package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// check if received path is just a "/". if not show a notFound error
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from a boring project"))
}

func messageView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display message ID: %d...", id)
}

func messageCreate(w http.ResponseWriter, r *http.Request) {

	// Use r.Method to check whether the request is using POST or not.
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new message"))
}
