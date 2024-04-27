package role

import (
	"github.com/gin-gonic/gin"
)

func HasOneRole(ctx *gin.Context, roles ...string) bool {
	roleReal, ok := ctx.Get("role")
	if !ok {
		return false
	}

	for _, role := range roles {
		if roleReal == role {
			return true
		}
	}

	return false
}
