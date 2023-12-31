// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type AreaRide struct {
	ID       int32
	AreaID   sql.NullInt32
	RideName string
}

type ParentPost struct {
	ID           int32
	UserID       uuid.UUID
	ParentID     uuid.NullUUID
	AreaID       sql.NullInt32
	RideID       sql.NullInt32
	IsOpen       bool
	RideTime     time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	NumberOfKids int32
}

type ParkArea struct {
	ID   int32
	Name string
}

type Role struct {
	ID        int32
	RoleName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	UserID    uuid.UUID
	Username  string
	Password  string
	CreatedAt time.Time
}

type UserInRole struct {
	ID        int32
	UserID    uuid.NullUUID
	RoleID    sql.NullInt32
	CreatedAt time.Time
	UpdatedAt time.Time
}
