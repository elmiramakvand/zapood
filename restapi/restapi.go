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
	handler := NewUserModel(DB)
	apiRouter := r.PathPrefix("/api/User/").Subrouter()
	apiRouter.Methods("GET").Path("/GetAllUsers").HandlerFunc(handler.GetAllUsers)
}
