package main

import (
	"net/http"
	"time"

	"github.com/mauricekoreman/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefreshAccessToken(w http.ResponseWriter, req *http.Request) {
	type response struct {
		Token string `json:"token"`
	}

	token, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "error getting token", err)
		return
	}

	user, err := cfg.db.GetUserFromRefreshToken(req.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't get user for refresh token", err)
		return
	}

	accessToken, err := auth.MakeJWT(user.ID, cfg.secret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't validate token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		Token: accessToken,
	})
}

func (cfg *apiConfig) handlerRevokeRefreshToken(w http.ResponseWriter, req *http.Request) {
	refreshToken, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "error getting token", err)
		return
	}

	_, err = cfg.db.RevokeRefreshToken(req.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error revoking refresh token", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
