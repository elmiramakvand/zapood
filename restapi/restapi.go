package restapi

import (
	"log"
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

	ManageAuthRoutes(r, DB)

	ManageUserRoutes(r, DB)
}

func ManageAuthRoutes(r *mux.Router, DB *gorm.DB) {
	Handler := NewAuthModel(DB)
	AuthRoute := r.PathPrefix("/api/auth").Subrouter()
	AuthRoute.Methods("POST").Path("/login").HandlerFunc(Handler.Login)
}

func ManageUserRoutes(r *mux.Router, DB *gorm.DB) {
	Handler := NewUserModel(DB)
	UserRoutes := r.PathPrefix("/api/User").Subrouter()
	UserRoutes.Use(Authentication) // middleware fuction runs before api actions
	UserRoutes.Methods("GET").Path("/GetAllUsers").HandlerFunc(Handler.GetAllUsers)
	UserRoutes.Methods("POST").Path("/{operation:(?:add|edit)}").HandlerFunc(Handler.Operation)
	UserRoutes.Methods("POST").Path("/delete/{id}").HandlerFunc(Handler.Delete)
}

func Authentication(next http.Handler) http.Handler {

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if info, logedin := IsLogedin(r); logedin {
			// We found the token in our map
			log.Printf("Authenticated user %v\n", info)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "شما به این صفحه دسترسی ندارید", http.StatusForbidden)
		}
	})

	return h
}
