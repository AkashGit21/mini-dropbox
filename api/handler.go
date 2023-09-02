package api

import (
	"github.com/AkashGit21/typeface-assignment/utils"
	"github.com/gorilla/mux"
)

type APIHandler struct {
	utils.MetadataOps
	// s3object
}

func NewAPIHandler() *APIHandler {
	persistenceDB, err := utils.NewPersistenceDBLayer()
	if err != nil {
		panic(err)
	}

	return &APIHandler{
		persistenceDB,
	}
}

func dropboxHandler(r *mux.Router) {
	dh := NewAPIHandler()

	r.HandleFunc("/files/upload", dh.uploadFile).Methods("POST")
	r.HandleFunc("/files/{fileID}", dh.getFile).Methods("GET")
	r.HandleFunc("/files/{fileID}", dh.updateFile).Methods("PUT")
	r.HandleFunc("/files/{fileID}", dh.deleteFile).Methods("DELETE")
}
