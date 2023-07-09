// Mapping profiles from interla to external representation
package main

import (
	"time"

	"github.com/acemouty/disney-pass/internal/database"
	"github.com/google/uuid"
	"github.com/relvacode/iso8601"
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
	PostId       int32        `json:"id"`
	UserName     string       `json:"userName"`
	ParentName   string       `json:"parentName"`
	AreaName     string       `json:"areaName"`
	RideName     string       `json:"rideName"`
	IsOpen       bool         `json:"isOpen"`
	RideTime     iso8601.Time `json:"rideTime"`
	NumberOfKids int32        `json:"numberOfKids"`
}

func databasePostToFriendlyPost(post database.GetAllParentPostsRow) database.ParentPostFriendly {

}
func databaseFriendlyPostToFriendlyPost(post database.ParentPostFriendly) ParentPost {
	return ParentPost{
		PostId:       post.ID,
		UserName:     post.Username,
		ParentName:   post.Parentname,
		AreaName:     post.ParkAreaName,
		RideName:     post.RideName,
		NumberOfKids: post.NumberOfKids,
	}
}

func databaseParentPostSliceToParentPostSlice(slice *[]database.ParentPostFriendly) []ParentPost {
	parkAreas := make([]ParentPost, len(*slice))
	for idx, parkArea := range *slice {
		parkAreas[idx] = databaseFriendlyPostToFriendlyPost(parkArea)
	}

	return parkAreas
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
