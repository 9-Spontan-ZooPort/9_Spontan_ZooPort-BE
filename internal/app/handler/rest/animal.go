package rest

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/app/service"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/model"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

type IAnimalHandler interface {
	CreateAnimal(ctx *gin.Context)
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
