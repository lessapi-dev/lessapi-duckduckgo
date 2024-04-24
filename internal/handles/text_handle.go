package handles

import (
	"github.com/lessapidev/lessapi-duckduckgo/internal/searchs"
	"github.com/lessapidev/lessapi-duckduckgo/internal/types"
	"github.com/lessapidev/lessapi-duckduckgo/internal/utils"

	"github.com/gin-gonic/gin"
)

// SearchTextHandle search text
func SearchTextHandle(c *gin.Context) {
	// parse request params
	var payload types.SearchTextPayload
	if err := c.ShouldBindQuery(&payload); err != nil {
		utils.SendResponse(c, utils.BuildApiError("invalid_params", "invalid request params"))
		return
	}

	// search text
	resp, err := searchs.SearchText(payload)
	if err != nil {
		utils.SendResponse(c, utils.BuildApiError("search_error", err.Error()))
		return
	}

	// return response
	utils.SendResponse(c, utils.BuildApiSuccessData(resp))
}
