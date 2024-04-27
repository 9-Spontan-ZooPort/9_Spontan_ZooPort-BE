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
