package dispatcher

import "github.com/arethuza/perspective/items"

func init() {
	name := "RootItem"
	AddAction(name, AuthRoot, "get", "get", items.GetRoot)
	AddAction(name, AuthRoot, "post", "createAccount", items.CreateAccount)
}
