package app

import (
	"oauth_cassandra_golang/src/domain/oauth"
	"oauth_cassandra_golang/src/infrastructure/db/cassandra"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication func to initialize starting app
func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	atService := oauth.NewService(oauth.NewDBRepository())
	atHandler := oauth.NewAccessTokenDelivery(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":9090")
}
