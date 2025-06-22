package collection

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func (smap *SyncMap) InitialScrapeAndParse(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	smap.InitUrl = url

	smap.Rmap[url].To = parseLinks(doc)

}

func parseLinks(doc *html.Node) []string {
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

func (smap *SyncMap) RecursiveScrape(depth int) {
	for i, url := range smap.Rmap[smap.InitUrl].To {
		if i >= depth {
			break
		}
		go func() {
			ok := smap.SetNewKey(url)
			if !ok {
				panic("key already exists in recur scrape")
			}

			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}

			doc, err := html.Parse(resp.Body)
			if err != nil {
				panic(err)
			}
			resp.Body.Close()
			smap.SetToRefs(url, parseLinks(doc))

		}()

}
