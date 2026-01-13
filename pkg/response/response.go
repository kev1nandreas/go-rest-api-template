package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int         `json:"-"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
}

func NewSuccessResponse(message string, data interface{}) Response {
	return Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    message,
		Data:       data,
	}
}

func NewErrorResponse(statusCode interface{}, message string, err string) Response {
	return Response{
		StatusCode: statusCode.(int),
		Success:    false,
		Message:    message,
		Error:      err,
	}
}

func (r Response) Send(ctx *gin.Context) {
	ctx.JSON(r.StatusCode, r)
}
