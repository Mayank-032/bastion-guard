package repository

import (
	"context"

	"github.com/Mayank-032/bastion-guard/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type upsertUser struct {
	DB *mongo.Client
}

func NewUpsertUserRepository(db *mongo.Client) UpsertUser {
	return &upsertUser{
		DB: db,
	}
}

func (uu *upsertUser) CreateUser(ctx context.Context, user domain.User) error {
	return nil
}

func (uu *upsertUser) UpdatePassword(ctx context.Context, newPassword string, user domain.User) error {
	return nil
}
