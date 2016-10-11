package items

type RootItem struct {
}

func (item RootItem) TypeName() string {
	return "RootItem"
}

func GetRoot(item Item) (ActionResult, *HttpError) {
	var a [0]string
	return JsonResult{value: a}, nil
}

func CreateAccount(item Item) (ActionResult, *HttpError) {
	return nil, nil
}
