package dispatcher

import "github.com/arethuza/perspective/items"

func init() {
	name := "RootItem"
	addAction(name, AuthSuper, "get", "get", items.GetRoot)
	addAction(name, AuthSuper, "post", "post", items.CreateTenancy)
	addAction(name, AuthNone, "post", "init", items.InitSystem)
	addAction(name, AuthNone, "post", "login", items.LoginSuperUser)
}
