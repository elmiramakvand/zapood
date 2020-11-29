package restapi

import (
	"database/sql"
	"net/http"
)

type UserModel struct {
	DB *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{
		DB: db,
	}
}

func (userModel UserModel) GetAllUsers(w http.ResponseWriter, r *http.Request) {

}
