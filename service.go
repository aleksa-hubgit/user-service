package main

import (
	"context"
	"errors"

	"github.com/aleksa-hubgit/user-service/data"
)

type CreateRequest struct {
	Username string
	Email    string
	Password string
}
type UserUpdateRequest struct {
	Username string
	Email    string
}
type UserDeleteRequest struct {
	Username string
	Password string
}

type Service interface {
	GetUserByUsername(context.Context, string) (*data.User, error)
	ListUsers(context.Context) ([]data.User, error)
	CreateUser(context.Context, CreateRequest) (*data.User, error)
	UpdateUser(context.Context, UserUpdateRequest) error
	DeleteUser(context.Context, UserDeleteRequest) error
}

type UserService struct {
	queries *data.Queries
}

func (u *UserService) exists(ctx context.Context, username string) bool {
	_, err := u.queries.GetUserByUsername(ctx, username)
	return err == nil
}

// CreateUser implements Service.
func (u *UserService) CreateUser(ctx context.Context, rr CreateRequest) (*data.User, error) {
	if u.exists(ctx, rr.Username) {
		return nil, errors.New("user exists")
	}
	user, err := u.queries.CreateUser(ctx, data.CreateUserParams{Username: rr.Username, Email: rr.Email, Password: rr.Password})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser implements Service.
func (u *UserService) DeleteUser(ctx context.Context, udr UserDeleteRequest) error {
	toDelete, err := u.queries.GetUserByUsername(ctx, udr.Username)
	if err != nil {
		return err
	}
	if toDelete.Password != udr.Password {
		return errors.New("passwords don't match")
	}
	u.queries.DeleteUser(ctx, toDelete.ID)
	return nil
}

// GetUserByUsername implements Service.
func (u *UserService) GetUserByUsername(ctx context.Context, uname string) (*data.User, error) {
	user, err := u.queries.GetUserByUsername(ctx, uname)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ListUsers implements Service.
func (u *UserService) ListUsers(ctx context.Context) ([]data.User, error) {
	users, err := u.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser implements Service.
func (u *UserService) UpdateUser(ctx context.Context, uur UserUpdateRequest) error {
	toUpdate, err := u.queries.GetUserByUsername(ctx, uur.Username)
	if err != nil {
		return err
	}
	err = u.queries.UpdateUser(ctx, data.UpdateUserParams{Username: uur.Username, Email: uur.Email, Password: toUpdate.Password, ID: toUpdate.ID})
	if err != nil {
		return err
	}
	return nil
}

func NewUserService(queries *data.Queries) Service {
	return &UserService{queries: queries}
}
