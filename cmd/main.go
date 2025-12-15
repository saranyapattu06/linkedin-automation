package main

import (
	"linkedin-automation/browser"
	"linkedin-automation/config"
	"linkedin-automation/stealth"
	"time"
)

func main() {
	cfg := config.LoadConfig()

	browserInstance, page := browser.Launch(cfg.Headless)
	defer browserInstance.MustClose()

	stealth.ApplyFingerprintMask(page)

	page.MustNavigate("https://www.linkedin.com")
	page.MustWaitLoad()

	stealth.Think()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		stealth.MoveMouseHuman(page)
	}
}
