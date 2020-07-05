package users

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://localhost:8080",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@gmail.com","password":"password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	repository := UserRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid rest client response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://localhost:8080",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email": "email@gmail.com","password": "password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials", "status": "404", "error": "not_found"}`,
	})
	repository := UserRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid rest client response when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://localhost:8080",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email": "email@gmail.com","password": "password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials", "status": 404", "error": "not_found"}`,
	})
	repository := UserRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid login credentials", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://localhost:8080",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email": "email@gmail.com","password": "password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "1", "first_name": "Mulia", "last_name": "Ichsan", "email": "email@gmail.com"}`,
	})
	repository := UserRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying unmarshal user response", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://localhost:8080",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email": "email@gmail.com","password": "password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1, "first_name": "Mulia", "last_name": "Ichsan", "email": "email@gmail.com"}`,
	})
	repository := UserRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Mulia", user.FirstName)
	assert.EqualValues(t, "Ichsan", user.LastName)
	assert.EqualValues(t, "email", user.Email)
}
