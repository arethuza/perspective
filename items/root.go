package items


type RootItem struct {
}

func (item RootItem) TypeName() string {
	return "RootItem"
}

func GetRoot(item Item) (*ActionResult, *HttpError) {
	return nil, nil
}

func CreateAccount(item Item) (*ActionResult, *HttpError){
	return nil, nil
}

