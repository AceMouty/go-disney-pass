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
