package stealth

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Think simulates human thinking delay
func Think() {
	time.Sleep(randomDuration(800, 2000))
}

func randomDuration(minMs, maxMs int) time.Duration {
	return time.Duration(rand.Intn(maxMs-minMs)+minMs) * time.Millisecond
}
