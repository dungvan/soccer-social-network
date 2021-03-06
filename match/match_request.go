package match

import (
	"time"
)

// IndexRequest struct
type IndexRequest struct {
	Page uint `form:"page" validate:"omitempty,min=1"`
}

// CreateRequest struct
type CreateRequest struct {
	TournamentID *uint     `json:"tourmanent_id" validate:"omitempty,required"`
	UserID       uint      `validate:"required"`
	Description  string    `json:"description" validate:"required"`
	StartDate    time.Time `json:"start_date" validate:"required"`
	Team1ID      uint      `json:"team1_id" validate:"required,nefield=Team2ID"`
	Team2ID      uint      `json:"team2_id" validate:"required,nefield=Team1ID"`
}

// UpdateGoaldsRequest struct
type UpdateGoaldsRequest struct {
	ID         uint  `validate:"required,min=1"`
	MasterID   uint  `validate:"required,min=1"`
	Team1Goals *uint `json:"team1_goals" validate:"omitempty,min=0"`
	Team2Goals *uint `json:"team2_goals" validate:"omitempty,min=0"`
}
