package main

import (
	"github.com/arethuza/perspective/dispatcher"
	"net/http"
	"path"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := path.Clean(r.URL.Path)
	method := strings.ToLower(r.Method)
	// Load all of the supplied params from the request into a map - apart from "action"
	r.ParseForm()
	action := strings.ToLower(r.Form.Get("action"))
	if action == "" {
		action = method
	}
	args := make(map[string]string)
	for name, value := range r.Form {
		args[strings.ToLower(name)] = value[0]
	}
	delete(args, "action")
	// Invoke the dispatcher to process the request
	actionResult, err := dispatcher.Process(path, method, action, &args)
	if err == nil {
		// No error so return a normal response
		actionResult.SendResponse(w)
	} else {
		// We got an error so return its details
		http.Error(w, err.Error(), err.Code)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
