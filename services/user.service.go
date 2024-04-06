package service

import (
	db "github.com/SJ22032003/go-ems/db"
	model "github.com/SJ22032003/go-ems/models"
)

func CreateUserService(user *model.User) (*model.User, error) {

	result, err := db.Execute("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = id
	return user, nil
}

func FindOneUserByEmail(email string) (*model.User, error) {
	row := db.DB.QueryRow("SELECT * FROM users WHERE email = ? LIMIT 1", email)

	user := model.User{}

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil

}
