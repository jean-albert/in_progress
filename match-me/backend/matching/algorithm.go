package matching

import (
	"match-me-backend/models"
	"sort"
)

type UserScore struct {
	UserID int
	Score  float64
}

// CalculateMatchScore computes compatibility between two user profiles
func CalculateMatchScore(user1, user2 models.Bio) float64 {
	score := 0.0

	// 1. Check gender preferences (50% weight)
	if user1.LookingFor != "" && user1.LookingFor != user2.Gender {
		return 0
	}
	if user2.LookingFor != "" && user2.LookingFor != user1.Gender {
		return 0
	}
	score += 0.5 // Basic compatibility if genders match preferences

	// 2. Compare interests (30% weight)
	commonInterests := 0
	for _, i1 := range user1.Interests {
		for _, i2 := range user2.Interests {
			if i1 == i2 {
				commonInterests++
				break
			}
		}
	}
	score += float64(commonInterests) / float64(len(user1.Interests)) * 0.3

	// 3. Age compatibility (10% weight)
	ageDiff := abs(user1.Age - user2.Age)
	if ageDiff <= 5 {
		score += 0.1
	} else if ageDiff <= 10 {
		score += 0.05
	}

	// 4. Location (10% weight)
	if user1.Location == user2.Location {
		score += 0.1
	}

	return clamp(score, 0, 1)
}

// GetTopMatches returns best matches for the current user
func GetTopMatches(currentUser models.Bio, users []models.Bio, limit int) []UserScore {
	var scores []UserScore

	for _, user := range users {
		if user.UserID == currentUser.UserID {
			continue // Skip self
		}

		score := CalculateMatchScore(currentUser, user)
		if score > 0.3 { // Only include matches with score > 30%
			scores = append(scores, UserScore{
				UserID: user.UserID,
				Score:  score,
			})
		}
	}

	// Sort by score (highest first)
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	// Limit results
	if len(scores) > limit {
		return scores[:limit]
	}
	return scores
}

// Helper functions
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
