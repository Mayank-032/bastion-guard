package repository

import (
	"context"
	"database/sql"

	"github.com/Mayank-032/bastion-guard/internal/domain"
)

type readUser struct {
	DB *sql.DB
}

func NewReadUserRepository(db *sql.DB) ReadUser {
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
