package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)



func writeJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	res, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Failed to marshal JSON:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(res)
	if err != nil {
		fmt.Println("Failed to write response:", err)
	}
}

// todo: rework this to be a middleware that just does the validation part.
func getValidUuid(r *http.Request) uuid.UUID {
	id := chi.URLParam(r, "id")
	parsedUUID, err:= uuid.Parse(id)
	println("error",err)
	return parsedUUID
}