package repository

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/entity"
	"gorm.io/gorm"
)

type IAnimalRepository interface {
	CreateAnimal(animal entity.Animal) error
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
