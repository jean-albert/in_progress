package matching

import "match-me-backend/models"

func GetTopMatches(currentUser models.Bio, users []models.Bio, limit int) []UserScore {
	var scores []UserScore

	for _, user := range users {
		if user.UserID == currentUser.UserID { // Now using UserID
			continue
		}

		score := calculateScore(currentUser, user)
		scores = append(scores, UserScore{
			UserID: user.UserID, // Now using UserID
			Score:  score,
		})
	}

	// Sort and return top matches
	return []int{} // simplified
}

func calculateScore(a, b models.Bio) float64 {
	score := 0.0
	// Implement matching logic
	return score
}
