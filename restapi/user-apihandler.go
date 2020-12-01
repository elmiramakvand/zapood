package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zapood/entities"

	"gorm.io/gorm"
)

type UserModel struct {
	DB *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		DB: db,
	}
}

func (userModel UserModel) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var User []entities.User
	rows := userModel.DB.Find(&User)
	sqlRows, _ := rows.Rows()
	var users []entities.User
	for sqlRows.Next() {
		u := entities.User{}
		err2 := sqlRows.Scan(&u.ID, &u.Name, &u.Family, &u.UserName, &u.Password)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			users = append(users, u)
		}
	}
	json.NewEncoder(w).Encode(users)
	return
}
