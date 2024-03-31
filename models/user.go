package models

import (
	"crud-events/db"
	"crud-events/utils"
	"errors"
	"fmt"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func CreateNewUser(user User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	insert, err := db.DB.Query(
		"INSERT INTO users (email,password) VALUES (?,?)",
		user.Email, hashedPassword)

	if err != nil {
		return err
	}
	defer insert.Close()
	return nil
}

func ValidateCredentials(user User) error {
	row := db.DB.QueryRow("SELECT id, password FROM users WHERE email=?", user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	fmt.Println(user.Password, retrievedPassword)
	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)
	fmt.Println(passwordIsValid)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
