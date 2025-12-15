package stealth

import "github.com/go-rod/rod"

// ApplyFingerprintMask removes common automation fingerprints
func ApplyFingerprintMask(page *rod.Page) {
	page.MustEvalOnNewDocument(`
		Object.defineProperty(navigator, 'webdriver', {
			get: () => undefined
		});

		Object.defineProperty(navigator, 'languages', {
			get: () => ['en-US', 'en']
		});

		Object.defineProperty(navigator, 'platform', {
			get: () => 'Win32'
		});
	`)
}
