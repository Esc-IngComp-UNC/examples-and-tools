// Package helpers provides helper functions for handling HTTP responses and errors.
package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	appErrors "apirestful-go/pkg/errors"
)

// RespondWithError sends a JSON error response to the client with the provided AppError.
//
// Parameters:
//
//	w (http.ResponseWriter): The HTTP response writer to send the response.
//	err (appErrors.AppError): The AppError containing the error code and message to send.
//
// Behavior:
//
//	RespondWithError sets the "Content-Type" header to "application/json",
//	writes the HTTP status code from err.Code to the response writer,
//	and encodes a JSON object containing {"error": err.Message}.
//	If there is an error encoding the JSON, it logs the error to the standard logger.
func RespondWithError(w http.ResponseWriter, err *appErrors.AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)

	var response map[string]string
	msg := make(map[string]interface{})
	if unmarshalErr := json.Unmarshal([]byte(err.Message), &response); unmarshalErr != nil {
		msg["error"] = err.Message
	} else {
		msg["error"] = response
	}

	if encodeErr := json.NewEncoder(w).Encode(msg); encodeErr != nil {
		log.Printf("Error encoding JSON response: %v", encodeErr)
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
	}
}

// RespondWithJSON sends a JSON response with the given status code and payload.
//
// Parameters:
//
//	w (http.ResponseWriter): The HTTP response writer to send the response.
//	code (int): The HTTP status code to send in the response.
//	payload (interface{}): The data to encode as JSON and send in the response body.
//
// Behavior:
//
//	RespondWithJSON sets the "Content-Type" header to "application/json",
//	writes the provided HTTP status code to the response writer,
//	and encodes the payload interface{} as JSON.
//	If there is an error encoding the JSON, it logs the error to the standard logger.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(map[string]any{
		"data": payload,
	}); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

// RespondWithoutData sends a JSON response with the given status code and no data.
//
// Parameters:
//
//	w (http.ResponseWriter): The HTTP response writer to send the response.
//	code (int): The HTTP status code to send in the response.
//
// Behavior:
//
//	RespondWithoutData sets the "Content-Type" header to "application/json",
func RespondWithoutData(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}
