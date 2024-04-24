package types

type IResponse interface {
	ToLLMStyle() string
}

type ApiResponse struct {
	Code    string    `json:"code"`
	Message *string   `json:"message,omitempty"`
	Data    IResponse `json:"data,omitempty"`
}

type EmptyResponse struct{}

func (e EmptyResponse) ToLLMStyle() string {
	return "success"
}
