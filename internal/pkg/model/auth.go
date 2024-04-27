package model

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email,max=320"`
	Name     string `json:"name" binding:"required,min=1,max=255"`
	Password string `json:"password" binding:"required,min=8,max=255"`
	Role     string `json:"role" binding:"required,oneof=admin zookeeper"`
}

type RegisterResponse struct {
	ID string `json:"id"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
