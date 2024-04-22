package searchs

import (
	"fmt"

	"github.com/lessapidev/lessapi-duckduckgo/internal/types"

	"github.com/playwright-community/playwright-go"
)

func SearchText(param types.SearchTextPayload) (*types.SearchTextResponse, error) {

	// validate and set default values
	if param.Keyword == "" {
		return nil, fmt.Errorf("keyword is required")
	}
	if param.Region == "" {
		param.Region = "en-US"
	}
	if param.TimeLimit == "" {
		param.TimeLimit = ""
	}
	if param.MaxCount < 1 {
		param.MaxCount = 20
	}
	if param.ViaProxyUrl == "" {
		param.ViaProxyUrl = ""
	}

	// -----------------------------------------------------------
	// open browser
	// -----------------------------------------------------------

	// create new browser context
	browserOpt := BrowserOptions{
		ProxyServer:     &param.ViaProxyUrl,
		EnableStealthJs: playwright.Bool(true),
		Language:        &param.Region,
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

	// -----------------------------------------------------------
	// search text
	// -----------------------------------------------------------

	// open homepage, input keyword and search
	if _, err = page.Goto("https://duckduckgo.com/", playwright.PageGotoOptions{Timeout: playwright.Float(10 * 1000)}); err != nil {
		return nil, fmt.Errorf("could not goto: %w", err)
	}
	if err = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{Timeout: playwright.Float(10 * 1000)}); err != nil {
		return nil, fmt.Errorf("could not wait for load: %w", err)
	}
	inputLocator := page.Locator("input[name=q]")
	if err := inputLocator.WaitFor(playwright.LocatorWaitForOptions{Timeout: playwright.Float(10 * 1000)}); err != nil {
		return nil, fmt.Errorf("could not locate input: %w", err)
	}
	if err = inputLocator.Fill(param.Keyword); err != nil {
		return nil, fmt.Errorf("could not fill: %w", err)
	}

	// goto search result page
	if err = page.Keyboard().Press("Enter"); err != nil {
		return nil, fmt.Errorf("could not press Enter: %w", err)
	}
	majorAreaLocator := page.Locator("section[data-testid=mainline]")
	if err = majorAreaLocator.WaitFor(playwright.LocatorWaitForOptions{Timeout: playwright.Float(10 * 1000)}); err != nil {
		return nil, fmt.Errorf("could not find major div: %w", err)
	}
	if err = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
		State:   playwright.LoadStateNetworkidle,
		Timeout: playwright.Float(10 * 1000),
	}); err != nil {
		return nil, fmt.Errorf("could not wait for load result: %w", err)
	}

	// -----------------------------------------------------------
	// parse search results & auto next page
	// -----------------------------------------------------------

	// find all search result elements
	searchResultList := make([]types.SearchTextResultItem, 0)
	for len(searchResultList) < param.MaxCount {
		elResultLiList, err := page.Locator("ol[class=react-results--main]").Locator("article[data-testid=result]").All()
		if err != nil {
			return nil, fmt.Errorf("could not get search results: %w", err)
		}
		searchResultList = make([]types.SearchTextResultItem, 0)
		for i, elResultLi := range elResultLiList {
			title, err := elResultLi.Locator("a[data-testid=result-title-a]").InnerText()
			if err != nil {
				continue
			}
			href, err := elResultLi.Locator("a[data-testid=result-title-a]").GetAttribute("href")
			if err != nil {
				continue
			}
			description, err := elResultLi.Locator("div[data-result=snippet]").InnerText()
			if err != nil {
				continue
			}

			searchResultList = append(searchResultList, types.SearchTextResultItem{
				Order:       i + 1,
				Title:       title,
				Url:         href,
				Description: description,
			})
			if len(searchResultList) >= param.MaxCount {
				break // enough results
			}
		}
		if len(searchResultList) >= param.MaxCount {
			searchResultList = searchResultList[:param.MaxCount]
			break // enough results
		}
		// scroll to bottom
		if _, err = page.Evaluate("window.scrollTo(0, document.body.scrollHeight)"); err != nil {
			continue
		}
		// wait for elements to load
		if err = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
			State:   playwright.LoadStateNetworkidle,
			Timeout: playwright.Float(10 * 1000),
		}); err != nil {
			break
		}
		// auto next page
		locator := page.Locator("button[id=more-results]:not([disabled])")
		if err = locator.WaitFor(playwright.LocatorWaitForOptions{Timeout: playwright.Float(5 * 1000)}); err != nil {
			break
		}
		exist, err := locator.IsVisible()
		if err != nil {
			break
		}
		if exist {
			if err = locator.Click(); err != nil {
				break
			}
			// wait for elements to load
			if err = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
				State:   playwright.LoadStateNetworkidle,
				Timeout: playwright.Float(10 * 1000),
			}); err != nil {
				break
			}
		} else {
			break // no more results
		}
	}

	return &types.SearchTextResponse{
		Results: searchResultList,
	}, nil
}
