package storage

import (
	"context"

	"github.com/HsiaoCz/app-stone/session/types"
	"github.com/uptrace/bun"
)

type SessionStoreInter interface {
	CreateSession(context.Context, *types.Sessions) (*types.Sessions, error)
	GetSessionByToken(context.Context, string) (*types.Sessions, error)
	DeleteSessionByToken(context.Context, string) error
}

type SessionStore struct {
	db *bun.DB
}

func SessionStoreInit(db *bun.DB) *SessionStore {
	return &SessionStore{
		db: db,
	}
}

func (s *SessionStore) CreateSession(ctx context.Context, session *types.Sessions) (*types.Sessions, error) {
	return nil, nil
}

func (s *SessionStore) GetSessionByToken(ctx context.Context, token string) (*types.Sessions, error) {
	return nil, nil
}

func (s *SessionStore) DeleteSessionByToken(ctx context.Context, token string) error {
	return nil
}
