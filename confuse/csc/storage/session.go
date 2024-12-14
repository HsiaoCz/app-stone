package storage

import (
	"context"

	"github.com/HsiaoCz/app-stone/confuse/csc/types"
	"gorm.io/gorm"
)

type SessionStorer interface {
	CreateSession(context.Context, *types.Sessions) error
	GetSessionByToken(context.Context, string) (*types.Sessions, error)
	DeleteSessionByToken(context.Context, string) error
}

type SessionStore struct {
	db *gorm.DB
}

func SessionStoreInit(db *gorm.DB) *SessionStore {
	return &SessionStore{
		db: db,
	}
}

func (s *SessionStore) CreateSession(ctx context.Context, session *types.Sessions) error {
	tx := s.db.Debug().WithContext(ctx).Model(&types.Sessions{}).Create(session)
	return tx.Error
}

func (s *SessionStore) GetSessionByToken(ctx context.Context, token string) (*types.Sessions, error) {
	var session types.Sessions
	tx := s.db.Debug().WithContext(ctx).Model(&types.Sessions{}).Where("token = ?", token).First(&session)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &session, nil
}

func (s *SessionStore) DeleteSessionByToken(ctx context.Context, token string) error {
	var session types.Sessions
	return s.db.Debug().WithContext(ctx).Model(&types.Sessions{}).Where("token = ?", token).Delete(&session).Error
}
