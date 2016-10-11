package main

import (
	"github.com/arethuza/perspective/dispatcher"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := path.Clean(r.URL.Path)
	method := strings.ToLower(r.Method)
	args := make(map[string]string)
	var body []byte = nil
	var action string
	if method == "get" || method == "delete" {
		r.ParseForm()
		action = strings.ToLower(r.Form.Get("action"))
		if action == "" {
			action = method
		}
		for name, value := range r.Form {
			args[strings.ToLower(name)] = value[0]
		}
		delete(args, "action")
	} else if method == "post" || method == "put" {
		action = method
		body, _ = ioutil.ReadAll(r.Body)
	}
	// Invoke the dispatcher to process the request
	actionResult, err := dispatcher.Process(path, method, action, &args, body)
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
