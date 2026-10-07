package main

import (
	"fmt"
	"net/http"
)

const port = "8080"

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, World from Go!</h1>")
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Go server running at http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
