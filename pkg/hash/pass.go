package hash

import (
	"golang.org/x/crypto/bcrypt"
)

//HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

//VerifyPassword checks the input password while verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) (isPasswordValid bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword)); err != nil {
		return false
	}
	return true

}
