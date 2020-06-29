package oauth

import (
	"oauth_cassandra_golang/src/infrastructure/db/cassandra"
	"oauth_cassandra_golang/src/utils/errors"
)

type dbRepository struct {
}

func (r *dbRepository) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return nil, errors.NewInternalServerError("database connection not implemented")
}

// NewDBRepository constructor func
func NewDBRepository() IDbRepository {
	return &dbRepository{}
}
