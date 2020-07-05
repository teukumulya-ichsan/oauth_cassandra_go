package oauth

import (
	"oauth_cassandra_golang/src/utils/errors"

	"github.com/gin-gonic/gin"
)

// IOAuthDelivery interface
type IOAuthDelivery interface {
	GetByID(ctx *gin.Context)
	Create(ctx *gin.Context)
}

// IOAuthService interface
type IOAuthService interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

// IOAuthRepository interface
type IOAuthRepository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}
