package repository

import (
	"context"

	"github.com/Mayank-032/bastion-guard/internal/domain"
)

type ReadUser interface {
	FetchUser(ctx context.Context, username string) (*domain.User, error)
	FetchHistory(ctx context.Context, userId string) (*domain.History, error)
}

type UpsertUser interface {
	CreateUser(ctx context.Context, user domain.User) error
	UpdatePassword(ctx context.Context, oldPassword string, user domain.User) error
}

type DeleteUser interface {
	Delete(ctx context.Context, username string) error
}
