package repository

import (
	"context"
	"database/sql"
)

type deleteUser struct {
	DB *sql.DB
}

func NewDeleteUserRepository(db *sql.DB) DeleteUser {
	return &deleteUser{
		DB: db,
	}
}

func (du *deleteUser) Delete(ctx context.Context, username string) error {
	return nil
}
