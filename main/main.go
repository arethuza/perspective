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
	user := authenticate(r)
	path := path.Clean(r.URL.Path)
	context, _ := misc.CreateContext(path, config)
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
	actionResult, requestErr := dispatcher.Process(context, user, path, method, action, &args, body)
	if requestErr == nil {
		// No error so return a normal response
		actionResult.SendResponse(w)
	} else {
		// We got an HTTP error so return its details
		http.Error(w, requestErr.Error(), requestErr.Code)
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

func authenticate(r *http.Request) *items.User {
	return nil
}
