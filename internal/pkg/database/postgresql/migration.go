package postgresql

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/entity"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		entity.User{},
	)
}
