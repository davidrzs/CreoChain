package chain

// The contents of this file have been taken from https://gowebexamples.com/password-hashing/

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes a password and returns a hashed version of it and an error if one occurred.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// HashPassword takes a password an already hashes password and returns if they are the same.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
