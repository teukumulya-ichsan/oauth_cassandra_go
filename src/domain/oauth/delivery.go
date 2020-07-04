package oauth

import (
	"net/http"
	"oauth_cassandra_golang/src/utils/errors"

	"github.com/gin-gonic/gin"
)

// AccessTokenDelivery struct
type AccessTokenDelivery struct {
	Service IService
}

// GetByID controller
func (delivery *AccessTokenDelivery) GetByID(c *gin.Context) {
	accessToken, err := delivery.Service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

// Create controller
func (delivery *AccessTokenDelivery) Create(c *gin.Context) {
	var accessToken AccessToken
	if err := c.ShouldBindJSON(&accessToken); err != nil {
		restErr := errors.NewBadRequestError(invalidJSONBody)
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := delivery.Service.Create(accessToken); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}

// NewAccessTokenDelivery constructor
func NewAccessTokenDelivery(service IService) IAccessTokenDelivery {
	return &AccessTokenDelivery{
		Service: service,
	}
}
