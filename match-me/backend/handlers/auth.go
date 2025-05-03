package handlers

import (
	"database/sql"
	"match-me-backend/models"
	"match-me-backend/utils"
	"net/http"
)

type Handler struct {
	DB        *sql.DB
	JWTSecret string
}

func NewHandler(db *sql.DB, jwtSecret string) *Handler {
	return &Handler{DB: db, JWTSecret: jwtSecret}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := utils.ParseJSON(r, &user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Validate and create user
	userID, err := utils.CreateUser(h.DB, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	token, err := utils.GenerateToken(userID, h.JWTSecret)
	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, map[string]string{"token": token})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// Similar implementation to Register
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	// Clear token logic
}

func (h *Handler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		userID, err := utils.ValidateToken(token, h.JWTSecret)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		r = r.WithContext(utils.SetUserIDInContext(r.Context(), userID))
		next(w, r)
	}
}
