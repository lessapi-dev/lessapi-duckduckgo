package utils

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"

	"github.com/lessapidev/lessapi-duckduckgo/internal/types"
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

func ReadLocalFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
