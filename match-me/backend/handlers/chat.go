package handlers

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Handler) HandleChat(w http.ResponseWriter, r *http.Request) {
	userID := utils.GetUserIDFromContext(r.Context())

	switch r.Method {
	case "GET":
		chats, err := utils.GetChats(h.DB, userID)
		if err != nil {
			http.Error(w, "Failed to get chats", http.StatusInternalServerError)
			return
		}
		utils.RespondJSON(w, chats)

	case "POST":
		var msg models.Message
		if err := utils.ParseJSON(r, &msg); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if err := utils.SaveMessage(h.DB, userID, msg); err != nil {
			http.Error(w, "Failed to send message", http.StatusInternalServerError)
			return
		}
		utils.RespondJSON(w, map[string]string{"status": "success"})
	}
}

func (h *Handler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// WebSocket implementation
}
