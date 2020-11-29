package restapi

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func RunApi(endpoint string, DB *sql.DB) error {
	r := mux.NewRouter()
	RunApiOnRouter(r, DB)
	return http.ListenAndServe(endpoint, r)
}

func RunApiOnRouter(r *mux.Router, DB *sql.DB) {
	handler := NewUserModel(DB)
	apiRouter := r.PathPrefix("api/User/").Subrouter()
	apiRouter.Methods("GET").Path("GetAllUsers").HandlerFunc(handler.GetAllUsers)
}
