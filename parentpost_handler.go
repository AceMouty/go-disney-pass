package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/acemouty/disney-pass/internal/database"
	"github.com/google/uuid"
	"github.com/relvacode/iso8601"
)

type newParentPost struct {
	UserID       uuid.UUID `json:"userId"`
	AreaID       int32     `json:"areaId"`
	RideID       int32     `json:"rideId"`
	IsOpen       bool      `json:"isOpen"`
	RideTime     string    `json:"rideTime"` // ISO 8601 time string
	NumberOfKids int32     `json:"numberOfKids"`
}

func (cfg *apiConfig) handleGetParkInformation(w http.ResponseWriter, r *http.Request) {
	dbParkAreas, err := cfg.DB.GetAllParkAreas(r.Context())
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "handleGetParkInformation::err unable to get park areas")
		return
	}

	dbAreaRides, err := cfg.DB.GetAllAreaRides(r.Context())
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "handleGetParkInformation::err unable to get area rides")
		return
	}

	parkAreas := databaseParkAreaSliceToParkAreaSlice(&dbParkAreas)
	areaRides := databaseAreaRideSliceToAreaRideSlice(&dbAreaRides)

	res := map[string]interface{}{
		"parkAreas": parkAreas,
		"parkRides": areaRides,
	}

	respondWithJSON(w, http.StatusOK, res)
	return
}

func (cfg *apiConfig) handleGetAllPosts(w http.ResponseWriter, r *http.Request) {
	dbPosts, err := cfg.DB.GetAllParentPosts(r.Context())
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "handleGetAllPosts::err unable to get posts")
		return
	}
}

func (cfg *apiConfig) handleCreateParentPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	req := newParentPost{}

	err := decoder.Decode(&req)
	if err != nil {
		log.Print(err)
		respondWithError(w, http.StatusInternalServerError, "handleCreateParentPost::err unable to create post")
		return
	}

	rideTime, err := iso8601.ParseString(req.RideTime)
	if err != nil {
		log.Print(err)
		respondWithError(w, http.StatusInternalServerError, "handleCreateParentPost::err unable to create post")
		return
	}

	newPost, err := cfg.DB.CreateParentPost(r.Context(), database.CreateParentPostParams{
		UserID:       req.UserID,
		AreaID:       sql.NullInt32{Int32: req.AreaID},
		RideID:       sql.NullInt32{Int32: req.RideID},
		IsOpen:       req.IsOpen,
		RideTime:     rideTime,
		NumberOfKids: req.NumberOfKids,
	})
	if err != nil {
		log.Print("handleCreateParentPost::sql::err", err)
		respondWithError(w, http.StatusInternalServerError, "handleCreateParentPost::err unable to create post")
		return
	}

	mappedpost := databasePostToPost(newPost)
	res := map[string]interface{}{
		"createdPost": mappedpost.PostId,
	}
	respondWithJSON(w, http.StatusCreated, res)
}
