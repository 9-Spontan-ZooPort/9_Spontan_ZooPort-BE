package repository

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/entity"
	"gorm.io/gorm"
)

type SpeciesRepository struct {
	db *gorm.DB
}

type ISpeciesRepository interface {
	CreateSpecies(species entity.Species) error
}

func NewSpeciesRepository(db *gorm.DB) ISpeciesRepository {
	return &SpeciesRepository{db: db}
}

func (r *SpeciesRepository) CreateSpecies(species entity.Species) error {
	return r.db.Create(&species).Error
}
