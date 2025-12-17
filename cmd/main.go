package main

import (
	"fmt"
	"log"

	"linkedin-automation/browser"
	"linkedin-automation/search"
	"linkedin-automation/stealth"
)

func main() {
	// Launch browser (headful for demo)
	browserInstance, page := browser.Launch(false)
	defer browserInstance.MustClose()

	// Apply stealth
	stealth.ApplyFingerprintMask(page)

	// Perform search (POC)
	results, err := search.SearchProfiles(page, "software engineer", 2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Collected profile URLs:")
	for _, url := range results {
		fmt.Println(url)
	}

	// Keep browser open for inspection
	select {}
}
