package oauth

import (
	"github.com/gocql/gocql"
	"oauth_cassandra_golang/src/infrastructure/db/cassandra"
	"oauth_cassandra_golang/src/utils/errors"
)

type dbRepository struct {
}

func (r *dbRepository) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {
	var result AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, accessTokenID).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError(errNotFoundAccessToken)
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (r *dbRepository) Create(accessToken AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		accessToken.AccessToken,
		accessToken.UserID,
		accessToken.ClientID,
		accessToken.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(accessToken AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires,
		accessToken.Expires,
		accessToken.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

// NewDBRepository constructor func
func NewDBRepository() IDbRepository {
	return &dbRepository{}
}
