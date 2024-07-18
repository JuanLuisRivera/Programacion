package utils

import "golang.org/x/crypto/bcrypt" // Package that allowss us to encrypt data

func HashPassword(password string) (string, error) { // We generate a new string with a hash password to store on the db
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // We import the "GenerateFromPassword" function from the package
	return string(bytes), err
}

func CheckHashedPass(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
