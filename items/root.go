package items

import "encoding/json"

type RootItem struct {
}

func (item RootItem) TypeName() string {
	return "RootItem"
}

func GetRoot(item Item, args RequestArgs, body []byte) (ActionResult, *HttpError) {
	var a [0]string
	return JsonResult{value: a}, nil
}

type CreateTenancyRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func CreateTenancy(item Item, args RequestArgs, body []byte) (ActionResult, *HttpError) {
	var request CreateTenancyRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return nil, &HttpError{Message: err.Error()}
	}
	return nil, nil
}
