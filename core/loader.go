package core


type Item interface {
	TypeName() string
}


func Load(path string) (Item, error) {
	if path == "/" {
		return RootItem{}, nil
	}
	return nil, nil
}
