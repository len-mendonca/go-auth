package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(password string) (string, error) {
	hashBytePassword, err := bcrypt.GenerateFromPassword([]byte(password), 10) //bcrypt works on bytes so convert string to bytes

	return string(hashBytePassword), err

}
