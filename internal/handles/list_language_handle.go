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
		utils.SendResponse(c, utils.BuildApiError("invalid_params", "invalid request params"))
		return
	}

	// get from cache
	languages, found := languagesCache.Get("languages")
	if found {
		resp := types.ListLanguageResponse{
			Languages: languages.([]types.LanguageType),
		}
		utils.SendResponse(c, utils.BuildApiSuccessData(resp))
		return
	}

	// search text
	resp, err := searchs.ListLanguage(payload)
	if err != nil {
		utils.SendResponse(c, utils.BuildApiError("fetch_error", err.Error()))
		return
	}

	if len(resp.Languages) > 0 {
		languagesCache.Set("languages", resp.Languages, cache.DefaultExpiration)
	}

	// return response
	utils.SendResponse(c, utils.BuildApiSuccessData(resp))
}
