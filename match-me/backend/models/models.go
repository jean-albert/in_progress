package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Profile struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	AboutMe   string `json:"about_me"`
	Picture   string `json:"picture"`
	Location  string `json:"location"`
}

type Bio struct {
	Interests   []string `json:"interests"`
	LookingFor  string   `json:"looking_for"`
	Age         int      `json:"age"`
	Gender      string   `json:"gender"`
	MaxDistance int      `json:"max_distance"`
}

type Connection struct {
	UserID int    `json:"user_id"`
	Status string `json:"status"`
}

type Message struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	SenderID  int    `json:"sender_id"`
	CreatedAt string `json:"created_at"`
}
