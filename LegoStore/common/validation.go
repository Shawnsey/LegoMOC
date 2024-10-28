package common

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func ValidateParams(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("started: %s %s %s", r.Method, r.URL.Path, chi.URLParam(r,"id"))
        uuidParam := chi.URLParam(r, "id")
		fmt.Printf("uuid: %d", len(uuidParam))
        // Validate UUID
		_, err := uuid.Parse(uuidParam) 
		if err != nil {
			fmt.Printf("err is there: %s", err)
			w.Write([]byte("Invalid UUID"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

        next.ServeHTTP(w, r)
    })
}