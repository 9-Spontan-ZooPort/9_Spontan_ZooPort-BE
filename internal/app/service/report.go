package service

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/app/repository"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/entity"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/model"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/response"
	"github.com/google/uuid"
	"time"
)

type IReportService interface {
	Create(request model.CreateReportRequest, userId uuid.UUID) response.ApiResponse
	GetByID(id uuid.UUID) response.ApiResponse
	GetByAnimalID(animalID uuid.UUID) response.ApiResponse
	GetByUserID(userID uuid.UUID) response.ApiResponse
	Approve(id uuid.UUID) response.ApiResponse
}

type ReportService struct {
	r repository.IReportRepository
}

func NewReportService(r repository.IReportRepository) IReportService {
	return &ReportService{r: r}
}

func (s *ReportService) Create(request model.CreateReportRequest, userId uuid.UUID) response.ApiResponse {
	id, err := uuid.NewRandom()
	if err != nil {
		return response.NewApiResponse(500, "fail to generate id", nil)
	}

	report := entity.Report{
		ID:              id,
		Description:     request.Description,
		PhotoUrl:        request.PhotoUrl,
		IsRequestDoctor: request.IsRequestDoctor,
		IsApproved:      false,
		AnimalID:        request.AnimalID,
		UserID:          userId,
	}

	if err = s.r.Create(report); err != nil {
		return response.NewApiResponse(500, "fail to create report", nil)
	}

	return response.NewApiResponse(201, "successfully created report", model.CreateReportResponse{ID: id.String()})
}

func (s *ReportService) GetByID(id uuid.UUID) response.ApiResponse {
	report, err := s.r.GetByID(id)
	if err != nil {
		return response.NewApiResponse(500, "fail to get report", nil)
	}

	res := model.GetReportResponse{
		ID:              report.ID.String(),
		Description:     report.Description,
		PhotoUrl:        report.PhotoUrl,
		IsRequestDoctor: report.IsRequestDoctor,
		IsApproved:      report.IsApproved,
		AnimalID:        report.AnimalID.String(),
		UserID:          report.UserID.String(),
		CreatedAt:       report.CreatedAt.Format(time.RFC3339),
	}

	return response.NewApiResponse(200, "successfully get report", res)
}

func (s *ReportService) GetByAnimalID(animalID uuid.UUID) response.ApiResponse {
	reports, err := s.r.GetByAnimalID(animalID)
	if err != nil {
		return response.NewApiResponse(500, "fail to get reports", nil)
	}

	res := make([]model.GetReportResponse, 0)
	for _, report := range reports {
		res = append(res, model.GetReportResponse{
			ID:              report.ID.String(),
			Description:     report.Description,
			PhotoUrl:        report.PhotoUrl,
			IsRequestDoctor: report.IsRequestDoctor,
			IsApproved:      report.IsApproved,
			AnimalID:        report.AnimalID.String(),
			UserID:          report.UserID.String(),
			CreatedAt:       report.CreatedAt.Format(time.RFC3339),
		})
	}

	return response.NewApiResponse(200, "successfully get reports", res)
}

func (s *ReportService) GetByUserID(userID uuid.UUID) response.ApiResponse {
	reports, err := s.r.GetByUserID(userID)
	if err != nil {
		return response.NewApiResponse(500, "fail to get reports", nil)
	}

	res := make([]model.GetReportResponse, 0)
	for _, report := range reports {
		res = append(res, model.GetReportResponse{
			ID:              report.ID.String(),
			Description:     report.Description,
			PhotoUrl:        report.PhotoUrl,
			IsRequestDoctor: report.IsRequestDoctor,
			IsApproved:      report.IsApproved,
			AnimalID:        report.AnimalID.String(),
			UserID:          report.UserID.String(),
			CreatedAt:       report.CreatedAt.Format(time.RFC3339),
		})
	}

	return response.NewApiResponse(200, "successfully get reports", res)
}

func (s *ReportService) Approve(id uuid.UUID) response.ApiResponse {
	if err := s.r.Approve(id); err != nil {
		return response.NewApiResponse(500, "fail to approve report", nil)
	}

	// TODO: Implement doctor email notification

	return response.NewApiResponse(200, "successfully approve report", nil)
}
