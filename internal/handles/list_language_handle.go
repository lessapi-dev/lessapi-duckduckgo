package handles

import (
	"sync"
	"time"

	"github.com/lessapidev/lessapi-duckduckgo/internal/searchs"
	"github.com/lessapidev/lessapi-duckduckgo/internal/types"
	"github.com/lessapidev/lessapi-duckduckgo/internal/utils"

	"github.com/gin-gonic/gin"
)

type languagesCache struct {
	languages []types.LanguageType
	ttl       time.Duration
	lock      sync.RWMutex
}

// ListLanguageHandle list all language types
func ListLanguageHandle(c *gin.Context) {
	// parse request params
	var payload types.ListLanguagePayload
	if err := c.ShouldBindQuery(&payload); err != nil {
		c.JSON(200, utils.BuildApiError("invalid_params", "invalid request params"))
		return
	}

	// search text
	resp, err := searchs.ListLanguage(payload)
	if err != nil {
		c.JSON(200, utils.BuildApiError("search_error", err.Error()))
		return
	}

	// return response
	c.JSON(200, utils.BuildApiSuccessData(resp))
}
