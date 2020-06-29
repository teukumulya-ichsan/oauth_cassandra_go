package oauth

import (
	"oauth_cassandra_golang/src/utils/errors"

	"github.com/gin-gonic/gin"
)

// IAccessTokenDelivery interface
type IAccessTokenDelivery interface {
	GetByID(ctx *gin.Context)
	Create(ctx *gin.Context)
}

// IService interface
type IService interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

// IDbRepository interface
type IDbRepository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}
