package middleware

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/jwt"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type IAuthMiddleware interface {
	Authenticate(ctx *gin.Context)
	RequireRole(role string) gin.HandlerFunc
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

	if claims.ExpiresAt.Time.Before(time.Now()) {
		response.NewApiResponse(401, "token expired", nil).Send(ctx)
		ctx.Abort()
		return
	}

	ctx.Set("claims", claims)
	ctx.Next()
}

func (m AuthMiddleware) RequireRole(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claimsTemp, ok := ctx.Get("claims")
		if !ok {
			response.NewApiResponse(401, "unauthorized", nil).Send(ctx)
			ctx.Abort()
			return
		}
		claims := claimsTemp.(jwt.Claims)
		if claims.Role != role {
			response.NewApiResponse(403, "no permission", nil).Send(ctx)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
