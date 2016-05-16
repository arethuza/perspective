package main

import (
	"fmt"
	"net/http"
	"path"
	"dispatcher"
)

const dispatcher = Dispatcher()

func handler(w http.ResponseWriter, r *http.Request) {
	path := path.Clean(r.URL.Path)
	fmt.Fprintf(w, "%s\n", path)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

