package model

import (
	"github.com/google/uuid"
	"time"
)

type CreateAnimalRequest struct {
	SpeciesID       string  `json:"species_id" binding:"required,max=255"`
	Nickname        string  `json:"nickname" binding:"required,max=255"`
	BirthDate       string  `json:"birth_date" binding:"required"`
	Gender          string  `json:"gender" binding:"required,oneof=jantan betina"`
	Weight          float64 `json:"weight" binding:"required,min=0"`
	PhotoUrl        string  `json:"photo_url" binding:"required,max=255,url"`
	StatusKesehatan string  `json:"status_kesehatan" binding:"required,oneof=sehat pemulihan sakit kritis"`
	Description     string  `json:"description" binding:"required,max=65535"`
}

type CreateAnimalResponse struct {
	ID string `json:"id"`
}

type GetAnimalResponse struct {
	ID              uuid.UUID `json:"id"`
	SpeciesID       string    `json:"species_id"`
	Nickname        string    `json:"nickname"`
	BirthDate       time.Time `json:"birth_date"`
	Gender          string    `json:"gender"`
	Weight          float64   `json:"weight"`
	StatusKesehatan string    `json:"status_kesehatan"`
	PhotoUrl        string    `json:"photo_url"`
	Description     string    `json:"description"`
}
