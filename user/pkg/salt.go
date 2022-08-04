package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

// HashEncode /*
func HashEncode(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

// ComparePasswords /*
func ComparePasswords(hashedPwd string, sourcePwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(sourcePwd))
	if err != nil {
		return false
	}
	return true
}
