package dispatcher

import (
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

func AddAction(typeName string, authLevel AuthorizationLevel, method string, name string, action items.Action) {
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

func Process(path, method, action string, args *map[string]string) (*items.ActionResult, *items.HttpError) {
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

func getAction(typeName string, userAuthLevel AuthorizationLevel, method string, name string) (items.Action, *items.HttpError) {
	actionTable, ok := dispatchTable[typeName]
	if !ok {
	}
	key := method + ":" + name
	dispatchList, ok := actionTable[key]
	if !ok {
		return nil, &items.HttpError{}
	}
	for _, entry := range dispatchList {
		if userAuthLevel >= entry.authorization {
			return entry.action, nil
		}
	}
	return nil, &items.HttpError{}
}
