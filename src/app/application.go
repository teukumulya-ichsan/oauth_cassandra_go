package app

import (
	"oauth_cassandra_golang/src/domain/oauth"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication func to initialize starting app
func StartApplication() {
	atService := oauth.NewService(oauth.NewDBRepository())
	atHandler := oauth.NewAccessTokenDelivery(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)

	router.Run(":9090")
}
