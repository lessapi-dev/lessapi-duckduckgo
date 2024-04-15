package searchs

// BrowserOptions Playwright browser options
type BrowserOptions struct {
	ProxyServer     *string
	EnableStealthJs *bool
	Locate          *string
	viewportWidth   *int
	viewportHeight  *int
}

type CloseBrowserHandler func() error
