package handles

import (
	"github.com/gin-gonic/gin"

	"github.com/lessapidev/lessapi-duckduckgo/internal/searchs"
	"github.com/lessapidev/lessapi-duckduckgo/internal/types"
	"github.com/lessapidev/lessapi-duckduckgo/internal/utils"
)

// SearchTextHandle search text
func SearchTextHandle(c *gin.Context) {
	// parse request params
	var payload types.SearchTextPayload
	if err := c.ShouldBindQuery(&payload); err != nil {
		c.JSON(200, utils.BuildApiError("invalid_params", "invalid request params"))
		return
	}

	// search text
	resp, err := searchs.SearchText(payload)
	if err != nil {
		c.JSON(200, utils.BuildApiError("search_error", err.Error()))
		return
	}

	// return response
	c.JSON(200, utils.BuildApiSuccessData(resp))
}
