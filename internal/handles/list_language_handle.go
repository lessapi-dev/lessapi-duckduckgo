package handles

import (
	"time"

	"github.com/lessapidev/lessapi-duckduckgo/internal/searchs"
	"github.com/lessapidev/lessapi-duckduckgo/internal/types"
	"github.com/lessapidev/lessapi-duckduckgo/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

var languagesCache = cache.New(5*time.Minute, 10*time.Minute)

// ListLanguageHandle list all language types
func ListLanguageHandle(c *gin.Context) {
	// parse request params
	var payload types.ListLanguagePayload
	if err := c.ShouldBindQuery(&payload); err != nil {
		c.JSON(200, utils.BuildApiError("invalid_params", "invalid request params"))
		return
	}

	// get from cache
	languages, found := languagesCache.Get("languages")
	if found {
		c.JSON(200, utils.BuildApiSuccessData(types.ListLanguageResponse{
			Languages: languages.([]types.LanguageType),
		}))
		return
	}

	// search text
	resp, err := searchs.ListLanguage(payload)
	if err != nil {
		c.JSON(200, utils.BuildApiError("search_error", err.Error()))
		return
	}

	if len(resp.Languages) > 0 {
		languagesCache.Set("languages", resp.Languages, cache.DefaultExpiration)
	}

	// return response
	c.JSON(200, utils.BuildApiSuccessData(resp))
}
