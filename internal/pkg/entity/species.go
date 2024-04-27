package entity

import (
	"gorm.io/gorm"
)

type Species struct {
	gorm.Model
	ID      string `gorm:"type:VARCHAR(255); PRIMARY_KEY; NOT NULL; UNIQUE"`
	Name    string `gorm:"type:VARCHAR(255); NOT NULL"`
	Class   string `gorm:"type:VARCHAR(255); NOT NULL"`
	Animals []Animal
}
