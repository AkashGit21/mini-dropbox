package api

import (
	"net/http"

	"github.com/AkashGit21/typeface-assignment/utils"
	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {
	router := mux.NewRouter()
	dropboxRouter := router.PathPrefix("/api").Subrouter()
	dropboxRouter.Use(PanicRecoveryMiddleware)

	dropboxHandler(dropboxRouter)
	return router, nil
}

func PanicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				// Handle the panic
				utils.InfoLog("Panic recovered: ", r)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
