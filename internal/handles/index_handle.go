package handles

import (
	"github.com/gentleshare/lessapi-duckduckgo/internal/utils"
	"github.com/gin-gonic/gin"
)

func IndexHandle(c *gin.Context) {
	c.JSON(200, utils.BuildApiSuccess())
}
