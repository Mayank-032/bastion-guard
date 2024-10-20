package repository

import (
	"context"

	"github.com/Mayank-032/bastion-guard/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type readUser struct {
	DB *mongo.Client
}

func NewReadUserRepository(db *mongo.Client) ReadUser {
	return &readUser{
		DB: db,
	}
}

func (ru *readUser) FetchUser(ctx context.Context, username string) (*domain.User, error) {
	return nil, nil
}

func (ru *readUser) FetchHistory(ctx context.Context, userId string) (*domain.History, error) {
	return nil, nil
}
