package model

import (
	"github.com/google/uuid"
)

type CreateReportRequest struct {
	Description     string    `json:"description" binding:"required,max=65535"`
	PhotoUrl        string    `json:"photo_url" binding:"required,max=255,uri"`
	IsRequestDoctor bool      `json:"is_request_doctor" binding:"boolean"`
	AnimalID        uuid.UUID `json:"animal_id" binding:"required,uuid"`
}

type CreateReportResponse struct {
	ID string `json:"id"`
}

type GetReportResponse struct {
	ID              string `json:"id"`
	Description     string `json:"description"`
	PhotoUrl        string `json:"photo_url"`
	IsRequestDoctor bool   `json:"is_request_doctor"`
	IsApproved      bool   `json:"is_approved"`
	AnimalID        string `json:"animal_id"`
	UserID          string `json:"user_id"`
	CreatedAt       string `json:"created_at"`
}
