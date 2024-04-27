package rest

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/app/service"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/model"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IAnimalHandler interface {
	CreateAnimal(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetBySpecies(ctx *gin.Context)
}

type AnimalHandler struct {
	s service.IAnimalService
}

func NewAnimalHandler(s service.IAnimalService) IAnimalHandler {
	return &AnimalHandler{s: s}
}

func (h *AnimalHandler) CreateAnimal(ctx *gin.Context) {
	var request model.CreateAnimalRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.InvalidRequestBody(err).Send(ctx)
		return
	}

	h.s.CreateAnimal(request).Send(ctx)
}

func (h *AnimalHandler) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.NewApiResponse(400, "invalid id", nil).Send(ctx)
		return
	}

	h.s.GetByID(id).Send(ctx)
}

func (h *AnimalHandler) GetBySpecies(ctx *gin.Context) {
	speciesID := ctx.Query("species")

	h.s.GetBySpecies(speciesID).Send(ctx)
}
