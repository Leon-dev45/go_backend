package main

import (
	"fmt"
	"net/http"

	"github.com/Leon-dev45/go_backend/internal/auth"
	"github.com/Leon-dev45/go_backend/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middleWareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, e := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if e != nil {
			responseWithError(w, 400, fmt.Sprintf("User not found: %v", err))
			return
		}

		handler(w, r, user)
	}
}
