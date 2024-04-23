package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/lessapidev/lessapi-duckduckgo/internal/types"
)

// BuildApiSuccess returns a new ApiResponse with code "success" and empty data.
func BuildApiSuccess() *types.ApiResponse[any] {
	return &types.ApiResponse[any]{
		Code:    "success",
		Message: nil,
		Data:    gin.H{},
	}
}

// BuildApiSuccessData returns a new ApiResponse with code "success" and the given data.
func BuildApiSuccessData[T any](data T) *types.ApiResponse[T] {
	return &types.ApiResponse[T]{
		Code:    "success",
		Message: nil,
		Data:    data,
	}

}

// BuildApiError returns a new ApiResponse with the given error code and message.
func BuildApiError(errCode string, errorMessage string) *types.ApiResponse[any] {
	return &types.ApiResponse[any]{
		Code:    errCode,
		Message: &errorMessage,
	}
}
