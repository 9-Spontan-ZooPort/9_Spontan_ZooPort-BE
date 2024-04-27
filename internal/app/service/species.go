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
	GetByID(id string) response.ApiResponse
	GetAll() response.ApiResponse
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

func (s *SpeciesService) GetByID(id string) response.ApiResponse {
	species, err := s.r.GetByID(id)
	if err != nil {
		return response.NewApiResponse(500, "fail to get species", nil)
	}

	res := model.GetSpeciesResponse{
		ID:    species.ID,
		Name:  species.Name,
		Class: species.Class,
	}

	return response.NewApiResponse(200, "successfully get species", res)
}

func (s *SpeciesService) GetAll() response.ApiResponse {
	species, err := s.r.GetAll()
	if err != nil {
		return response.NewApiResponse(500, "fail to get species", nil)
	}

	var res []model.GetSpeciesResponse
	for _, v := range species {
		res = append(res, model.GetSpeciesResponse{
			ID:    v.ID,
			Name:  v.Name,
			Class: v.Class,
		})
	}

	return response.NewApiResponse(200, "successfully get species", res)
}
