package main

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/app/handler/rest"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/app/middleware"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/app/repository"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/app/service"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/database/postgresql"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	db := postgresql.Connect()

	jwtAuth := jwt.NewJWT(os.Getenv("JWT_SECRET"), os.Getenv("JWT_TTL"))

	authRepo := repository.NewAuthRepository(db)

	authService := service.NewAuthService(authRepo, jwtAuth)

	_ = middleware.NewAuthMiddleware(jwtAuth) //TODO: Implement auth middleware in the router

	authHandler := rest.NewAuthHandler(authService)

	gin.SetMode(os.Getenv("GIN_MODE"))

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.POST("/auth/login", authHandler.Login)
	v1.POST("/auth/register", middleware.RequireSuperAdmin, authHandler.Register)

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalln(err)
	}
}
