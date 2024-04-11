package searchs

import (
	"github.com/playwright-community/playwright-go"
	"io"
	"log"
	"os"
	"time"
)

func Lab() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	// 启动浏览器
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
		//Proxy:    &playwright.Proxy{Server: "http://127.0.0.1:7890"},
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	// 创建新上下文
	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		Locale: playwright.String("en-Us"),
		Viewport: &playwright.Size{
			Width:  1920,
			Height: 1080,
		},
	})
	if err != nil {
		log.Fatalf("could not create context: %v", err)
	}
	// 添加 stealth.min.js
	stealthJs, err := readLocalFile("./resource/stealth.min.js")
	if err != nil {
		log.Fatalf("could not read stealth.min.js: %v", err)
	}
	stealthJsScript := playwright.Script{
		Content: playwright.String(stealthJs),
	}
	if err = context.AddInitScript(stealthJsScript); err != nil {
		log.Fatalf("could not add stealth.min.js: %v", err)
	}
	// 创建新页面
	page, err := context.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	// 打开首页 输入关键字 搜索
	if _, err = page.Goto("https://duckduckgo.com/"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	waitOpt := playwright.PageWaitForLoadStateOptions{
		State: playwright.LoadStateNetworkidle,
	}
	if err = page.WaitForLoadState(waitOpt); err != nil {
		log.Fatalf("could not wait for load: %v", err)
	}
	time.Sleep(1 * time.Second)
	keyword := "ChatGPT"
	if err = page.Fill("input[name=q]", keyword); err != nil {
		log.Fatalf("could not fill: %v", err)
	}
	if err = page.Keyboard().Press("Enter"); err != nil {
		log.Fatalf("could not press Enter: %v", err)
	}

	// ----------------------
	maxCount := 500
	// ----------------------

	pageCount := maxCount / 20

	waitOpt = playwright.PageWaitForLoadStateOptions{
		State: playwright.LoadStateNetworkidle,
	}
	if err = page.WaitForLoadState(waitOpt); err != nil {
		log.Fatalf("could not wait for load: %v", err)
	}
	time.Sleep(1 * time.Second)

	locator := page.Locator("button[id='more-results']")
	for i := 0; i < pageCount; i++ {
		// 页面滚动到底部
		if _, err = page.Evaluate("window.scrollTo(0, document.body.scrollHeight)"); err != nil {
			log.Fatalf("could not scroll to bottom: %v", err)
		}
		time.Sleep(1 * time.Second)

		exist, err := locator.IsVisible()
		if err != nil {
			log.Fatalf("could not check if visible: %v", err)
		}
		if exist {
			if err = locator.Click(); err != nil {
				log.Fatalf("could not click: %v", err)
			}
			if err = page.WaitForLoadState(waitOpt); err != nil {
				log.Fatalf("could not wait for load: %v", err)
			}
		} else {
			break
		}
	}

	// 判断是否存在
	// 查询
	// 找到 class="react-results--main" 的 ol 元素
	elResultMain, err := page.WaitForSelector(".react-results--main")
	if err != nil {
		log.Fatalf("could not find .react-results--main: %v", err)
	}

	// 遍历 elResultMain 下的所有 li 元素
	elsResultLi, err := elResultMain.QuerySelectorAll("li")
	if err != nil {
		log.Fatalf("could not find li: %v", err)
	}
	for _, elResultLi := range elsResultLi {
		// 找到 data-testid="result-title-a" 的 a 元素
		elResultTitleA, err := elResultLi.QuerySelector("a[data-testid=result-title-a]")
		if err != nil {
			log.Fatalf("could not find [data-testid=result-title-a]: %v", err)
		}
		if elResultTitleA == nil {
			continue
		}
		// 获取 a 元素的文本
		title, err := elResultTitleA.TextContent()
		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}
		// 获取 a 元素的 href 属性
		href, err := elResultTitleA.GetAttribute("href")
		if err != nil {
			log.Fatalf("could not get attribute: %v", err)
		}

		// 找到 data-result="snippet" 的 div 元素
		elResultSnippet, err := elResultLi.QuerySelector("[data-result=snippet]")
		if err != nil {
			log.Fatalf("could not find [data-result=snippet]: %v", err)
		}
		// 获取 div 元素的纯文本
		description, err := elResultSnippet.TextContent()
		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}

		log.Printf("title: %s, href: %s, description: %s", title, href, description)
	}

	//if _, err = page.Goto("https://bot.sannysoft.com/"); err != nil {
	//	log.Fatalf("could not goto: %v", err)
	//}
	//time.Sleep(5 * time.Second)
	//// 保持截图
	//_, err = page.Screenshot(playwright.PageScreenshotOptions{
	//	Path: playwright.String("screenshot_default.png"),
	//})
	//if err != nil {
	//	log.Fatalf("could not take screenshot: %v", err)
	//}

	time.Sleep(1 * time.Hour)

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}

}

func readLocalFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
