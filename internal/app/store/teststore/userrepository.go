package teststore

import (
	"github.com/Oringik/fastexp/internal/app/model"
	"github.com/Oringik/fastexp/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[int]*model.User
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err

	}
	u.ID = len(r.users) + 1
	r.users[u.ID] = u

	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

// Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]

	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

// AddTags ...
func (r *UserRepository) AddTags(userID int, tags []string) error {
	// u, ok := r.users[userID]

	// if !ok {
	// 	return store.ErrRecordNotFound
	// }

	// for _, tag := range tags {

	// 	if _, ok2 := r.users[userID]; !ok2 {
	// 		return store.WrongUserTag
	// 	}

	// 	if utf8.RuneCountInString(tag) > 10 {
	// 		return store.TagWrongLength
	// 	}
	// 	if utf8.RuneCountInString(tag) == 0 {
	// 		return store.TagIsNull
	// 	}

	// 	tagStruct := model.Tag{
	// 		ID: userID,
	// 		Text:   tag,
	// 	}

	// 	u.Tags = append(u.Tags, tagStruct)
	// }

	return nil
}

// GetTags ...
func (r *UserRepository) GetTags(userID int) ([]model.Tag, error) {

	return nil, nil
}

// CreateTheme ...
func (r *UserRepository) CreateTheme(th *model.Theme) error {
	return nil

}

func (r *UserRepository) GetAllThemes() ([]model.Theme, error) {

	return nil, nil
}

// GetTags ...
func (r *UserRepository) GetThemeTags(themeID int) ([]model.TagTheme, error) {

	return nil, nil
}

// AddTags ...
func (r *UserRepository) AddThemeTags(themeID int, tags []string) error {

	return nil
}

func (r *UserRepository) AddUserTheme(userID int, th *model.Theme) error {
	return nil
}

func (r *UserRepository) CreateCard(card *model.Card) error {
	return nil
}

func (r *UserRepository) DeleteCard(cardID string) {

}
