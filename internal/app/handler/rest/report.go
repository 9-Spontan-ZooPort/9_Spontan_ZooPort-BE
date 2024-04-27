package rest

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/app/service"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/jwt"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/model"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IReportHandler interface {
	Create(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Get(ctx *gin.Context)
	Approve(ctx *gin.Context)
}

type ReportHandler struct {
	s service.IReportService
}

func NewReportHandler(s service.IReportService) IReportHandler {
	return &ReportHandler{s: s}
}

func (h *ReportHandler) Create(ctx *gin.Context) {
	var request model.CreateReportRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.InvalidRequestBody(err).Send(ctx)
		return
	}

	claimsTemp, _ := ctx.Get("claims")
	claims := claimsTemp.(jwt.Claims)
	id, _ := uuid.Parse(claims.Subject)

	h.s.Create(request, id).Send(ctx)
}

func (h *ReportHandler) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.NewApiResponse(400, "invalid id", nil).Send(ctx)
		return
	}

	h.s.GetByID(id).Send(ctx)
}

func (h *ReportHandler) Get(ctx *gin.Context) {
	animalId := ctx.Query("animal_id")
	if animalId != "" {
		id, err := uuid.Parse(animalId)
		if err != nil {
			response.NewApiResponse(400, "invalid animal id", nil).Send(ctx)
			return
		}
		h.s.GetByAnimalID(id).Send(ctx)
		return
	}

	userId := ctx.Query("user_id")
	if userId != "" {
		id, err := uuid.Parse(userId)
		if err != nil {
			response.NewApiResponse(400, "invalid user id", nil).Send(ctx)
			return
		}
		h.s.GetByUserID(id).Send(ctx)
		return

	}
}

func (h *ReportHandler) Approve(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.NewApiResponse(400, "invalid id", nil).Send(ctx)
		return
	}

	h.s.Approve(id).Send(ctx)
}
