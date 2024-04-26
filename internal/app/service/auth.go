package service

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/app/repository"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/bcrypt"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/entity"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/jwt"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/model"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/response"
	"github.com/google/uuid"
)

type IAuthService interface {
	Register(request model.RegisterRequest) response.ApiResponse
	Login(request model.LoginRequest) response.ApiResponse
}

type AuthService struct {
	r   repository.IAuthRepository
	jwt jwt.IJWT
}

func NewAuthService(r repository.IAuthRepository, j jwt.IJWT) IAuthService {
	return &AuthService{r, j}
}

func (s *AuthService) Register(request model.RegisterRequest) response.ApiResponse {
	hashedPassword, err := bcrypt.Hash(request.Password)
	if err != nil {
		return response.NewApiResponse(500, "fail to hash password", nil)
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return response.NewApiResponse(500, "fail to generate id", nil)
	}

	user := entity.User{
		ID:       id,
		Email:    request.Email,
		Name:     request.Name,
		Password: hashedPassword,
		Role:     request.Role,
	}

	if err := s.r.Create(&user); err != nil {
		return response.NewApiResponse(500, "fail to create user", nil)
	}

	return response.NewApiResponse(201, "successfully created user", model.RegisterResponse{ID: id.String()})
}

func (s *AuthService) Login(request model.LoginRequest) response.ApiResponse {
	user, err := s.r.FindByEmail(request.Email)
	if err != nil {
		return response.NewApiResponse(404, "user not found", nil)
	}

	if err := bcrypt.ValidateHash(request.Password, user.Password); err != nil {
		return response.NewApiResponse(401, "invalid email or password", nil)
	}

	token, err := s.jwt.Create(user)
	if err != nil {
		return response.NewApiResponse(500, "fail to create token", nil)
	}

	return response.NewApiResponse(200, "successfully logged in", model.LoginResponse{
		ID:    user.ID.String(),
		Token: token,
	})
}
