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
	speciesRepo := repository.NewSpeciesRepository(db)
	animalRepo := repository.NewAnimalRepository(db)

	authService := service.NewAuthService(authRepo, jwtAuth)
	speciesService := service.NewSpeciesService(speciesRepo)
	animalService := service.NewAnimalService(animalRepo)

	mid := middleware.NewAuthMiddleware(jwtAuth)

	authHandler := rest.NewAuthHandler(authService)
	speciesHandler := rest.NewSpeciesHandler(speciesService)
	animalHandler := rest.NewAnimalHandler(animalService)

	gin.SetMode(os.Getenv("GIN_MODE"))

	router := gin.Default()
	v1 := router.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/login", authHandler.Login)
	auth.POST("/register", middleware.RequireSuperAdmin, authHandler.Register)

	species := v1.Group("/species")
	species.POST("/", mid.Authenticate, mid.RequireRole("zookeeper"), speciesHandler.CreateSpecies)
	species.GET("/:id", speciesHandler.GetByID)
	species.GET("/", speciesHandler.GetAll)

	animals := v1.Group("/animals")
	animals.POST("/", mid.Authenticate, mid.RequireRole("zookeeper"), animalHandler.CreateAnimal)

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalln(err)
	}
}
