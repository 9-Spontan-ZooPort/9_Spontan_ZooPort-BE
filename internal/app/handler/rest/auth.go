package rest

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/app/service"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/model"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	s service.IAuthService
}

type IAuthHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

func NewAuthHandler(s service.IAuthService) IAuthHandler {
	return &AuthHandler{s}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var request model.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.InvalidRequestBody(err).Send(ctx)
		return
	}

	h.s.Login(request).Send(ctx)
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var request model.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.InvalidRequestBody(err).Send(ctx)
		return
	}

	h.s.Register(request).Send(ctx)
}
