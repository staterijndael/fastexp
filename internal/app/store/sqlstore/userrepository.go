package sqlstore

import (
	"database/sql"
	"log"
	"unicode/utf8"

	"github.com/Oringik/fastexp/internal/app/model"
	"github.com/Oringik/fastexp/internal/app/store"
	"github.com/lib/pq"
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
		"SELECT id, email, encrypted_password FROM users WHERE email=$1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
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

	var tags2 []*string

	if err := r.store.db.QueryRow("SELECT tags FROM users WHERE id=$1", userID).Scan(&tags2); err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}

	}

	for _, tag := range tags {

		err := r.store.db.QueryRow("SELECT id,msg FROM tags WHERE msg = $1", tag)
		if err == nil {
			return store.WrongUserTag
		}

		if utf8.RuneCountInString(tag) > 10 {
			return store.TagWrongLength
		}
		if utf8.RuneCountInString(tag) == 0 {
			return store.TagIsNull
		}

		var msg *string

		// rows, err := r.store.db.NamedQuery("INSERT INTO tags (msg) VALUES (:msg) RETURNING id", tag)

		// if rows.Next() {

		// }

		r.store.db.QueryRow(
			"INSERT INTO tags (msg) VALUES ($1) RETURNING msg",
			tag,
		).Scan(msg)

		tags2 = append(tags2, msg)

	}

	r.store.db.Exec("UPDATE users SET tags = $1 WHERE id = $2", pq.Array(tags2), userID)

	return nil
}

// GetTags ...
func (r *UserRepository) GetTags(userID int) ([]model.Tag, error) {

	var tags []string

	err := r.store.db.QueryRow("SELECT tags FROM users WHERE id = $1", userID).Scan(
		&tags,
	)
	if err != nil {
		log.Print(err)
		return nil, store.TagsNotfound
	}

	var modelsTags []model.Tag

	for _, tag := range tags {
		var id int
		var text string

		r.store.db.QueryRow("SELECT id,msg FROM tags WHERE id = $1", tag).Scan(
			&id,
			&text,
		)

		modelTag := model.Tag{
			ID:   id,
			Text: text,
		}

		modelsTags = append(modelsTags, modelTag)
	}

	return modelsTags, nil
}

// AddTags ...
func (r *UserRepository) AddThemeTags(themeID int, tags []string) error {

	if err := r.store.db.QueryRow("SELECT * FROM tagstheme WHERE id=$1").Scan(); err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}

	}

	for _, tag := range tags {

		err := r.store.db.QueryRow("SELECT * FROM theme WHERE id = $1", themeID)
		if err == nil {
			return store.WrongUserTag
		}

		if utf8.RuneCountInString(tag) > 10 {
			return store.TagWrongLength
		}
		if utf8.RuneCountInString(tag) == 0 {
			return store.TagIsNull
		}

		r.store.db.Exec("INSERT INTO tagstheme (title, description) VALUES ($1,$2)", themeID, tag)

	}

	return nil
}

// GetTags ...
func (r *UserRepository) GetThemeTags(themeID int) ([]model.TagTheme, error) {

	var tags []model.TagTheme

	rows, err := r.store.db.Query("SELECT theme_id,msg  FROM tagstheme WHERE theme_id = $1", themeID)
	if err != nil {
		return nil, store.TagsNotfound
	}

	for rows.Next() {
		var tag model.TagTheme

		rows.Scan(
			&tag.ThemeID,
			&tag.Text,
		)

		tags = append(tags, tag)
	}

	return tags, nil
}

// CreateTheme ...
func (r *UserRepository) CreateTheme(th *model.Theme) error {
	if err := th.ValidateTheme(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO theme (title, description) VALUES ($1,$2) RETURNING id",
		th.Title,
		th.Description,
	).Scan(&th.ID)

}

// AddUserTheme ...
func (r *UserRepository) AddUserTheme(userID int, th *model.Theme) error {

	var themesID []int

	err := r.store.db.QueryRow("SELECT themes FROM users WHERE id = ?", userID).Scan(
		&themesID,
	)
	if err != nil {
		return store.ErrRecordNotFound
	}

	for _, themeID := range themesID {
		if th.ID == themeID {
			return store.RepeatedValue
		}
	}

	themesID = append(themesID, th.ID)

	r.store.db.Exec("INSERT INTO users (themes) VALUES ($1)", themesID)

	return nil

}

// GetTheme ...
func (r *UserRepository) GetAllThemes() ([]model.Theme, error) {

	var themes []model.Theme

	rows, err := r.store.db.Query("SELECT id,title,description FROM theme")
	if err != nil {
		return nil, store.ThemeNotFound
	}

	for rows.Next() {
		var tag model.Theme

		rows.Scan(
			&tag.ID,
			&tag.Title,
			&tag.Description,
		)

		themes = append(themes, tag)
	}

	return themes, nil
}

func (r *UserRepository) CreateCard(card *model.Card) error {
	if err := card.ValidateCard(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO cards (name, shortdesc, fulldesc) VALUES ($1, $2) RETURNING id",
		card.Name,
		card.ShortDesc,
		card.FullDesc,
	).Scan(&card.ID)
}

func (r *UserRepository) DeleteCard(cardName string) {

	r.store.db.Exec(
		"DELETE FROM cards WHERE name = $1", cardName,
	)
}
