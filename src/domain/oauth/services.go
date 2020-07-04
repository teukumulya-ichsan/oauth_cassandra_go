package oauth

import (
	"oauth_cassandra_golang/src/utils/errors"
	"strings"
)

// Service struct
type Service struct {
	Repository IDbRepository
}

// GetByID service func to get token by ID
func (s *Service) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {
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
func (s *Service) Create(accessToken AccessToken) *errors.RestErr {
	if err := accessToken.Validate(); err != nil {
		return err
	}

	return s.Repository.Create(accessToken)
}

// UpdateExpirationTime func to update expires access token
func (s *Service) UpdateExpirationTime(accessToken AccessToken) *errors.RestErr {
	if err := accessToken.Validate(); err != nil {
		return err
	}

	return s.Repository.UpdateExpirationTime(accessToken)
}

// NewService constructor func
func NewService(repo IDbRepository) IService {
	return &Service{Repository: repo}
}
