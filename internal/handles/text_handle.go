package handles

import (
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/gentleshare/lessapi-duckduckgo/internal/searchs"
	"github.com/gentleshare/lessapi-duckduckgo/internal/types"
	"github.com/gentleshare/lessapi-duckduckgo/internal/utils"
)

// SearchTextHandle search text
func SearchTextHandle(c *gin.Context) {
	// parse request params
	var payload types.SearchTextPayload
	if err := c.ShouldBindQuery(&payload); err != nil {
		c.JSON(200, utils.BuildApiError("invalid_params", "invalid request params"))
		return
	}
	// call search function
	var wg sync.WaitGroup
	wg.Add(1)

	// Channel to receive the search result
	resultCh := make(chan interface{}, 1)
	errCh := make(chan error, 1)

	go func() {
		defer wg.Done()
		resp, err := searchs.SearchText(payload)
		if err != nil {
			errCh <- err
			return
		}
		resultCh <- resp
	}()

	// Wait for the goroutine to finish
	wg.Wait()
	close(resultCh)
	close(errCh)

	// Check if there was an error
	if err := <-errCh; err != nil {
		c.JSON(200, utils.BuildApiError("search_error", err.Error()))
		return
	}

	// Get the result from the channel
	resp := <-resultCh

	// return response
	c.JSON(200, utils.BuildApiSuccessData(resp))
}
