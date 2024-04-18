package main

import (
	"fmt"

	// import colly
	"github.com/gocolly/colly"
)

func main() {
	// scrapping logic...
	fmt.Println("Scraping logic...")
	c := colly.NewCollector()
	c.Visit("https://en.wikipedia.org/wiki/Main_Page")

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		// Print all URLs associated with the a links in the page
		fmt.Printf("%v", e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, "scraped!")
	})
}
