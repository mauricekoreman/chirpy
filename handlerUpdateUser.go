package main

import (
	"encoding/json"
	"net/http"

	"github.com/mauricekoreman/chirpy/internal/auth"
	"github.com/mauricekoreman/chirpy/internal/database"
)

func (cfg *apiConfig) handlerUpdateUser(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		User
	}

	token, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "user not authenticated", err)
		return
	}

	userId, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid access token", err)
		return
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error parsing the parameters", err)
		return
	}

	newHashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error creating new password", err)
		return
	}

	updatedUser, err := cfg.db.UpdateUserById(req.Context(), database.UpdateUserByIdParams{
		ID:             userId,
		Email:          params.Email,
		HashedPassword: newHashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error updating user", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:        updatedUser.ID,
			Email:     updatedUser.Email,
			CreatedAt: updatedUser.CreatedAt,
			UpdatedAt: updatedUser.UpdatedAt,
		},
	})
}
