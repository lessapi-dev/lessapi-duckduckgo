package handles

import (
	"github.com/gin-gonic/gin"

	"github.com/gentleshare/lessapi-duckduckgo"
	"github.com/gentleshare/lessapi-duckduckgo/internal/utils"
)

func IndexHandle(c *gin.Context) {
	resp := gin.H{
		"server":  lessapi_duckduckgo.ServerName,
		"github":  lessapi_duckduckgo.ProjectGithub,
		"version": lessapi_duckduckgo.Version,
	}
	c.JSON(200, utils.BuildApiSuccessData(resp))
}
