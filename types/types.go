package types

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
