package models

import (
	"errors"

	"example.com/Course_Project/db"
	"example.com/Course_Project/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM user WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPass string
	err := row.Scan(&u.ID, &retrievedPass)

	if err != nil {
		return errors.New("credentials invalid")
	}

	validPassword := utils.CheckHashedPass(u.Password, retrievedPass)
	if !validPassword {
		return errors.New("credentials invalid")
	}

	return nil
}

func (u User) Save() error {
	query := `INSERT INTO user(email, password) VALUES (?, ?)`
	stm, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stm.Close()

	hashedPass, err := utils.HashPassword(u.Password) // We hash the password in order to save it to the database
	if err != nil {
		return err
	}

	result, err := stm.Exec(u.Email, hashedPass) // If we chose to save the password with "u.Password" as it is the data is insecure as a vulnerability may cause a leak of the information

	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()

	u.ID = userID
	return err
}
