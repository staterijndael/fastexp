package store

import "github.com/Oringik/fastexp/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	AddTags(int, []string) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
