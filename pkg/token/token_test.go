package token

import (
	"app/diplom/pkg/models"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCreateToken(t *testing.T) {
	user := models.User{
		ID:       102,
		Username: "alex_prett@gmail.com",
		Password: "Doe123Q!",
		FIO:      "Uliy Alexandr",
		RoleID:   1,
	}
	secretKey := "JWT_SALT"
	oldJwtSalt := os.Getenv(secretKey)
	defer os.Setenv(secretKey, oldJwtSalt)
	os.Setenv(secretKey, "verysecretsaltfortesting")

	actual, err := Create(&user)
	assert.NoError(t, err, "test failed")
	t.Run("CreateToken", func(t *testing.T) {
		assert.NotEmpty(t, actual, "token string should not be empty")
		assert.NoError(t, err, "should be no errors")
	})
}
func TestParseToken(t *testing.T) {
	user := models.User{
		ID:       10,
		Username: "John",
		Password: "Doedoe",
	}

	secretKey := "JWT_SALT"
	oldJwtSalt := os.Getenv(secretKey)
	defer os.Setenv(secretKey, oldJwtSalt)
	os.Setenv(secretKey, "verysecretsaltfortesting")

	testToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAsImV4cCI6MTYzNDgwNjU4MX0.JJEmvlrdYPTv2lAeE3wSrPAjb188fl9QbC0l3hU7LLE"
	actual, err := Parse(testToken)
	assert.NoError(t, err, "test failed")
	t.Run("ParseToken", func(t *testing.T) {
		assert.Equal(t, user.ID, actual, "ID should be equal")
		assert.NoError(t, err, "should be no errors")
	})
}