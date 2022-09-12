package teststore

import (
	"Rest/internal/app/model"
	"Rest/internal/app/store"
	"database/sql"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
