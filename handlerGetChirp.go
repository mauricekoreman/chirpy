package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, req *http.Request) {
	chirpId := req.PathValue("chirpId")

	// Convert chirpId to uuid.UUID
	chirpUUID, err := uuid.Parse(chirpId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong...", err)
	}

	chirp, err := cfg.db.GetChirp(req.Context(), chirpUUID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong retrieving the chirp", err)
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		Body:      chirp.Body,
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		UserId:    chirp.UserID,
	})
}
