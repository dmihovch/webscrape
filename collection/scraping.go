package collection

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func (mm *MasterMap) Init() error {
	site, err := http.Get(mm.InitUrl)

	if err != nil {
		return err
	}
	defer site.Body.Close()

	doc, err := html.Parse(site.Body)
	if err != nil {
		return err
	}

	mm.Refs.ToRefs[mm.InitUrl] = ParseLinks(doc)

	return nil
}

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

func (mm *MasterMap) GoScrape() {

	for url := range mm.Channel {
		mm.Size++
		if mm.Size >= mm.SizeLimit {
			close(mm.Channel)
			break
		}

	}

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
