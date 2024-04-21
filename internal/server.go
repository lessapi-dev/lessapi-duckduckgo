package internal

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/lessapidev/lessapi-duckduckgo/internal/handles"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/", handles.IndexHandle)
	r.GET("/openapi/openapi.json", handles.OpenAPIHandle)
	r.GET("/search/text", handles.SearchTextHandle)
	r.GET("/base/list-language", handles.ListLanguageHandle)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
