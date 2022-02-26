package utils

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(pwd string) (string, error) {
	pwdByte := []byte(pwd)

	hashedPwd, err := bcrypt.GenerateFromPassword(pwdByte, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPwd), nil
}

func DecryptPassowrd(hashed string, pwd string) bool {
	hashedByte := []byte(hashed)
	pwdByte := []byte(pwd)
	if err := bcrypt.CompareHashAndPassword(hashedByte, pwdByte); err != nil {
		return false
	}
	return true
}
