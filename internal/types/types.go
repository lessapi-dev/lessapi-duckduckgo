package types

type ApiResponse[T any] struct {
	Code    string  `json:"code"`
	Message *string `json:"message,omitempty"`
	Data    T       `json:"data,omitempty"`
}

type SearchTextPayload struct {
	// Keywords to search (required)
	Keyword string `json:"keyword" form:"keyword" validate:"required"`
	// Region to search (optional) (default "wt-wt")
	Region string `json:"region" form:"region"`
	// Time limit to search (optional) ( "" default all , "d" past day, "w" past week, "m" past month, "y" past year)
	TimeLimit string `json:"timeLimit" form:"timeLimit"`
	// Max count of search results (optional) (default 10)
	MaxCount int `json:"maxCount" form:"maxCount"`
	// Proxy url to use (optional)
	ViaProxyUrl string `json:"viaProxyUrl" form:"viaProxyUrl"`
}

type SearchTextResponse struct {
	Results []SearchTextResultItem `json:"results"`
}

type SearchTextResultItem struct {
	Title       string `json:"title"`
	Href        string `json:"href"`
	Description string `json:"description"`
}
