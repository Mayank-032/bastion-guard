package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type deleteUser struct {
	DB *mongo.Client
}

func NewDeleteUserRepository(db *mongo.Client) DeleteUser {
	return &deleteUser{
		DB: db,
	}
}

func (du *deleteUser) Delete(ctx context.Context, username string) error {
	return nil
}
