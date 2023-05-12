package types

import (
	"time"

	"github.com/infinitybotlist/eureka/dovewing"
	"github.com/jackc/pgx/v5/pgtype"
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
	User            *dovewing.DiscordUser `json:"user" db:"-" description:"The user object."`
	Votes           []string              `json:"votes" db:"votes"`
	Rates           []string              `json:"rates" db:"-"` //TODO: Add this to the DB
	Banned          bool                  `json:"banned" db:"banned"`
	Staff           bool                  `json:"staff" db:"staff"`
	Premium         bool                  `json:"premium" db:"premium"`
	Lifetime        bool                  `json:"lifetime_premium" db:"lifetime_premium"`
	Notifications   map[string]string     `json:"notifications" db:"notifications"`
	PremiumDuration time.Time             `json:"premium_duration" db:"premium_duration"`
}

type CV struct {
	User      *dovewing.DiscordUser `json:"user" db:"-" description:"The user object."`
	Overview  string                `json:"overview" db:"overview"`
	Hire      string                `json:"hire" db:"hire"`
	Birthday  time.Time             `json:"birthday" db:"birthday"`
	Link      pgtype.Text           `json:"link" db:"link"`
	Email     pgtype.Text           `json:"email" db:"email"`
	Job       pgtype.Text           `json:"job" db:"job"`
	Vanity    pgtype.Text           `json:"vanity" db:"vanity"`
	Private   bool                  `json:"private" db:"private"`
	Developer bool                  `json:"developer" db:"developer"`
	Current   bool                  `json:"current" db:"current"`
	ExpToggle bool                  `json:"exptoggle" db:"exptoggle"`
	Nitro     bool                  `json:"nitro" db:"nitro"`
	Views     int                   `json:"views" db:"views"`
	Likes     int                   `json:"likes" db:"likes"`
	Date      time.Time             `json:"date" db:"date"`
}
