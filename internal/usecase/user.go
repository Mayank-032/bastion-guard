package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/Mayank-032/bastion-guard/internal/domain"
	"github.com/Mayank-032/bastion-guard/internal/repository"
)

type user struct {
	ReadUserRepo   repository.ReadUser
	UpsertUserRepo repository.UpsertUser
	DeleteUserRepo repository.DeleteUser
}

func NewLoginUsecase(readUserRepo repository.ReadUser,
	upsertUserRepo repository.UpsertUser,
	deleteUserRepo repository.DeleteUser) *user {
	return &user{
		ReadUserRepo:   readUserRepo,
		UpsertUserRepo: upsertUserRepo,
		DeleteUserRepo: deleteUserRepo,
	}
}

// check if user is registered to our application or not
func (u *user) IsCreated(ctx context.Context, username, password string) (bool, error) {
	user, err := u.ReadUserRepo.FetchUser(ctx, username)
	if err != nil {
		log.Printf("err: %v\n", err.Error())
		return false, errors.New("unable to fetch user details")
	}

	// if user is not nil and password is wrong, invalidate user
	if user != nil && user.Password != password {
		return false, errors.New("invalid user")
	}

	// if user is nil, which means user is not created
	if user == nil {
		return false, nil
	}

	return true, nil
}

// registers the user to our application
func (u *user) Create(ctx context.Context, username, password string) error {
	err := u.UpsertUserRepo.CreateUser(ctx, domain.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Printf("err: %v\n", err.Error())
		return errors.New("unable to create user")
	}

	return nil
}

// mark user inactive
func (u *user) MarkInactive(ctx context.Context, username string) error {
	err := u.DeleteUserRepo.Delete(ctx, username)
	if err != nil {
		log.Printf("err: %v\n", err.Error())
		return errors.New("unable to delete user")
	}
	return nil
}

// update old password with new password
func (u *user) UpdatePassword(ctx context.Context, username, oldPassword, newPassword string) error {
	err := u.UpsertUserRepo.UpdatePassword(ctx, newPassword, domain.User{
		Username: username,
		Password: oldPassword,
	})
	if err != nil {
		log.Printf("err: %v", err.Error())
		return errors.New("unable to update old password to newpassword")
	}
	
	return nil
}
