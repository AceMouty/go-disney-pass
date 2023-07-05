package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/acemouty/disney-pass/internal/database"
	"github.com/acemouty/disney-pass/internal/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type requestBody struct {
	Username string
	Password string
}

func (cfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	req := requestBody{}

	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "createuser::err couldnt decode request body")
		return
	}

	hashedPassword, err := hashUserPassword(req.Password)
	if err != nil {
		log.Println("hashUserPassword::err::", err)
		respondWithError(w, http.StatusInternalServerError, "Couldnt create user")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		UserID:    uuid.New(),
		Username:  req.Username,
		Password:  hashedPassword,
		CreatedAt: time.Now().UTC(),
	})

	if err != nil {
		log.Println("db::createuser::err", err)
		respondWithError(w, http.StatusInternalServerError, "Couldnt create user")
		return
	}

	res := map[string]interface{}{
		"created_user": databaseUserToUser(user),
	}

	token, err := jwt.GenerateToken(&user)
	if err != nil {
		log.Println("generateToken::err:", err)
		respondWithError(w, http.StatusInternalServerError, "Couldnt create user")
		return
	}

	res["token"] = token

	respondWithJSON(w, http.StatusOK, res)
	return
}

func (cfg *apiConfig) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	req := requestBody{}

	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "createuser::err couldnt decode request body")
		return
	}

	dbUser, err := cfg.DB.GetUserByUsername(r.Context(), req.Username)
	if err != nil {
		log.Println("handleUserLogin::err::unable to find username:", req.Username)
		log.Println("db::err::", err)
		respondWithError(w, http.StatusInternalServerError, "unable to login user")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(req.Password))
	passwordIsNotMatching := err != nil
	if passwordIsNotMatching {
		log.Println("handleUserLogin::err::password-mismatch", err)
		respondWithError(w, http.StatusInternalServerError, "unable to login user")
		return
	}

	token, err := jwt.GenerateToken(&dbUser)
	if err != nil {
		log.Println("handleUserLogin::err::jwt-token", err)
		respondWithError(w, http.StatusInternalServerError, "unable to login user")
		return
	}

	res := map[string]interface{}{"token": token}
	respondWithJSON(w, http.StatusOK, res)
	return
}

func hashUserPassword(password string) (string, error) {
	hasedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hasedBytes), err
}
