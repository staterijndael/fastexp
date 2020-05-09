package sqlstore

import (
	"database/sql"
	"unicode/utf8"

	"github.com/Oringik/fastexp/internal/app/model"
	"github.com/Oringik/fastexp/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)

}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password, tags FROM users WHERE email=$1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
		&u.Tags,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
	}

	return u, nil
}

// Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password,tags FROM users WHERE id=$1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
		&u.Tags,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
	}

	return u, nil
}

// AddTags ...
func (r *UserRepository) AddTags(userID int, tags []string) error {

	if err := r.store.db.QueryRow("SELECT * FROM users WHERE id=$1").Scan(); err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}

	}

	for _, tag := range tags {

		err := r.store.db.QueryRow("SELECT user_id,text FROM tags WHERE user_id = $1 AND msg = $2", userID, tag)
		if err == nil {
			return store.WrongUserTag
		}

		if utf8.RuneCountInString(tag) > 10 {
			return store.TagWrongLength
		}
		if utf8.RuneCountInString(tag) == 0 {
			return store.TagIsNull
		}

		r.store.db.Exec("INSERT INTO tags (user_id, msg) VALUES ($1,$2)", userID, tag)

	}

	return nil
}
