package helper

import "golang.org/x/crypto/bcrypt"

func BcryptGenerate(password string) (string, error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPassword := string(hashedPasswordBytes)
	if err != nil {
		return hashedPassword, err
	} else {
		return hashedPassword, nil
	}
}

func BcryptValidate(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
