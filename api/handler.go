package api

import (
	"net/http"

	"github.com/AkashGit21/typeface-assignment/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type APIHandler struct {
	utils.MetadataOps
	utils.S3Ops
}

func NewAPIHandler() *APIHandler {
	persistenceDB, err := utils.NewPersistenceDBLayer()
	if err != nil {
		panic(err)
	}

	s3Client, err := utils.NewS3Client()
	if err != nil {
		panic(err)
	}

	return &APIHandler{
		persistenceDB,
		s3Client,
	}
}

func dropboxHandler(r *mux.Router) {
	dh := NewAPIHandler()
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Apply the CORS middleware to your router
	http.Handle("/", corsMiddleware.Handler(r))

	r.HandleFunc("/files/upload", dh.uploadFile).Methods("POST")
	r.HandleFunc("/files/{fileID}", dh.getFile).Methods("GET")
	r.HandleFunc("/files/{fileID}", dh.updateFile).Methods("PUT")
	r.HandleFunc("/files/{fileID}", dh.deleteFile).Methods("DELETE")
	r.HandleFunc("/files", dh.listFiles).Methods("GET")
}
