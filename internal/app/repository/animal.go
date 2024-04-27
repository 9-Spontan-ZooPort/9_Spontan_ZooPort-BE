package repository

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IAnimalRepository interface {
	CreateAnimal(animal entity.Animal) error
	GetByID(id uuid.UUID) (entity.Animal, error)
	GetBySpecies(speciesID string) ([]entity.Animal, error)
}

type AnimalRepository struct {
	db *gorm.DB
}

func NewAnimalRepository(db *gorm.DB) IAnimalRepository {
	return &AnimalRepository{db: db}
}

func (r *AnimalRepository) CreateAnimal(animal entity.Animal) error {
	return r.db.Create(&animal).Error
}

func (r *AnimalRepository) GetByID(id uuid.UUID) (entity.Animal, error) {
	animal := entity.Animal{
		ID: id,
	}
	err := r.db.First(&animal).Error
	return animal, err
}

func (r *AnimalRepository) GetBySpecies(speciesID string) ([]entity.Animal, error) {
	var animals []entity.Animal
	err := r.db.Where("species_id = ?", speciesID).Find(&animals).Error
	return animals, err
}
