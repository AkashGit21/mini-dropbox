package api

import (
	"github.com/AkashGit21/typeface-assignment/utils"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/mux"
)

type APIHandler struct {
	utils.MetadataOps
	s3Client *s3.S3
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

	r.HandleFunc("/files/upload", dh.uploadFile).Methods("POST")
	r.HandleFunc("/files/{fileID}", dh.getFile).Methods("GET")
	r.HandleFunc("/files/{fileID}", dh.updateFile).Methods("PUT")
	r.HandleFunc("/files/{fileID}", dh.deleteFile).Methods("DELETE")
	r.HandleFunc("/files", dh.listFiles).Methods("GET")
}
