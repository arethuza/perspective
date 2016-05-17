package main

import (
	"net/http"
	"path"
	dispatch "github.com/arethuza/perspective/dispatch"
	"strings"
)

var dispatcher = dispatch.Dispatcher()

func handler(w http.ResponseWriter, r *http.Request) {
	path := path.Clean(r.URL.Path)
	method := strings.ToLower(r.Method)
	r.ParseForm()
	action := r.Form.Get("action")
	args := make(map[string]string)
	for name, value := range r.Form {
		args[strings.ToLower(name)] = value[0]
	}
	delete(args, "action")
	result, err := dispatcher.Process(path, method, action, &args)
	if err != nil {
		http.Error(w, err.Error(), err.Code)
	} else {
		switch result.(type) {
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

