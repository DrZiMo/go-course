package models

import (
	"errors"
	"rest_api/db"
	"rest_api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func GetUsers() ([]User, error) {
	query := `SELECT * FROM users`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetSingleUser(id int64) (*User, error) {
	query := `SELECT * FROM users WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u User) Save() error {
	query := `
	INSERT INTO users (email, password)
	VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) Login() error {
	query := `SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)

	if err != nil {
		return err
	}

	validPassword := utils.CheckPasswordHash(u.Password, retrivedPassword)

	if !validPassword {
		return errors.New("Credentials invalid")
	}

	return nil
}

func (u User) Update(id int64) error {
	query := `
	UPDATE users
	SET email = ?, password = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Email, u.Password, id)

	return err
}

func (u User) Delete(id int64) error {
	query := `
	DELETE FROM users
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)

	return err
}
