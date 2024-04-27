package service

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/app/repository"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/entity"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/model"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/response"
)

type SpeciesService struct {
	r repository.ISpeciesRepository
}

type ISpeciesService interface {
	CreateSpecies(request model.CreateSpeciesRequest) response.ApiResponse
}

func NewSpeciesService(r repository.ISpeciesRepository) ISpeciesService {
	return &SpeciesService{r: r}
}

func (s *SpeciesService) CreateSpecies(request model.CreateSpeciesRequest) response.ApiResponse {
	species := entity.Species{
		ID:    request.ID,
		Name:  request.Name,
		Class: request.Class,
	}

	if err := s.r.CreateSpecies(species); err != nil {
		return response.NewApiResponse(500, "fail to create species", nil)
	}

	return response.NewApiResponse(201, "successfully created species", nil)
}
