package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:CHAR(36); PRIMARY_KEY; NOT NULL; UNIQUE"`
	Email    string    `gorm:"type:VARCHAR(320); NOT NULL; UNIQUE"`
	Name     string    `gorm:"type:VARCHAR(255); NOT NULL"`
	Password string    `gorm:"type:VARCHAR(255); NOT NULL"`
	Role     string    `gorm:"type:VARCHAR(255); NOT NULL"`
	Reports  []Report
}
