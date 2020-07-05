package oauth

import (
	"oauth_cassandra_golang/src/utils/errors"
	"strings"
)

// OAuthService struct
type OAuthService struct {
	Repository IOAuthRepository
}

// GetByID service func to get token by ID
func (s *OAuthService) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError(invalidAccessTokenID)
	}
	accessToken, err := s.Repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

// Create service func to Create access token
func (s *OAuthService) Create(accessToken AccessToken) *errors.RestErr {
	if err := accessToken.Validate(); err != nil {
		return err
	}

	return s.Repository.Create(accessToken)
}

// UpdateExpirationTime func to update expires access token
func (s *OAuthService) UpdateExpirationTime(accessToken AccessToken) *errors.RestErr {
	if err := accessToken.Validate(); err != nil {
		return err
	}

	return s.Repository.UpdateExpirationTime(accessToken)
}

// NewOAuthService constructor func
func NewOAuthService(repo IOAuthRepository) IOAuthService {
	return &OAuthService{Repository: repo}
}
