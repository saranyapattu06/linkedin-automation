package auth

import (
	"errors"
	"time"

	"github.com/go-rod/rod"
)

func Login(page *rod.Page, email, password string) error {
	page.MustNavigate("https://www.linkedin.com/login")
	page.MustWaitLoad()

	page.MustElement(`#username`).MustInput(email)
	time.Sleep(1 * time.Second)

	page.MustElement(`#password`).MustInput(password)
	time.Sleep(1 * time.Second)

	page.MustElement(`button[type="submit"]`).MustClick()
	page.MustWaitLoad()

	if page.MustHas(`.alert.error`) || page.MustHas(`div[data-test-id="error"]`) {
		return errors.New("invalid credentials detected")
	}

	if page.MustHas(`input[name="pin"]`) {
		return errors.New("2FA detected, manual intervention required")
	}

	if page.MustHas(`iframe[src*="captcha"]`) {
		return errors.New("captcha detected, stopping automation")
	}

	return nil
}
