package app

import (
	"github.com/gin-gonic/gin"
	"oauth_cassandra_golang/src/domain/oauth"
)

var (
	router = gin.Default()
)

// StartApplication func to initialize starting app
func StartApplication() {
	atService := oauth.NewOAuthService(oauth.NewOAuthRepository())
	atHandler := oauth.NewOAuthDelivery(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":9090")
}
