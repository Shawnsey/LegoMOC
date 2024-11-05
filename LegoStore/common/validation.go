package common

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func ValidateParams(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        uuidParam := chi.URLParam(r, "id")
        // Validate UUID
		_, err := uuid.Parse(uuidParam) 
		if err != nil {
			fmt.Printf("Parsing of uuid error: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid UUID"))
			return
		}

        next.ServeHTTP(w, r)
    })
}