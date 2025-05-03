package handlers

import (
	"match-me-backend/models"
	"net/http"
)

func (h *Handler) HandleProfile(w http.ResponseWriter, r *http.Request) {
	userID := utils.GetUserIDFromContext(r.Context())

	switch r.Method {
	case "GET":
		profile, err := utils.GetProfile(h.DB, userID)
		if err != nil {
			http.Error(w, "Profile not found", http.StatusNotFound)
			return
		}
		utils.RespondJSON(w, profile)

	case "POST", "PUT":
		var profile models.Profile
		if err := utils.ParseJSON(r, &profile); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if err := utils.SaveProfile(h.DB, userID, profile); err != nil {
			http.Error(w, "Failed to save profile", http.StatusInternalServerError)
			return
		}
		utils.RespondJSON(w, map[string]string{"status": "success"})
	}
}

func (h *Handler) HandleBio(w http.ResponseWriter, r *http.Request) {
	// Similar structure to HandleProfile
}
