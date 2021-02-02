package model

import (
	"time"

	"github.com/google/uuid"
)

// AccessibilityRequest models a 508 request
type AccessibilityRequest struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}