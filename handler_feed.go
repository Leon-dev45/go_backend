package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Leon-dev45/go_backend/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintln("Error parsing JSON: ", err))
		return
	}
	feed, e := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if e != nil {
		responseWithError(w, 400, fmt.Sprintln("Could not create a user"))
		return
	}

	responseWithJSON(w, 201, databaseUserToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeed(w http.ResponseWriter, r *http.Request) {

	feeds, e := apiCfg.DB.GetFeeds(r.Context())
	if e != nil {
		responseWithError(w, 400, fmt.Sprintln("Could not get feeds"))
		return
	}

	responseWithJSON(w, 201, databaseFeedsToFeed(feeds))
}
