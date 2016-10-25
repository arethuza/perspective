package items

import (
	"encoding/json"
	"github.com/arethuza/perspective/cache"
	"github.com/arethuza/perspective/database"
	"github.com/arethuza/perspective/misc"
	"golang.org/x/crypto/bcrypt"
	"strconv"
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

//
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
//

type InitSystemResponse struct {
	Password string `json:"password"`
}

func InitSystem(context *misc.Context, user *User, item Item, args RequestArgs, body []byte) (ActionResult, *HttpError) {
	// Get the root superuser
	initSuperUser, err := database.ReadSuperUserById(context.DatabaseConnection, 1)
	if initSuperUser != nil {
		// Already exists
		return nil, &HttpError{Message: "init has already been executed"}
	}
	// Create password
	password, _, err := misc.GenerateRandomString(context.Config.PasswordLength)
	if err != nil {
		return nil, &HttpError{Message: err.Error()}
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), context.Config.BcryptCost)
	if err != nil {
		return nil, &HttpError{Message: err.Error()}
	}
	// Actually create the root superuser
	superUserId, err := database.CreateSuperUser(context.DatabaseConnection, "root", hash)
	if superUserId != 1 {
		return nil, &HttpError{Message: "Error creating root superuser - id != 1"}
	} else if err != nil {
		return nil, &HttpError{Message: err.Error()}
	}
	// Return response with password
	response := InitSystemResponse{Password: password}
	return JsonResult{value: response}, nil
}

//
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
//

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginSuperUser(context *misc.Context, user *User, item Item, args RequestArgs, body []byte) (ActionResult, *HttpError) {
	var request LoginRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return nil, &HttpError{Message: err.Error()}
	}
	superUser, err := database.ReadSuperUserByName(context.DatabaseConnection, request.UserName)
	if err != nil {
		return nil, &HttpError{Message: err.Error()}
	}
	passwordBytes := []byte(request.Password)
	err = bcrypt.CompareHashAndPassword(superUser.PasswordHash, passwordBytes)
	if err != nil {
		return nil, &HttpError{Message: "bad username or password"}
	}
	// Password has matched
	token, err := cache.CreateUserSession(context.Config, superUser)
	if err != nil {
		return nil, &HttpError{Message: err.Error()}
	}
	//
	var response LoginResponse
	response.Token = token
	return JsonResult{value: response}, nil
}

//
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
//

type CreateTenancyRequest struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type CreateTenancyResponse struct {
	Name      string `json:"name"`
	UserName  string `json:"username"`
	TenancyId string `json:"tenantid"`
	Password  string `json:"password"`
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
		password, _, err = misc.GenerateRandomString(context.Config.PasswordLength)
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
	response := CreateTenancyResponse{Name: request.Name, UserName: request.UserName, TenancyId: strconv.Itoa(tenancyId)}
	if request.Password == "" {
		response.Password = password
	}
	return JsonResult{value: response}, nil
}
