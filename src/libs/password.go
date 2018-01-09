package libs

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Pass string
}

var SecretKey string = "cairo"

func (w Password) Gen(p string) (string, error) {

	password := []byte(p + SecretKey)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	w.Pass = string(hashedPassword)

	return string(hashedPassword), nil
}

func (w Password) Compare(hs string, p string) (bool, error) {
	hash := []byte(hs)
	pass := []byte(p + SecretKey)
	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword(hash, pass)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (w Password) Token() (string, error) {
	t, err := w.Random(32)
	if err != nil {
		return "", err
	}

	return t, nil
}

func (w Password) RandomByte(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (w Password) Random(s int) (string, error) {
	b, err := w.RandomByte(s)
	return base64.URLEncoding.EncodeToString(b), err
}
