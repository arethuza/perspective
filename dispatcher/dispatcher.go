package dispatcher

import (
	"fmt"
	"github.com/arethuza/perspective/items"
	"strings"
)

type AuthorizationLevel int

const (
	AuthRoot AuthorizationLevel = iota
	AuthAdmin
	AuthWriter
	AuthReader
	AuthNone
)

type dispatchEntry struct {
	authorization AuthorizationLevel
	action        items.Action
}

var dispatchTable = make(map[string](map[string][]dispatchEntry))

func addAction(typeName string, authLevel AuthorizationLevel, method string, name string, action items.Action) {
	actionTable, ok := dispatchTable[typeName]
	if !ok {
		actionTable = make(map[string][]dispatchEntry)
		dispatchTable[typeName] = actionTable
	}
	key := strings.ToLower(method) + ":" + strings.ToLower(name)
	dispatchList, _ := actionTable[key]
	entry := dispatchEntry{authorization: authLevel, action: action}
	actionTable[key] = append(dispatchList, entry)
}

func Process(path, method, action string, args *map[string]string) (items.ActionResult, *items.HttpError) {
	item, err := Load(path)
	if err != nil {
		return nil, &items.HttpError{}
	}
	itemAction, err := getAction(item.TypeName(), AuthRoot, method, action)
	if err != nil {
		return nil, err
	}
	actionResult, err := itemAction(item)
	return actionResult, err
}

func getAction(typeName string, userAuthLevel AuthorizationLevel, method string, action string) (items.Action, *items.HttpError) {
	actionTable, ok := dispatchTable[typeName]
	if !ok {
		message := fmt.Sprintf("Unknown item type=%s", typeName)
		return nil, &items.HttpError{Code: 500, Message: message}
	}
	key := method + ":" + action
	dispatchList, ok := actionTable[key]
	if !ok {
		message := fmt.Sprintf("No actions for item type=%s, method=%s, action=%s", typeName, method, action)
		return nil, &items.HttpError{Code: 404, Message: message}
	}
	// Return the first action that we are authorized for, if it exists!
	for _, entry := range dispatchList {
		if userAuthLevel >= entry.authorization {
			return entry.action, nil
		}
	}
	// We haven't found any actions we are authorized to perform
	return nil, &items.HttpError{Code: 403, Message: "Unauthorized"}
}
