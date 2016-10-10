package dispatcher

import (
	"github.com/arethuza/perspective/items"
)

func Load(path string) (items.Item, *items.HttpError) {
	if path == "/" {
		return items.RootItem{}, nil
	}
	return nil, nil
}

