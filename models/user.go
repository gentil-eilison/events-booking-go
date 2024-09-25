package models

import (
	"errors"

	"github.com/gentil-eilison/events-booking-go/db"
	"github.com/gentil-eilison/events-booking-go/utils"
)

type User struct {
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := `
	INSERT INTO user(email, password) VALUES (?, ?);
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId
	return err
}

func (user User) ValidateCredentials() error {
	query := `SELECT password FROM user WHERE email = ?`
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}

	validPassword := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !validPassword {
		return errors.New("invalid credentials")
	}

	return nil
}