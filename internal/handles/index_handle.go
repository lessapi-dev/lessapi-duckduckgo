package handles

import (
	"github.com/gin-gonic/gin"

	"github.com/lessapidev/lessapi-duckduckgo"
	"github.com/lessapidev/lessapi-duckduckgo/internal/utils"
)

func IndexHandle(c *gin.Context) {
	resp := gin.H{
		"server":  lessapi_duckduckgo.ServerName,
		"github":  lessapi_duckduckgo.ProjectGithub,
		"version": lessapi_duckduckgo.Version,
	}
	c.JSON(200, utils.BuildApiSuccessData(resp))
}
