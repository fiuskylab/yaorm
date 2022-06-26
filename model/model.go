package model

import "time"

// Model is the base for yaORM
type Model struct {
	ID        int       `json:"id" yaorm:"pk,autoincrement"`
	CreatedAt time.Time `json:"created_at" yaorm:"trigger:create"`
	UpdatedAt time.Time `json:"updated_at" yaorm:"trigger:update"`
	// TODO:
	// 	Add DeletedAt field logic.
}
