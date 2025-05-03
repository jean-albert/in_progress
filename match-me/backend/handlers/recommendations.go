package handlers

import (
	"match-me-backend/matching"
	"net/http"
)

func (h *Handler) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	userID := utils.GetUserIDFromContext(r.Context())

	bio, err := utils.GetBio(h.DB, userID)
	if err != nil {
		http.Error(w, "Bio not found", http.StatusNotFound)
		return
	}

	potentialMatches, err := utils.GetPotentialMatches(h.DB, userID)
	if err != nil {
		http.Error(w, "Failed to get matches", http.StatusInternalServerError)
		return
	}

	recommendations := matching.GetTopMatches(bio, potentialMatches, 10)
	utils.RespondJSON(w, recommendations)
}

func (h *Handler) HandleConnection(w http.ResponseWriter, r *http.Request) {
	// Handle connection requests
}
