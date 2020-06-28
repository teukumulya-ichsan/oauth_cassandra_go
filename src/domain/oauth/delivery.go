package oauth

import (
	"net/http"

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

// NewAccessTokenDelivery constructor
func NewAccessTokenDelivery(service IService) IAccessTokenDelivery {
	return &AccessTokenDelivery{
		Service: service,
	}
}
