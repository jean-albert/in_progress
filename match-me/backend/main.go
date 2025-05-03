package main

import (
	"log"
	"match-me/backend/config"
	"match-me/backend/handlers"
	"match-me/backend/utils"
	"net/http"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := utils.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Set up handlers
	h := handlers.NewHandler(db, cfg.JWTSecret)

	// Set up routes
	http.HandleFunc("/register", h.Register)
	http.HandleFunc("/login", h.Login)
	http.HandleFunc("/logout", h.Logout)

	http.HandleFunc("/profile", h.AuthMiddleware(h.HandleProfile))
	http.HandleFunc("/bio", h.AuthMiddleware(h.HandleBio))

	http.HandleFunc("/recommendations", h.AuthMiddleware(h.GetRecommendations))
	http.HandleFunc("/connect", h.AuthMiddleware(h.HandleConnection))

	http.HandleFunc("/chat", h.AuthMiddleware(h.HandleChat))
	http.HandleFunc("/ws", h.AuthMiddleware(h.HandleWebSocket))

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
