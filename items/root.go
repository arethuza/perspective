package items

import (
	"encoding/json"
	"github.com/arethuza/perspective/misc"
	"github.com/arethuza/perspective/database"
)

type RootItem struct {
}

func (item RootItem) TypeName() string {
	return "RootItem"
}

func GetRoot(context *misc.Context, user *User, item Item, args RequestArgs, body []byte) (ActionResult, *HttpError) {
	var a [0]string
	return JsonResult{value: a}, nil
}

type CreateTenancyRequest struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type CreateTenancyResponse struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
}

func CreateTenancy(context *misc.Context, user *User, item Item, args RequestArgs, body []byte) (ActionResult, *HttpError) {
	var request CreateTenancyRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return nil, &HttpError{Message: err.Error()}
	}
	//
	err = database.CreateTenancy(request.Name, request.UserName, request.Password)
	if err != nil {
		return nil, &HttpError{Message: err.Error()}
	}
	response := CreateTenancyResponse{Name: request.Name, UserName:request.UserName}
	return JsonResult{value:response}, nil
}
