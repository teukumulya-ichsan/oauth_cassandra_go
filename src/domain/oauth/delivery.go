package oauth

import (
	"net/http"
	"oauth_cassandra_golang/src/utils/errors"

	"github.com/gin-gonic/gin"
)

// OAuthDelivery struct
type OAuthDelivery struct {
	Service IOAuthService
}

// GetByID controller
func (delivery *OAuthDelivery) GetByID(c *gin.Context) {
	accessToken, err := delivery.Service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

// Create controller
func (delivery *OAuthDelivery) Create(c *gin.Context) {
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

// NewOAuthDelivery constructor
func NewOAuthDelivery(service IOAuthService) IOAuthDelivery {
	return &OAuthDelivery{
		Service: service,
	}
}
