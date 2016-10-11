package dispatcher

import "github.com/arethuza/perspective/items"

func init() {
	name := "RootItem"
	addAction(name, AuthRoot, "get", "get", items.GetRoot)
	addAction(name, AuthRoot, "post", "post", items.CreateTenancy)
}
