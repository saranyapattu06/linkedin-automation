package message

import (
	"log"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"linkedin-automation/state"
)

func SendMessage(page *rod.Page, profileURL string, template string, st *state.State) {
	if st.SentMessages[profileURL] {
		log.Println("Message already sent to:", profileURL)
		return
	}

	msg := strings.ReplaceAll(template, "{{profile}}", profileURL)
	log.Println("Prepared message:", msg)

	time.Sleep(3 * time.Second)

	st.MarkMessageSent(profileURL)
}
