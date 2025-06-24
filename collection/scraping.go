package collection

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func (smap *SyncMap) InitialScrapeAndParse() {
	resp, err := http.Get(smap.InitUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	smap.Rmap[smap.InitUrl].To = parseLinks(doc)

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

func (smap *SyncMap) RecursiveScrape(url string) {
	for _, urlinmap := range smap.Rmap[url].To {

		go func() {

			ok := smap.SetNewKey(urlinmap)
			if !ok {
				fmt.Println(urlinmap)
				panic("key already exists in recur scrape")
			}

			resp, err := http.Get(urlinmap)
			if err != nil {
				panic(err)
			}

			doc, err := html.Parse(resp.Body)
			if err != nil {
				panic(err)
			}
			resp.Body.Close()
			smap.SetToRefs(urlinmap, parseLinks(doc))
			smap.RecursiveScrape(urlinmap)
		}()
	}
}
