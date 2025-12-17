package connect

import (
	"log"
	"time"

	"github.com/go-rod/rod"
	"linkedin-automation/state"
)

func SendConnection(page *rod.Page, profileURL string, st *state.State) {
	if st.SentConnections[profileURL] {
		log.Println("Already sent connection to:", profileURL)
		return
	}

	log.Println("Visiting profile:", profileURL)
	page.MustNavigate(profileURL)
	page.MustWaitLoad()

	if page.MustHas(`button:contains("Connect")`) {
		log.Println("Connect button detected (POC only)")
	}

	time.Sleep(5 * time.Second)

	st.MarkConnectionSent(profileURL)
}
