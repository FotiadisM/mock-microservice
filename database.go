package main

import (
	"context"
	"sync"

	userv1 "github.com/findit-it/users-svc/api/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserNotFound      = status.Error(codes.NotFound, "user not found")
	ErrUserAlreadyExists = status.Error(codes.AlreadyExists, "user already exists")
)

type inMemoryDB struct {
	mu    sync.Mutex
	users map[string]*userv1.User
}

func newInMemoryDB() *inMemoryDB {
	return &inMemoryDB{
		users: make(map[string]*userv1.User),
	}
}

func (db *inMemoryDB) Ping(_ context.Context) error {
	return nil
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
