package store

import (
	"os"
	"testing"

	"github.com/icrowley/fake"
	"github.com/poudels14/Neptro/utils"
	"github.com/stretchr/testify/assert"
)

func getTestUserStore() (UserStore, error) {
	utils.LoadAllKeys()

	dbUser := os.Getenv("TEST_DB_USER")
	dbPass := os.Getenv("TEST_DB_PASS")
	dbName := os.Getenv("TEST_DB_NAME")

	return InitializeUserDb(dbUser, dbPass, dbName)
}

func getUserFixture() User {
	return User{
		FirstName:    fake.FirstName(),
		LastName:     fake.LastName(),
		Email:        fake.EmailAddress(),
		Phone:        fake.Phone(),
		PasswordHash: fake.SimplePassword(),
		CountryCode:  "us",
	}
}

func TestCreatingUser(t *testing.T) {
	store, err := getTestUserStore()
	assert.Nil(t, err)

	u := getUserFixture()

	id, err := store.Create(u)
	assert.Nil(t, err)
	assert.NotZero(t, id, "User was not created properly")
}

func TestUpdatingUser(t *testing.T) {
	store, err := getTestUserStore()
	assert.Nil(t, err)

	u := getUserFixture()
	id, err := store.Create(u)
	assert.Nil(t, err)

	u.ID = id
	u.FirstName = "Elmo"

	err = store.Update(u)
	updated, _ := store.Get(id)

	assert.Nil(t, err, "Updating user failed")
	assert.NotNil(t, updated, "Did not retrieve user with id %d", id)
	assert.Equal(t, "Elmo", updated.FirstName)
}

func TestDeletingUser(t *testing.T) {
	store, err := getTestUserStore()
	assert.Nil(t, err)

	id, _ := store.Create(getUserFixture())
	err = store.Deactivate(id)
	assert.Nil(t, err)

	// TODO: expand this
}

func TestGettingUser(t *testing.T) {
	store, err := getTestUserStore()
	assert.Nil(t, err)

	u := getUserFixture()
	id, _ := store.Create(u)
	user, err := store.Get(id)
	assert.Nil(t, err)

	assert.Equal(t, id, user.ID, "Did not retrieve user correcly")
}

func TestGettingUnknownUser(t *testing.T) {
	store, err := getTestUserStore()
	assert.Nil(t, err)

	user, err := store.Get(9001)
	assert.Nil(t, err)
	assert.Nil(t, user)
}

func testCountingUsers(t *testing.T) {
	store, err := getTestUserStore()
	assert.Nil(t, err)

	oldCount, err := store.Count()
	assert.Nil(t, err)

	store.Create(getUserFixture())
	newCount, err := store.Count()
	assert.Nil(t, err)

	assert.Equal(t, 1, newCount-oldCount, "Did not receive count properly")
}
