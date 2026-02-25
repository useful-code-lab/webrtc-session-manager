package service

import "context"

type SessionStorage interface {
	CreateSession(ctx context.Context, id, userId string) error
}

type SessionService struct {
	storage SessionStorage
}

func New(storage SessionStorage) *SessionService {
	return &SessionService{storage: storage}
}

func (s *SessionService) Start(ctx context.Context, id, userID string) error {
	return s.storage.CreateSession(ctx, id, userID)
}
