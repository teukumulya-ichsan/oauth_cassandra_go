package users

import (
	"encoding/json"
	"github.com/mercadolibre/golang-restclient/rest"
	"oauth_cassandra_golang/src/utils/errors"
	"time"
)

var (
	userRestClient = rest.RequestBuilder{
		Timeout: 100 * time.Millisecond,
		BaseURL: "https://localhost:8080",
	}
)

type UserRepository struct{}

func (u *UserRepository) LoginUser(email string, password string) (*User, *errors.RestErr) {
	request := LoginRequest{
		Email:    email,
		Password: password,
	}

	response := userRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid rest client response when trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying unmarshal user response")
	}
	return &user, nil
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}
