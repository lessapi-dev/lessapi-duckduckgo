package searchs

import (
	"fmt"
	"time"

	"github.com/playwright-community/playwright-go"

	"github.com/lessapidev/lessapi-duckduckgo/internal/types"
)

var pageWaitOpt = playwright.PageWaitForLoadStateOptions{
	State: playwright.LoadStateDomcontentloaded,
}

func SearchText(param types.SearchTextPayload) (*types.SearchTextResponse, error) {

	// validate and set default values
	if param.Keyword == "" {
		return nil, fmt.Errorf("keyword is required")
	}
	if param.Region == "" {
		param.Region = "wt-wt"
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
		Locate:          &param.Region,
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
	if _, err = page.Goto("https://duckduckgo.com/"); err != nil {
		return nil, fmt.Errorf("could not goto: %w", err)
	}
	if err = page.WaitForLoadState(pageWaitOpt); err != nil {
		return nil, fmt.Errorf("could not wait for load: %w", err)
	}
	time.Sleep(1 * time.Second) // wait for page to load
	if err = page.Fill("input[name=q]", param.Keyword); err != nil {
		return nil, fmt.Errorf("could not fill: %w", err)
	}
	if err = page.Keyboard().Press("Enter"); err != nil {
		return nil, fmt.Errorf("could not press Enter: %w", err)
	}
	// wait for search result page to load
	if err = page.WaitForLoadState(pageWaitOpt); err != nil {
		return nil, fmt.Errorf("could not wait for load: %w", err)
	}
	time.Sleep(5 * time.Second) // wait for page to load

	// -----------------------------------------------------------
	// fetch pages
	// -----------------------------------------------------------

	// calculate page count
	pageCount := param.MaxCount/20 + 1
	if pageCount < 1 {
		pageCount = 1
	}
	// get enough search results at first
	locator := page.Locator("button[id='more-results']")
	for i := 0; i < pageCount; i++ {
		// scroll to bottom
		if _, err = page.Evaluate("window.scrollTo(0, document.body.scrollHeight)"); err != nil {
			continue
		}
		// wait for elements to load
		time.Sleep(1 * time.Second) // wait for page to load
		// click "more results" button
		exist, err := locator.IsVisible()
		if err != nil {
			continue
		}
		if exist {
			if err = locator.Click(); err != nil {
				continue
			}
			// wait for elements to load
			time.Sleep(500 * time.Microsecond)
			if err = page.WaitForLoadState(pageWaitOpt); err != nil {
				continue
			}
		} else {
			break // no more results
		}
	}

	// -----------------------------------------------------------
	// parse search results
	// -----------------------------------------------------------

	// find all search result elements
	searchResultList := make([]types.SearchTextResultItem, 0)
	elResultLiList, err := page.Locator("ol[class=react-results--main]").Locator("li[data-layout=organic]").All()
	if err != nil {
		return nil, fmt.Errorf("could not get search results: %w", err)
	}
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

	return &types.SearchTextResponse{
		Results: searchResultList,
	}, nil
}
