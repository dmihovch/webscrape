package collection

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func ParseLinks(doc *html.Node) []string {
	var wikiLinks []string

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {

			for _, atr := range n.Attr {
				if atr.Key == "href" && strings.HasPrefix(atr.Val, "/wiki/") && !strings.Contains(atr.Val, ":") {
					wikiLinks = append(wikiLinks, "https://en.wikipedia.org"+atr.Val)
				}
			}

		}
	}
	return wikiLinks
}

func (mm *MasterMap) Collect() {

	go func() {
		mm.Channel <- mm.InitUrl
	}()

	mm.WaitGroup.Add(1)
	go mm.ScrapeLoop()

	mm.WaitGroup.Wait()

}

func (mm *MasterMap) ScrapeLoop() {
	defer mm.WaitGroup.Done()
	fmt.Println("In Scrape Loop")

	for {
		fmt.Println("about to block")
		select {
		case NewUrl, ok := <-mm.Channel:
			if !ok {
				return
			}
			fmt.Println(NewUrl, ": new url from channel")

			fmt.Println(NewUrl, ": about to scrape")
			mm.WaitGroup.Add(1)
			go func(url string) {
				defer mm.WaitGroup.Done()
				mm.GoScrape(url)
			}(NewUrl)
			fmt.Println(NewUrl, ": scraped url")

		}
	}

}

func (mm *MasterMap) GoScrape(url string) {

	site, err := http.Get(url)

	if err != nil {
		delete(mm.Refs.ToRefs, url)
		return
	}

	doc, err := html.Parse(site.Body)
	if err != nil {
		delete(mm.Refs.ToRefs, url)
		return
	}
	site.Body.Close()

	mm.SetToRefs(url, ParseLinks(doc))

}
