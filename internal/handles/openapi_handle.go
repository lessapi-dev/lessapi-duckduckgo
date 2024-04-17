package handles

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lessapidev/lessapi-duckduckgo/internal/utils"
)

// OpenAPIHandle handle OpenAPI request
func OpenAPIHandle(c *gin.Context) {
	// Read OpenAPI file from resources
	jsonRaw, err := utils.ReadLocalFile("./resource/openapi.json")
	if err != nil {
		c.JSON(200, utils.BuildApiError("read_openapi_error", err.Error()))
		return
	}
	// Parse JSON & Rewrite "servers" field to point to the current server
	jsonObj := map[string]interface{}{}
	if err := json.Unmarshal([]byte(jsonRaw), &jsonObj); err != nil {
		c.JSON(200, utils.BuildApiError("parse_openapi_error", err.Error()))
		return
	}
	jsonObj["servers"] = []map[string]string{}
	jsonObj["servers"] = append(jsonObj["servers"].([]map[string]string), map[string]string{
		"url":         "http://" + c.Request.Host,
		"description": "LessAPI DuckDuckGo API Server",
	})
	// Return OpenAPI JSON
	c.JSON(200, jsonObj)
}
