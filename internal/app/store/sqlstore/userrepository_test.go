package sqlstore_test

import (
	"testing"

	"github.com/Oringik/fastexp/internal/app/model"
	"github.com/Oringik/fastexp/internal/app/store"
	"github.com/Oringik/fastexp/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))

	assert.NotNil(t, u)

}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_AddTag(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("tags")

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	s.User().Create(u1)
	tags := []string{"asdasdasd", "asdasdads", "212kkasd", "sadasd"}
	err := s.User().AddTags(u1.ID, tags)
	assert.NoError(t, err)
}

func TestUserRepository_GetTags(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("tags", "users")
	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	err := s.User().Create(u1)
	assert.NoError(t, err)
	tags1 := []string{"asdasdasd", "asdasdads", "212kkasd", "sadasd"}
	err = s.User().AddTags(u1.ID, tags1)
	assert.NoError(t, err)
	tags2, err := s.User().GetTags(u1.ID)
	assert.NoError(t, err)

	var readyTags []string

	for _, tag := range tags2 {
		readyTags = append(readyTags, tag.Text)
	}

	assert.Equal(t, tags1, readyTags)
}
