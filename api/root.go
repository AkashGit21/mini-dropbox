package api

import (
	"net/http"

	"github.com/AkashGit21/typeface-assignment/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {
	router := mux.NewRouter()

	// Enable CORS with the allowed origins you need.
	corsAllowedOrigins := handlers.AllowedOrigins([]string{"*"})

	// Enable CORS with other options as needed.
	cors := handlers.CORS(corsAllowedOrigins)

	dropboxRouter := router.PathPrefix("/api").Subrouter()
	dropboxRouter.Use(PanicRecoveryMiddleware)
	dropboxRouter.Use(cors)

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

func DeleteInactiveRecords() error {
	ah := NewAPIHandler()

	records, err := ah.MetadataOps.FetchInactiveRecords()
	if err != nil {
		return err
	}

	bucketName := utils.GetEnvValue("S3_BUCKET", "typeface-assignment")
	for _, record := range records {
		s3Key := getS3KeyFromURI(record.S3ObjectKey)
		if err = ah.S3Ops.DeleteObject(bucketName, s3Key); err != nil {
			utils.ErrorLog("unable to remove the s3 object with following details: ", record)
			continue
		}
	}
	return nil
}
