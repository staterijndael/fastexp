package sqlstore

import (
	"github.com/Oringik/fastexp/internal/app/store"
	"github.com/jmoiron/sqlx"

	// pq ...
	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	db             *sqlx.DB
	userRepository *UserRepository
}

// New ...
func New(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository

}
