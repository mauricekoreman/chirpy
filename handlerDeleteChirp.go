package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/mauricekoreman/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, req *http.Request) {
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

	chirpId := req.PathValue("chirpId")

	chirpUUID, err := uuid.Parse(chirpId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error parsing chirp ID", err)
		return
	}

	chirp, err := cfg.db.GetChirp(req.Context(), chirpUUID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "error retrieving chirp", err)
		return
	}

	if chirp.UserID != userId {
		respondWithError(w, http.StatusForbidden, "error deleting chirp", err)
		return
	}

	err = cfg.db.DeleteChirp(req.Context(), chirpUUID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error deleting chirp", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
