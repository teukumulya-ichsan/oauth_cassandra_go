package users

import "oauth_cassandra_golang/src/utils/errors"

type IUserRepository interface {
	LoginUser(string, string) (*User, *errors.RestErr)
}
