package main

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, req *http.Request) {
	chirps, err := cfg.db.GetAllChirps(req.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed retrieving chirps", err)
		return
	}

	authorID := uuid.Nil
	authorIdString := req.URL.Query().Get("author_id")
	if authorIdString != "" {
		authorID, err = uuid.Parse(authorIdString)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "error parsing authorId", err)
			return
		}
	}

	allChirps := []Chirp{}
	for _, chirp := range chirps {
		if authorID != uuid.Nil && chirp.UserID != authorID {
			continue
		}

		allChirps = append(allChirps, Chirp{
			Body:      chirp.Body,
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			UserId:    chirp.UserID,
		})
	}

	sortString := req.URL.Query().Get("sort")
	if sortString == "desc" {
		sort.Slice(allChirps, func(i, j int) bool { return allChirps[i].CreatedAt.After(allChirps[j].CreatedAt) })
	}

	respondWithJSON(w, http.StatusOK, allChirps)
}
