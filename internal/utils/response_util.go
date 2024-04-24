package utils

import (
	"fmt"

	"github.com/lessapidev/lessapi-duckduckgo/internal/types"

	"github.com/gin-gonic/gin"
)

// BuildApiSuccessData returns a new ApiResponse with code "success" and the given data.
func BuildApiSuccessData(data types.IResponse) *types.ApiResponse {
	return &types.ApiResponse{
		Code:    "success",
		Message: nil,
		Data:    data,
	}

}

// BuildApiError returns a new ApiResponse with the given error code and message.
func BuildApiError(errCode string, errorMessage string) *types.ApiResponse {
	return &types.ApiResponse{
		Code:    errCode,
		Message: &errorMessage,
	}
}

// -----------------------------------------------------------

// BuildLLMError returns an error message for LLM.
func BuildLLMError(errCode string, errorMessage string) string {
	return fmt.Sprintf("fail with error: %s\n error message is: %s", errCode, errorMessage)
}

// --------------------------------------------------------

func SendResponse(ctx *gin.Context, resp *types.ApiResponse) {
	// Check if LLM style is requested
	// Support both header and query parameter
	llmStyleByHeader := ctx.Request.Header.Get("x-llm-style")
	llmStyleByQuery := ctx.Query("llmStyle")

	if llmStyleByHeader == "1" || llmStyleByQuery == "1" {
		// LLM style
		if resp.Code == "success" {
			ctx.String(200, resp.Data.ToLLMStyle())
		} else {
			httpCode := 500
			if resp.Code == "invalid_params" {
				httpCode = 400
			}
			ctx.String(httpCode, BuildLLMError(resp.Code, *resp.Message))
		}
	} else {
		// API style
		if resp.Code == "success" {
			ctx.JSON(200, resp)
		} else {
			httpCode := 500
			if resp.Code == "invalid_params" {
				httpCode = 400
			}
			ctx.JSON(httpCode, resp)
		}
	}
}
