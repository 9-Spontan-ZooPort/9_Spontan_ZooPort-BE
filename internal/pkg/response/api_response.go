package response

import (
	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

func NewApiResponse(status int, message string, data any) ApiResponse {
	return ApiResponse{
		StatusCode: status,
		Message:    message,
		Data:       data,
	}
}

func (r ApiResponse) Send(ctx *gin.Context) {
	ctx.JSON(r.StatusCode, r)
}

func (r ApiResponse) isError() bool {
	return r.StatusCode >= 400
}

func InvalidRequestBody(err error) ApiResponse {
	return ApiResponse{
		StatusCode: 400,
		Message:    "invalid request body",
		Data:       err.Error(),
	}
}
