package searchs

import (
	"fmt"
	"io"
	"os"

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
	if opt.ProxyServer != nil && *opt.ProxyServer != "" {
		launchOptions.Proxy = &playwright.Proxy{Server: *opt.ProxyServer}
	}
	browser, err := pw.Chromium.Launch(launchOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("could not launch browser: %w", err)
	}

	// create new context
	pLocate := opt.Locate
	if pLocate == nil {
		pLocate = playwright.String("en-Us")
	}
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
	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		Locale:   pLocate,
		Viewport: pViewport,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("could not create context: %w", err)
	}

	// add init stealth.min.js
	if opt.EnableStealthJs != nil && *opt.EnableStealthJs {
		stealthJs, err := readLocalFile("./resource/stealth.min.js")
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

// TODO move to utils
func readLocalFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
