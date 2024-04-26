package middleware

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/entity"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/jwt"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"strings"
)

type IAuthMiddleware interface {
	Authenticate(ctx *gin.Context)
	RequireAdmin(ctx *gin.Context)
}

type AuthMiddleware struct {
	jwtAuth jwt.IJWT
}

func NewAuthMiddleware(jwtAuth jwt.IJWT) IAuthMiddleware {
	return AuthMiddleware{jwtAuth: jwtAuth}
}

func (m AuthMiddleware) Authenticate(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	if bearer == "" {
		response.NewApiResponse(401, "empty token", nil).Send(ctx)
		ctx.Abort()
		return
	}

	token := strings.Split(bearer, " ")[1]
	var claims jwt.Claims
	err := m.jwtAuth.Decode(token, &claims)
	if err != nil {
		response.NewApiResponse(401, "fail to validate token", err).Send(ctx)
		ctx.Abort()
		return
	}

	if claims.ExpiresAt.Time.Before(claims.IssuedAt.Time) {
		response.NewApiResponse(401, "token expired", nil).Send(ctx)
		ctx.Abort()
		return
	}

	ctx.Set("claims", claims)
	ctx.Next()
}

func (m AuthMiddleware) RequireAdmin(ctx *gin.Context) {
	userTemp, ok := ctx.Get("user")
	if !ok {
		response.NewApiResponse(401, "unauthorized", nil).Send(ctx)
		ctx.Abort()
		return
	}
	user := userTemp.(entity.User)
	if user.Role != "admin" {
		response.NewApiResponse(403, "no permission", nil).Send(ctx)
		ctx.Abort()
		return
	}

	ctx.Next()
}
