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
	GetByID(id string) (entity.Species, error)
	GetAll() ([]entity.Species, error)
}

func NewSpeciesRepository(db *gorm.DB) ISpeciesRepository {
	return &SpeciesRepository{db: db}
}

func (r *SpeciesRepository) CreateSpecies(species entity.Species) error {
	return r.db.Create(&species).Error
}

func (r *SpeciesRepository) GetByID(id string) (entity.Species, error) {
	var species entity.Species
	err := r.db.Where("id = ?", id).First(&species).Error
	return species, err
}

func (r *SpeciesRepository) GetAll() ([]entity.Species, error) {
	var species []entity.Species
	err := r.db.Find(&species).Error
	return species, err
}
