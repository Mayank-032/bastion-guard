package usecase

import "context"

type User interface {
	IsCreated(ctx context.Context, username, password string) (bool, error)
	Create(ctx context.Context, username, password string) error
	MarkInactive(ctx context.Context, username string) error
	UpdatePassword(ctx context.Context, username, oldPassword, newPassword string) error
}
