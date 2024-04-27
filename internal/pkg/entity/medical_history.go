package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicalRecord struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:CHAR(36); PRIMARY_KEY; NOT NULL; UNIQUE"`
	Description  string    `gorm:"type:TEXT; NOT NULL"`
	PhotoUrl     string    `gorm:"type:VARCHAR(255); NOT NULL"`
	IsSensitive  bool      `gorm:"NOT NULL"`
	IsCallDoctor bool      `gorm:"NOT NULL"`
	IsApproved   bool      `gorm:"NOT NULL"`
	AnimalID     uuid.UUID `gorm:"type:CHAR(36); NOT NULL"`
}
