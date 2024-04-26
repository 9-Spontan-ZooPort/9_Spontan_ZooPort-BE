package middleware

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"os"
)

func RequireSuperAdmin(ctx *gin.Context) {
	token := ctx.GetHeader("Superadmin")
	if token != os.Getenv("SUPERADMIN_API_KEY") {
		response.ApiResponse{
			StatusCode: 401,
			Message:    "no permission",
			Data:       nil,
		}.Send(ctx)
		ctx.Abort()
		return
	}
	ctx.Next()
}
