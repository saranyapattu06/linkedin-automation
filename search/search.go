package search

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

// SearchProfiles performs a public LinkedIn people search (POC)
func SearchProfiles(page *rod.Page, keyword string, pages int) ([]string, error) {
	profiles := make(map[string]bool)

	for i := 1; i <= pages; i++ {
		searchURL := fmt.Sprintf(
			"https://www.linkedin.com/search/results/people/?keywords=%s&page=%d",
			keyword, i,
		)

		page.MustNavigate(searchURL)
		page.MustWaitLoad()
		time.Sleep(3 * time.Second)

		links := page.MustElements(`a[href*="/in/"]`)
		for _, link := range links {
			href, _ := link.Attribute("href")
			if href != nil {
				profiles[*href] = true
			}
		}
	}

	// Convert map to slice
	result := []string{}
	for url := range profiles {
		result = append(result, url)
	}

	return result, nil
}
