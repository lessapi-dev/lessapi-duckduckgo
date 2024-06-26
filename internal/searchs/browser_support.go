package searchs

import (
	"fmt"
	"strings"

	"github.com/lessapidev/lessapi-duckduckgo/internal/envs"
	"github.com/lessapidev/lessapi-duckduckgo/internal/utils"

	"github.com/playwright-community/playwright-go"
)

// NewBrowserContextWithOptions creates a new browser context with options
func NewBrowserContextWithOptions(opt BrowserOptions) (playwright.BrowserContext, CloseBrowserHandler, error) {

	// start playwright
	pw, err := playwright.Run()
	if err != nil {
		return nil, nil, fmt.Errorf("could not start playwright: %w", err)
	}

	// launch browser
	launchOptions := playwright.BrowserTypeLaunchOptions{}
	launchOptions.Headless = playwright.Bool(true)

	// set proxy
	if opt.ProxyServer != nil && *opt.ProxyServer != "" {
		launchOptions.Proxy = &playwright.Proxy{Server: *opt.ProxyServer}
	} else if utils.GetEnv(envs.LessapiDefaultViaProxyUrl) != "" {
		launchOptions.Proxy = &playwright.Proxy{Server: utils.GetEnv(envs.LessapiDefaultViaProxyUrl)}
	}
	browser, err := pw.Chromium.Launch(launchOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("could not launch browser: %w", err)
	}
	// set locale
	pLocate := opt.Language
	if pLocate == nil || *pLocate == "" {
		pLocate = playwright.String(utils.GetEnvOrDefault(envs.LessapiDefaultLanguage, "en-US"))
	} else {
		pLocate = playwright.String(strings.ReplaceAll(*pLocate, "_", "-"))
	}
	// set viewport
	pViewport := &playwright.Size{
		Width:  1920,
		Height: 1080,
	}
	if opt.viewportWidth != nil && opt.viewportHeight != nil {
		pViewport = &playwright.Size{
			Width:  *opt.viewportWidth,
			Height: *opt.viewportHeight,
		}
	}
	// create context
	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		Locale:   pLocate,
		Viewport: pViewport,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("could not create context: %w", err)
	}

	// add init stealth.min.js
	if opt.EnableStealthJs != nil && *opt.EnableStealthJs {
		stealthJs, err := utils.ReadLocalFile("./resource/stealth.min.js")
		if err != nil {
			return nil, nil, fmt.Errorf("could not read stealth.min.js: %w", err)
		}
		stealthJsScript := playwright.Script{
			Content: playwright.String(stealthJs),
		}
		if err = context.AddInitScript(stealthJsScript); err != nil {
			return nil, nil, fmt.Errorf("could not add stealth.min.js: %w", err)
		}
	}

	// create close handler
	closeHandler := func() error {
		if err := browser.Close(); err != nil {
			return fmt.Errorf("could not close browser: %w", err)
		}
		if err := pw.Stop(); err != nil {
			return fmt.Errorf("could not stop Playwright: %w", err)
		}
		return nil
	}

	return context, closeHandler, nil

}
