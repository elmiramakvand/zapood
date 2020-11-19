package models

import (
	"database/sql"
	"zapood/entities"
)

type UserModel struct {
	DB *sql.DB
}

func (userModel UserModel) FindAll() ([]entities.User, error) {
	rows, err := userModel.DB.Query("select * from [iranvar].[rose].[User]")
	if err != nil {
		return nil, err
	} else {
		var users []entities.User
		for rows.Next() {
			u := entities.User{}
			err2 := rows.Scan(&u.ID, &u.Name, &u.Family, &u.UserName, &u.Password)
			if err2 != nil {
				return nil, err2
			} else {
				users = append(users, u)
			}
		}
		return users, nil
	}
}
