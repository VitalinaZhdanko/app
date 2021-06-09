package token

import (
	config "app/diplom/pkg/configs"
	"app/diplom/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var conf = config.New()

type tokenClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

// Create returns token string with given user's ID
func Create(u *models.User) (string, error) {
	claims := tokenClaims{
		ID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.JwtSalt))
}

// Parse extracts ID from JWT string
func Parse(token string) (int, error) {
	tk, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.JwtSalt), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := tk.Claims.(*tokenClaims); ok && tk.Valid {
		return claims.ID, nil
	}
	return 0, err
}

