package utils

import (
	"github.com/gin-gonic/gin"

	"github.com/gentleshare/lessapi-duckduckgo/internal/types"
)

func BuildApiSuccess() *types.ApiResponse[any] {
	return &types.ApiResponse[any]{
		Code:    "success",
		Message: nil,
		Data:    gin.H{},
	}
}

func BuildApiSuccessData[T any](data T) *types.ApiResponse[T] {
	return &types.ApiResponse[T]{
		Code:    "success",
		Message: nil,
		Data:    data,
	}

}

func BuildApiError(errCode string, errorMessage string) *types.ApiResponse[any] {
	return &types.ApiResponse[any]{
		Code:    errCode,
		Message: &errorMessage,
	}
}
