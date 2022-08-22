package repository

import (
	"context"

	"github.com/Edilberto-Vazquez/golang-rest-and-websockets/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, use *models.User) error
	GetUserById(ctx context.Context, id int64) (*models.User, error)
}

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}
