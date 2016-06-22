package core

type RootItem struct {

}

func (item RootItem) TypeName() string {
	return "RootItem"
}

func init() {
	AddAction("RootItem", AuthRoot, "get", "get", getRoot)
}

func getRoot(item Item) {

}
