package handles

import (
	"github.com/gentleshare/lessapi-duckduckgo/internal/searchs"
	"github.com/gentleshare/lessapi-duckduckgo/internal/types"
	"github.com/gentleshare/lessapi-duckduckgo/internal/utils"
	"github.com/gin-gonic/gin"
)

// SearchTextHandle search text
func SearchTextHandle(c *gin.Context) {
	// parse request params
	var payload types.SearchTextPayload
	if err := c.ShouldBindQuery(&payload); err != nil {
		c.JSON(200, utils.BuildApiError("invalid_params", "invalid request params"))
		return
	}
	// call search function
	resp, err := searchs.SearchText(payload)
	if err != nil {
		c.JSON(200, utils.BuildApiError("search_error", err.Error()))
		return
	}
	// return response
	c.JSON(200, utils.BuildApiSuccessData(resp))
}
