// Mapping profiles from interla to external representation
package main

import (
	"time"

	"github.com/acemouty/disney-pass/internal/database"
	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func databaseUserToUser(user database.User) User {
	return User{
		UserID:    user.UserID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
}

type ParentPost struct {
	PostId int32 `json:"id"`
}

func databasePostToPost(post database.ParentPost) ParentPost {
	return ParentPost{
		PostId: post.ID,
	}
}

// TODO: Figure out how to make this generic...
type ParkArea struct {
	AreaId   int32  `json:"areaId"`
	AreaName string `json:"areaName"`
}

func databaseParkAreaToParkArea(parkArea database.ParkArea) ParkArea {
	return ParkArea{
		AreaId:   parkArea.ID,
		AreaName: parkArea.Name,
	}
}

func databaseParkAreaSliceToParkAreaSlice(slice *[]database.ParkArea) []ParkArea {
	parkAreas := []ParkArea{}

	for _, parkArea := range *slice {
		parkAreas = append(parkAreas, databaseParkAreaToParkArea(parkArea))
	}

	return parkAreas
}

type AreaRide struct {
	AreaRideId   int32  `json:"areaRideId"`
	AreaRideName string `json:"areadRideName"`
}

func databaseAreaRideToAreaRide(areaRide database.AreaRide) AreaRide {
	return AreaRide{
		AreaRideId:   areaRide.ID,
		AreaRideName: areaRide.RideName,
	}
}

func databaseAreaRideSliceToAreaRideSlice(slice *[]database.AreaRide) []AreaRide {
	areaRides := []AreaRide{}

	for _, areaRide := range *slice {
		areaRides = append(areaRides, databaseAreaRideToAreaRide(areaRide))
	}

	return areaRides
}
