package api

import "encoding/json"

var (
	SUCCESS_MSG = map[string]interface{}{
		"status": "success",
	}
	FAILURE_MSG = map[string]interface{}{
		"status": "failure",
		"error":  nil,
	}
)

func getSuccessMessage() []byte {
	data, err := json.Marshal(SUCCESS_MSG)
	if err != nil {
		return nil
	}
	return data
}

func getFailureMessage(err error) []byte {
	// Append error to failure message
	FAILURE_MSG["error"] = err.Error()

	data, err := json.Marshal(FAILURE_MSG)
	if err != nil {
		return nil
	}
	return data
}
