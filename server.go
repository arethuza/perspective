package main

import (
	"net/http"
	"path"
	"strings"
	core "github.com/arethuza/perspective/core"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := path.Clean(r.URL.Path)
	method := strings.ToLower(r.Method)
	// Load all of the supplied params from the request into a map - apart from "action"
	r.ParseForm()
	action := r.Form.Get("action")
	if action == "" {
		action = method
	}
	args := make(map[string]string)
	for name, value := range r.Form {
		args[strings.ToLower(name)] = value[0]
	}
	delete(args, "action")
	// Invoke the dispatcher to process the request
	result, err := core.Process(path, method, action, &args)
	if err != nil {
		http.Error(w, err.Error(), err.Code)
		return
	} else {
		sendResponse(&w, result)
	}
}

func sendResponse(w *http.ResponseWriter, result interface{}) error {
	switch result.(type) {
	}
	return nil
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

