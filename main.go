package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from a boring project"))
}

func main() {

	// init server multiplexer and add root "page" funtion
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// start server on localhost port 4000. Albeit,
	// as I did not specified host, so the app is listening
	// to port 4000 in every computer network interface.

	// I can also use an alias as port name
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
