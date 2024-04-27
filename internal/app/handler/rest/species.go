package rest

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/app/service"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/model"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

type SpeciesHandler struct {
	s service.ISpeciesService
}

type ISpeciesHandler interface {
	CreateSpecies(ctx *gin.Context)
}

func NewSpeciesHandler(s service.ISpeciesService) ISpeciesHandler {
	return &SpeciesHandler{s: s}
}

func (h *SpeciesHandler) CreateSpecies(ctx *gin.Context) {
	var request model.CreateSpeciesRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.InvalidRequestBody(err).Send(ctx)
		return
	}

	h.s.CreateSpecies(request).Send(ctx)
}
