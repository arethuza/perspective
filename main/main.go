package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/arethuza/perspective/cache"
	"github.com/arethuza/perspective/database"
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
	context, _ := misc.CreateContext(path, config)
	user, authLevel, authErr := authenticate(r, path, config)
	if authErr != nil {
		http.Error(w, "problem authenticating request:"+authErr.Error(), 400)
		return
	}
	method := strings.ToLower(r.Method)
	args := make(map[string]string)
	var body []byte = nil
	var action string
	if action == "" {
		action = method
	}
	if method == "get" || method == "delete" {
		r.ParseForm()
		action = strings.ToLower(r.Form.Get("action"))
		for name, value := range r.Form {
			args[strings.ToLower(name)] = value[0]
		}
		delete(args, "action")
	} else if method == "post" || method == "put" {
		var err error
		body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		r.ParseForm()
		action = strings.ToLower(r.Form.Get("action"))
	}
	// Invoke the dispatcher to process the request
	actionResult, requestErr := dispatcher.Process(context, user, authLevel, path, method, action, &args, body)
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
	err = cache.CreateRedisClient(config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	http.HandleFunc("/", handler)
	addr := ":" + strconv.Itoa(config.Port)
	http.ListenAndServe(addr, nil)
}

func authenticate(r *http.Request, path string, config *misc.Config) (items.User, dispatcher.AuthorizationLevel, error) {
	token := getBearerToken(r)
	if token == "" {
		return nil, dispatcher.AuthNone, nil
	}
	sessionData, err := cache.GetUserSessionData(token)
	if err != nil {
		return nil, dispatcher.AuthNone, errors.New("no matching session for supplied token")
	}
	if path == "/" {
		var superUser database.SuperUser
		json.Unmarshal(sessionData, &superUser)
		return superUser, dispatcher.AuthSuper, nil
	}
	return nil, dispatcher.AuthNone, nil
}

func getBearerToken(r *http.Request) string {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return ""

	}
	a := strings.Split(authorizationHeader, " ")
	if len(a) != 2 {
		return ""
	}
	if strings.ToLower(a[0]) != "bearer" {
		return ""
	}
	return a[1]
}
