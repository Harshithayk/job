package models

import (
	"context"
	"project/models"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	CreateUser(ctx context.Context, nu UserSignUp) (User, error)
	Authenticate(ctx context.Context, email, password string) (jwt.RegisteredClaims, error)
	CreateCompany(ctx context.Context, ni CreateCompany, userId uint) (Company, error)
	ViewCompany(ctx context.Context, userId string) ([]Company, error)
	Getcompany(id int64) (Company, error)
	JobCreation(ctx context.Context, ni CreateJob, id uint64) (CreateJob, error)
}
type Store struct {
	Service
}

func NewStore(s Service) Store {
	return Store{Service: s}
}
