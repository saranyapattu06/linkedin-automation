package browser

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func Launch(headless bool) (*rod.Browser, *rod.Page) {
	l := launcher.New().
		Headless(headless).
		Leakless(false) // IMPORTANT: disable leakless to avoid Windows Defender issue

	u := l.MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()
	page := browser.MustPage()

	page.MustWaitLoad()
	time.Sleep(2 * time.Second)

	return browser, page
}
