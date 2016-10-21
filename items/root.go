package items

import (
	"encoding/json"
	"github.com/arethuza/perspective/database"
	"github.com/arethuza/perspective/misc"
	"strconv"
	"golang.org/x/crypto/bcrypt"
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
	TenancyId string `json:"tenantid"`
	Password string `json:"password"`
}

func CreateTenancy(context *misc.Context, user *User, item Item, args RequestArgs, body []byte) (ActionResult, *HttpError) {
	var request CreateTenancyRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return nil, &HttpError{Message: err.Error()}
	}
	var password string
	if request.Password != "" {
		password = request.Password
	} else {
		password, err = misc.GenerateRandomString(context.Config.PasswordLength)
		if err != nil {
			return nil, &HttpError{Message: err.Error()}
		}
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), context.Config.BcryptCost)
	if err != nil {
		return nil, &HttpError{Message: err.Error()}
	}
	tenancyId, err := database.CreateTenancy(context.DatabaseConnection, request.Name, request.UserName, hash)
	if err != nil {
		return nil, &HttpError{Message: err.Error()}
	}
	response := CreateTenancyResponse{Name: request.Name, UserName: request.UserName, TenancyId:strconv.Itoa(tenancyId)}
	if request.Password == "" {
		response.Password = password
	}
	return JsonResult{value: response}, nil
}
