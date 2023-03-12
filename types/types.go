package types

import (
	"time"

	"github.com/bwmarrin/discordgo"
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

// Internal response from Popplio
type IUser struct {
	ID            string           `json:"id"`
	Username      string           `json:"username"`
	Discriminator string           `json:"discriminator"`
	Avatar        string           `json:"avatar"`
	Bot           bool             `json:"bot"`
	Status        discordgo.Status `json:"status"`
	System        bool             `json:"system"`
	Nickname      string           `json:"nickname"`
	Guild         string           `json:"in_guild"`
}

type User struct {
	User            *IUser            `json:"user" db:"-"`
	Votes           []string          `json:"votes" db:"votes"`
	Rates           []string          `json:"rates" db:"-"` //TODO: Add this to the DB
	Banned          bool              `json:"banned" db:"banned"`
	Staff           bool              `json:"staff" db:"staff"`
	Premium         bool              `json:"premium" db:"premium"`
	Lifetime        bool              `json:"lifetime_premium" db:"lifetime_premium"`
	Notifications   map[string]string `json:"notifications" db:"notifications"`
	PremiumDuration time.Time         `json:"premium_duration" db:"premium_duration"`
}
