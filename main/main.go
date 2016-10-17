package main

import (
	"fmt"
	"github.com/arethuza/perspective/dispatcher"
	"github.com/arethuza/perspective/items"
	"github.com/arethuza/perspective/misc"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"
)

var config *misc.Config

func handler(w http.ResponseWriter, r *http.Request) {
	path := path.Clean(r.URL.Path)
	context, err := items.CreateContext(path, config)
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
	actionResult, err := dispatcher.Process(context, path, method, action, &args, body)
	if err == nil {
		// No error so return a normal response
		actionResult.SendResponse(w)
	} else if httpError, ok := err.(items.HttpError); ok {
		// We got an HTTP error so return its details
		http.Error(w, err.Error(), httpError.Code)
	} else {
		// Default return error and a 500
		http.Error(w, err.Error(), 500)
	}
}

func main() {
	var err error
	config, err = misc.LoadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	http.HandleFunc("/", handler)
	addr := ":" + strconv.Itoa(config.Port)
	http.ListenAndServe(addr, nil)
}
