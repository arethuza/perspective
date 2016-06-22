package core

type AuthorizationLevel int

const (
	AuthRoot AuthorizationLevel = iota
	AuthAdmin
	AuthWriter
	AuthReader
	AuthNone
)

type Action func(item Item)

type dispatchEntry struct {
	authorization AuthorizationLevel
	action Action
}

var dispatchTable = make(map[string](map[string][]dispatchEntry))

func AddAction(typeName string, authLevel AuthorizationLevel, method string, name string, action Action) {
	actionTable, ok := dispatchTable[typeName]
	if !ok {
		actionTable = make(map[string][]dispatchEntry)
		dispatchTable[typeName] = actionTable
	}
	key := method + ":" + name
	dispatchList, ok := actionTable[key]
	entry := dispatchEntry{authorization:authLevel, action:action}
	actionTable[key] = append(dispatchList, entry)
}

func Process(path, method, action string, args *map[string]string) (interface{}, *HttpError) {
	item, err := Load(path)
	if err != nil {
		return nil, &HttpError{}
	}
	itemAction, err := GetAction(item.TypeName(), AuthRoot, method, action)
	itemAction(item)
	return nil, &HttpError{}
}

func GetAction(typeName string, userAuthLevel AuthorizationLevel, method string, name string) (Action, error) {
	actionTable, ok := dispatchTable[typeName]
	if !ok {
	}
	key := method + ":" + name
	dispatchList, ok := actionTable[key]
	if !ok {
	}
	for _, entry := range dispatchList {
		if userAuthLevel >= entry.authorization {
			return entry.action, nil
		}
	}
	return nil, nil
}

type HttpError struct {
	Code int
	message string
}

func (he HttpError) Error() string {
	return he.message
}