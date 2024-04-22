package searchs

import (
	"fmt"
	"strings"

	"github.com/lessapidev/lessapi-duckduckgo/internal/types"
	"github.com/playwright-community/playwright-go"
)

func ListLanguage(param types.ListLanguagePayload) (*types.ListLanguageResponse, error) {

	// -----------------------------------------------------------
	// open browser
	// -----------------------------------------------------------

	// create new browser context
	browserOpt := BrowserOptions{
		ProxyServer:     &param.ViaProxyUrl,
		EnableStealthJs: playwright.Bool(true),
		viewportWidth:   playwright.Int(1920),
		viewportHeight:  playwright.Int(1080 - 35),
	}
	browserCtx, browserClose, err := NewBrowserContextWithOptions(browserOpt)
	if err != nil {
		return nil, fmt.Errorf("could not create virtual browser context: %w", err)
	}
	defer (func() {
		_ = browserClose()
	})()

	// create new page
	page, err := browserCtx.NewPage()
	if err != nil {
		return nil, fmt.Errorf("could not create page: %w", err)
	}

	// open search page
	if _, err = page.Goto("https://duckduckgo.com/?q=duckduckgo", playwright.PageGotoOptions{Timeout: playwright.Float(10 * 1000)}); err != nil {
		return nil, fmt.Errorf("could not goto: %w", err)
	}
	if err = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{Timeout: playwright.Float(10 * 1000)}); err != nil {
		return nil, fmt.Errorf("could not wait for load: %w", err)
	}

	// click settings
	settingLocator := page.Locator("div[class='dropdown dropdown--settings'] a")
	if err := settingLocator.WaitFor(playwright.LocatorWaitForOptions{Timeout: playwright.Float(5 * 1000)}); err != nil {
		return nil, fmt.Errorf("could not locate setting: %w", err)
	}
	if err := settingLocator.Click(); err != nil {
		return nil, fmt.Errorf("could not click: %w", err)
	}

	// find language select
	selectLocator := page.Locator("select[id=setting_kad]")
	if err := selectLocator.WaitFor(playwright.LocatorWaitForOptions{Timeout: playwright.Float(5 * 1000)}); err != nil {
		return nil, fmt.Errorf("could not locate select: %w", err)
	}

	// get languages
	optionLocator := selectLocator.Locator("option")
	options, err := optionLocator.All()
	if err != nil {
		return nil, fmt.Errorf("could not get options: %w", err)
	}
	l := make([]types.LanguageType, 0)
	for _, option := range options {
		code, err := option.GetAttribute("value")
		if err != nil {
			return nil, fmt.Errorf("could not get value: %w", err)
		}
		name, err := option.TextContent()
		if err != nil {
			return nil, fmt.Errorf("could not get text content: %w", err)
		}
		l = append(l, types.LanguageType{
			Code: strings.ReplaceAll(code, "_", "-"),
			Name: name,
		})
	}

	return &types.ListLanguageResponse{
		Languages: l,
	}, nil
}
