package main

import (
	"fmt"
	"net/http"
	"path"
	dispatch "github.com/arethuza/perspective/dispatch"
)

var dispatcher = dispatch.Dispatcher()

func handler(w http.ResponseWriter, r *http.Request) {
	path := path.Clean(r.URL.Path)
	fmt.Fprintf(w, "%s\n", path)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

