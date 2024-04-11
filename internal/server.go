package internal

import (
	"github.com/gentleshare/lessapi-duckduckgo/internal/handles"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", handles.IndexHandle)
	r.GET("/search/text", handles.SearchTextHandle)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
