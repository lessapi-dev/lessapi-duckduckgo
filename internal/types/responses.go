package types

import "fmt"

type SearchTextPayload struct {
	// Keywords to search (required)
	Keyword string `json:"keyword" form:"keyword" validate:"required"`
	// Region to search (optional) (default "en-US")
	Region string `json:"region" form:"region"`
	// Time limit to search (optional) ( "" default all , "d" past day, "w" past week, "m" past month, "y" past year)
	TimeLimit string `json:"timeLimit" form:"timeLimit"`
	// Max count of search results (optional) (default 20)
	MaxCount int `json:"maxCount" form:"maxCount"`
	// Proxy url to use (optional)
	ViaProxyUrl string `json:"viaProxyUrl" form:"viaProxyUrl"`
}

type SearchTextResponse struct {
	Results []SearchTextResultItem `json:"results"`
}

func (s SearchTextResponse) ToLLMStyle() string {
	var resultStr string
	for i, item := range s.Results {
		resultStr += fmt.Sprintf("# %d. %s\nURL:%s\nAbstract:%s\n\n", i+1, item.Title, item.Url, item.Description)
	}
	return resultStr
}

type SearchTextResultItem struct {
	Order       int    `json:"order"`       // order of the result
	Title       string `json:"title"`       // title of the result
	Url         string `json:"url"`         // url of the result
	Description string `json:"description"` // description of the result
}

// -----------------------------------------------------------

type LanguageType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ListLanguagePayload struct {
	// Proxy url to use (optional)
	ViaProxyUrl string `json:"viaProxyUrl" form:"viaProxyUrl"`
}

type ListLanguageResponse struct {
	Languages []LanguageType `json:"languages"`
}

func (l ListLanguageResponse) ToLLMStyle() string {
	var resultStr string
	for _, item := range l.Languages {
		resultStr += fmt.Sprintf("%s (%s)\n", item.Name, item.Code)
	}
	return resultStr
}
