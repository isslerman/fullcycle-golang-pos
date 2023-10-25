package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Boss", "boss@hugo.com", "123456")
	assert.Nil(t, err)                           //
	assert.NotNil(t, user)                       // user is not nil
	assert.NotEmpty(t, user.ID)                  // id not empty
	assert.NotEmpty(t, user.Password)            // password not empty
	assert.Equal(t, "John Boss", user.Name)      // name is equal
	assert.Equal(t, "boss@hugo.com", user.Email) // email is equal
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Boss", "boss@hugo.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidadePassword("123456"))  // must be true
	assert.False(t, user.ValidadePassword("333333")) // must be true
	assert.NotEqual(t, "123456", user.Password)
}
