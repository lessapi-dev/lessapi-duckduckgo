package handles

import (
	"github.com/gin-gonic/gin"

	"github.com/lessapidev/lessapi-duckduckgo"
	"github.com/lessapidev/lessapi-duckduckgo/internal/utils"
)

func IndexHandle(c *gin.Context) {
	resp := IndexResponse{
		Server:  lessapi_duckduckgo.ServerName,
		Github:  lessapi_duckduckgo.ProjectGithub,
		Version: lessapi_duckduckgo.Version,
	}
	utils.SendResponse(c, utils.BuildApiSuccessData(resp))
}

type IndexResponse struct {
	Server  string `json:"server"`
	Github  string `json:"github"`
	Version string `json:"version"`
}

func (i IndexResponse) ToLLMStyle() string {
	return "Server: " + i.Server + "\nGithub: " + i.Github + "\nVersion: " + i.Version
}
