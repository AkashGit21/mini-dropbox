package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/AkashGit21/typeface-assignment/utils"
	"github.com/gorilla/mux"
)

// TODO: Uploads the file to blob storage (s3 here).
func (ah *APIHandler) uploadFile(w http.ResponseWriter, r *http.Request) {
	utils.DebugLog("inside uploadFile")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
}

// Fetch the file metadata from persistent storage (s3 here).
func (ah *APIHandler) getFile(w http.ResponseWriter, r *http.Request) {
	utils.DebugLog("inside getFile")

	vars := mux.Vars(r)
	id := vars["fileID"]

	w.Header().Add("Content-Type", "application/json")
	if utils.IsEmptyString(id) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(getFailureMessage(errors.New("unique id of file is required")))
		return
	}

	fileID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(getFailureMessage(err))
		return
	}

	data, err := ah.MetadataOps.GetRecord(fileID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getFailureMessage(err))
		return
	}
	if data == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getFailureMessage(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// TODO: Upload the new file to blob storage and update metadata with new url.
func (ah *APIHandler) updateFile(w http.ResponseWriter, r *http.Request) {
	utils.DebugLog("inside updateFile")

	w.WriteHeader(http.StatusNotImplemented)
}

// Soft deletes the file from blob storage.
func (ah *APIHandler) deleteFile(w http.ResponseWriter, r *http.Request) {
	utils.DebugLog("inside deleteFile")

	vars := mux.Vars(r)
	id := vars["fileID"]

	w.Header().Add("Content-Type", "application/json")
	if utils.IsEmptyString(id) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(getFailureMessage(errors.New("unique id of file is required")))
		return
	}

	fileID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(getFailureMessage(err))
		return
	}

	// Validate if record with the given id exists
	exists, err := ah.MetadataOps.Exists(fileID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getFailureMessage(err))
		return
	}

	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(getFailureMessage(errors.New("no such record with given id exists")))
		return
	}

	if err = ah.MetadataOps.DeactivateRecord(fileID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(getFailureMessage(err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(getSuccessMessage())
}
