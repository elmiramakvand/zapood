package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"zapood/entities"

	"github.com/gorilla/mux"

	jwt "github.com/dgrijalva/jwt-go"
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
	// CheckToken := CheckJWTToken(*r, w)
	// if !CheckToken {
	// 	return
	// }

	cookie, err := r.Cookie("JWT_Token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("cookie error : %s", err)
		return
	}
	tokenString := cookie.Value
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(toke *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("token is invalid")
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println("err: %v", err)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("err: %v", err)
		return
	}

	var User []entities.User
	rows := userModel.DB.Find(&User)
	sqlRows, _ := rows.Rows()
	var users []entities.User
	for sqlRows.Next() {
		u := entities.User{}
		err2 := sqlRows.Scan(&u.ID, &u.Name, &u.Family, &u.UserName, &u.Password)
		if err2 != nil {
			fmt.Println(err2)
			return
		} else {
			users = append(users, u)
		}
	}
	json.NewEncoder(w).Encode(users)
	return
}

func (userModel UserModel) Operation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	operation, ok := vars["operation"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "operation not found!")
		return
	}
	var user entities.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "could not decode request body by error : %v", err)
		return
	}

	switch strings.ToLower(operation) {
	case "add":
		result := userModel.DB.Create(&user)
		if result.Error != nil {
			fmt.Fprintln(w, "Insert Error : %v", result.Error)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "edit":
		result := userModel.DB.Save(&user)
		if result.Error != nil {
			fmt.Fprintln(w, "Insert Error : %v", result.Error)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	return

}

func (userModel UserModel) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "id not found!")
		return
	}
	var User []entities.User
	result := userModel.DB.Delete(&User, id)
	if result.Error != nil {
		fmt.Fprintln(w, "Delete Error : %v", result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

func CheckJWTToken(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("JWTToken")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println(err)
		return false
	}
	tokenString := cookie.Value
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(toke *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("token is invalid")
		return false
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println("err: %v", err)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("err: %v", err)
		return false
	}
	return true
}
