package tools

import "golang.org/x/crypto/bcrypt"

func BcryptHash(pwd string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(bytes)
}

func BcryptCheck(pwd string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
