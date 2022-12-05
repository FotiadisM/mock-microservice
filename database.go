package main

import (
	"context"
	"errors"
	"sync"

	userv1 "github.com/findit-it/users-svc/api/user/v1"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type Database interface {
	GetUser(ctx context.Context, id string) (*userv1.User, error)
	StoreUser(ctx context.Context, user *userv1.User) error
}

type inMemoryDB struct {
	mu    sync.Mutex
	users map[string]*userv1.User
}

func newInMemoryDB() *inMemoryDB {
	return &inMemoryDB{
		users: make(map[string]*userv1.User),
	}
}

func (db *inMemoryDB) GetUser(_ context.Context, id string) (*userv1.User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	user, ok := db.users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (db *inMemoryDB) StoreUser(_ context.Context, user *userv1.User) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.users[user.GetId()]; ok {
		return ErrUserAlreadyExists
	}

	db.users[user.GetId()] = user
	return nil
}
