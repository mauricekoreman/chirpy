package main

import "net/http"

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, req *http.Request) {
	chirps, err := cfg.db.GetAllChirps(req.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed getting all chirps", err)
		return
	}

	allChirps := []Chirp{}

	for _, chirp := range chirps {
		allChirps = append(allChirps, Chirp{
			Body:      chirp.Body,
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			UserId:    chirp.UserID,
		})
	}

	respondWithJSON(w, http.StatusOK, allChirps)
}
