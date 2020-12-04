package restapi

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RunApi(endpoint string, DB *gorm.DB) error {
	r := mux.NewRouter()
	RunApiOnRouter(r, DB)
	return http.ListenAndServe(endpoint, r)
}

func RunApiOnRouter(r *mux.Router, DB *gorm.DB) {
	authHandler := NewAuthModel(DB)
	r.Methods("POST").Path("/api/auth/login").HandlerFunc(authHandler.Login)

	userHandler := NewUserModel(DB)
	r.Methods("GET").Path("/api/User/GetAllUsers").HandlerFunc(userHandler.GetAllUsers)
	r.Methods("POST").Path("/api/User/{operation:(?:add|edit)}").HandlerFunc(userHandler.Operation)
	r.Methods("POST").Path("/api/User/delete/{id}").HandlerFunc(userHandler.Delete)
}
