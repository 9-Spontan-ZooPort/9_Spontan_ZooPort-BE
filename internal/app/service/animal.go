package service

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/app/repository"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/entity"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/model"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/response"
	"github.com/google/uuid"
	"time"
)

type IAnimalService interface {
	CreateAnimal(request model.CreateAnimalRequest) response.ApiResponse
	GetByID(id uuid.UUID, hideDetails bool) response.ApiResponse
	GetBySpecies(speciesID string, hideDetails bool) response.ApiResponse
}

type AnimalService struct {
	r repository.IAnimalRepository
}

func NewAnimalService(r repository.IAnimalRepository) IAnimalService {
	return &AnimalService{r: r}
}

func (s *AnimalService) CreateAnimal(request model.CreateAnimalRequest) response.ApiResponse {
	id, err := uuid.NewRandom()
	if err != nil {
		return response.NewApiResponse(500, "fail to generate id", nil)
	}

	birthDate, err := time.Parse(time.RFC3339, request.BirthDate)
	if err != nil {
		return response.NewApiResponse(400, "invalid birth date", nil)
	}

	animal := entity.Animal{
		ID:              id,
		SpeciesID:       request.SpeciesID,
		Nickname:        request.Nickname,
		BirthDate:       birthDate,
		Gender:          request.Gender,
		Weight:          request.Weight,
		StatusKesehatan: request.StatusKesehatan,
		PhotoUrl:        request.PhotoUrl,
		Description:     request.Description,
	}

	if err = s.r.CreateAnimal(animal); err != nil {
		return response.NewApiResponse(500, "fail to create animal", nil)
	}

	return response.NewApiResponse(201, "successfully created animal", model.CreateAnimalResponse{ID: id.String()})
}

func (s *AnimalService) GetByID(id uuid.UUID, hideDetails bool) response.ApiResponse {
	animal, err := s.r.GetByID(id)
	if err != nil {
		return response.NewApiResponse(500, "fail to get animal", nil)
	}

	res := model.GetAnimalResponse{
		ID:              animal.ID,
		SpeciesID:       animal.SpeciesID,
		Nickname:        animal.Nickname,
		BirthDate:       animal.BirthDate,
		Gender:          animal.Gender,
		Weight:          animal.Weight,
		StatusKesehatan: animal.StatusKesehatan,
		PhotoUrl:        animal.PhotoUrl,
		Description:     animal.Description,
	}

	if hideDetails {
		res.StatusKesehatan = ""
	}

	return response.NewApiResponse(200, "successfully get animal", res)
}

func (s *AnimalService) GetBySpecies(speciesID string, hideDetails bool) response.ApiResponse {
	animals, err := s.r.GetBySpecies(speciesID)
	if err != nil {
		return response.NewApiResponse(500, "fail to get animals", nil)
	}

	var res []model.GetAnimalResponse
	for _, v := range animals {
		nextRes := model.GetAnimalResponse{
			ID:              v.ID,
			SpeciesID:       v.SpeciesID,
			Nickname:        v.Nickname,
			BirthDate:       v.BirthDate,
			Gender:          v.Gender,
			Weight:          v.Weight,
			StatusKesehatan: v.StatusKesehatan,
			PhotoUrl:        v.PhotoUrl,
			Description:     v.Description,
		}
		if hideDetails {
			nextRes.StatusKesehatan = ""
		}
		res = append(res, nextRes)
	}

	return response.NewApiResponse(200, "successfully get animals", res)
}
