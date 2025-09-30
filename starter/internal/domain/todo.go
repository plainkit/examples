// Package domain contains core business entities.
// These are pure data structures with no business logic.
package domain

import "time"

// Todo represents a todo item in the application.
type Todo struct {
	ID        string
	Title     string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
