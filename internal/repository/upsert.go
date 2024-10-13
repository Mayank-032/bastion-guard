package repository

import (
	"context"
	"database/sql"

	"github.com/Mayank-032/bastion-guard/internal/domain"
)

type upsertUser struct {
	DB *sql.DB
}

func NewUpsertUserRepository(db *sql.DB) UpsertUser {
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