package main

import (
	"log"
	"net/http"
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
	w.Write([]byte("Display a message here"))
}

func messageCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new message"))
}

func main() {

	// init server multiplexer and add page routes funtion
	// OBS: i am no t using default servermux and http.HandleFunc
	// to avoid a global scope servermux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/message/view", messageView)
	mux.HandleFunc("/message/create", messageCreate)

	// start server on localhost port 4000. Albeit,
	// as I did not specified host, so the app is listening
	// to port 4000 in every computer network interface.

	// I can also use an alias as port name
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
