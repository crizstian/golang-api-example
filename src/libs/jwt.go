package libs

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWT struct {
	UserName string
}

func (j JWT) Secret() []byte {
	return []byte("cairo")
}

// MapClaims creates a JSON web token
func (j JWT) MapClaims(email string, id uint) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	now := time.Now()
	secs := now.Unix()

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["role"] = 1
	claims["iat"] = secs
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["user_id"] = id

	// Generate encoded token and send it as response.
	t, err := token.SignedString(j.Secret())
	return t, err
}

func (j JWT) LargeMapClaims(email string, id uint) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	now := time.Now()
	secs := now.Unix()

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["role"] = 1
	claims["iat"] = secs
	claims["exp"] = time.Now().Add(time.Hour * 144).Unix()
	claims["user_id"] = id

	// Generate encoded token and send it as response.
	t, err := token.SignedString(j.Secret())
	return t, err
}
