package token

import (
	"time"
)

// Maker is interface for managing tokens
type Maker interface {

	// CreateToken create a new token for a specific user,
	// @param userID string is the user ID for claims,
	// @param duration time.Duration is for expired duration
	CreateToken(userID string, duration time.Duration) (string, error)

	// VerifyToken check if the token is valid or not,
	// @param signedToken string is the token from request
	VerifyToken(signedToken string) (*Payload, error)
}
