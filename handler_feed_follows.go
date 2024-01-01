package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Leon-dev45/go_backend/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintln("Error parsing JSON: ", err))
		return
	}
	feed, e := apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if e != nil {
		responseWithError(w, 400, fmt.Sprintln("You have already followed this feed"))
		return
	}

	responseWithJSON(w, 201, databaseUserToFeedFollows(feed))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feed_follows, e := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if e != nil {
		responseWithError(w, 400, fmt.Sprintln("Could not get the feeds that you followed"))
		return
	}

	responseWithJSON(w, 201, databasedatabaseUserToFeedFollowsToFeedFollow(feed_follows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedID := chi.URLParam(r, "feedFollowID")

	feedFollowID, err := uuid.Parse(feedID)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintln("You are not following this feed"))
		return
	}
	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintln("Server cannot unfollow the feed"))
		return
	}
	responseWithJSON(w, 200, struct{}{})

}
