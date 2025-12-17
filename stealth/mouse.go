package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func MoveMouseHuman(page *rod.Page) {
	actions := rand.Intn(3) + 3 // 3–5 small scroll actions

	for i := 0; i < actions; i++ {
		time.Sleep(time.Duration(rand.Intn(600)+400) * time.Millisecond)

		scrollY := float64(rand.Intn(40) + 10)
		if rand.Intn(2) == 0 {
			scrollY = -scrollY
		}
		page.Mouse.Scroll(0, scrollY, 1)
	}
}
