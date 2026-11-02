package main

import (
	"fmt"
	"log"
	"net/http"
)

// hello sends a web page back to the browser.
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, World from Go!</h1>")
}

func main() {
	// When someone opens the home page ("/"), run the hello function.
	http.HandleFunc("/", hello)

	// Start the server. It keeps running until you stop it.
	// 0.0.0.0 means it also works inside Docker.
	// If it cannot start, log.Fatal prints the reason.
	fmt.Println("Open http://127.0.0.1:8001 in your browser")
	log.Fatal(http.ListenAndServe("0.0.0.0:8001", nil))
}
