package middleware

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/jwt"
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/response"
	role2 "github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/role"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type IAuthMiddleware interface {
	Authenticate(ctx *gin.Context)
	RequireOneRole(role ...string) gin.HandlerFunc
	SoftAuthenticate(ctx *gin.Context)
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
	ctx.Set("role", claims.Role)
	ctx.Next()
}

func (m AuthMiddleware) SoftAuthenticate(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	if bearer == "" {
		ctx.Next()
		return
	}

	token := strings.Split(bearer, " ")[1]
	var claims jwt.Claims
	err := m.jwtAuth.Decode(token, &claims)
	if err != nil {
		ctx.Next()
		return
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		ctx.Next()
		return
	}

	ctx.Set("claims", claims)
	ctx.Set("role", claims.Role)
	ctx.Next()
}

func (m AuthMiddleware) RequireOneRole(role ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !role2.HasOneRole(ctx, role...) {
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
