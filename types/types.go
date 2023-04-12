package types

import (
	"time"

	"github.com/infinitybotlist/eureka/dovewing"
)

// This represents a IBL Popplio API Error
type ApiError struct {
	Context any    `json:"context,omitempty"`
	Message string `json:"message"`
	Error   bool   `json:"error"`
}

type TargetType int

const (
	TargetTypeUser TargetType = iota
	TargetTypeServer
)

type User struct {
	User            *dovewing.DiscordUser `json:"user" db:"-"`
	Votes           []string              `json:"votes" db:"votes"`
	Rates           []string              `json:"rates" db:"-"` //TODO: Add this to the DB
	Banned          bool                  `json:"banned" db:"banned"`
	Staff           bool                  `json:"staff" db:"staff"`
	Premium         bool                  `json:"premium" db:"premium"`
	Lifetime        bool                  `json:"lifetime_premium" db:"lifetime_premium"`
	Notifications   map[string]string     `json:"notifications" db:"notifications"`
	PremiumDuration time.Time             `json:"premium_duration" db:"premium_duration"`
}
