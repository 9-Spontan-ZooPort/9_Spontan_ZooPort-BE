package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Animal struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:CHAR(36); PRIMARY_KEY; NOT NULL; UNIQUE"`
	SpeciesID       string    `gorm:"type:VARCHAR(255); NOT NULL"`
	Nickname        string    `gorm:"type:VARCHAR(255); NOT NULL; UNIQUE"`
	BirthDate       time.Time `gorm:"NOT NULL"`
	Gender          string    `gorm:"type:VARCHAR(255); NOT NULL"`
	Weight          float64   `gorm:"NOT NULL"`
	StatusKesehatan string    `gorm:"type:VARCHAR(255); NOT NULL; DEFAULT:'sehat'"`
	PhotoUrl        string    `gorm:"type:VARCHAR(255); NOT NULL"`
	MedicalHistorys []Report
}
