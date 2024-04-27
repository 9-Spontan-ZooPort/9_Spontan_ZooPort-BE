package repository

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IReportRepository interface {
	Create(report entity.Report) error
	GetByID(id uuid.UUID) (entity.Report, error)
	GetByAnimalID(animalID uuid.UUID) ([]entity.Report, error)
	GetByUserID(userID uuid.UUID) ([]entity.Report, error)
	Approve(id uuid.UUID) error
}

type ReportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) IReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) Create(report entity.Report) error {
	return r.db.Create(&report).Error
}

func (r *ReportRepository) GetByID(id uuid.UUID) (entity.Report, error) {
	report := entity.Report{
		ID: id,
	}
	err := r.db.First(&report).Error
	return report, err
}

func (r *ReportRepository) GetByAnimalID(animalID uuid.UUID) ([]entity.Report, error) {
	var reports []entity.Report
	err := r.db.Where("animal_id = ?", animalID).Find(&reports).Error
	return reports, err
}

func (r *ReportRepository) GetByUserID(userID uuid.UUID) ([]entity.Report, error) {
	var reports []entity.Report
	err := r.db.Where("user_id = ?", userID).Find(&reports).Error
	return reports, err
}

func (r *ReportRepository) Approve(id uuid.UUID) error {
	return r.db.Model(&entity.Report{}).Where("id = ?", id).Update("is_approved", true).Error
}
