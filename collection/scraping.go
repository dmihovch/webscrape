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

	go mm.ScrapeLoop()

	mm.WaitGroup.Add(1)
	go func() {
		defer mm.WaitGroup.Done()
		mm.GoScrape(mm.InitUrl)
	}()

	fmt.Println("Waiting now")
	mm.WaitGroup.Wait()
	fmt.Println("Done waiting")

}

func (mm *MasterMap) ScrapeLoop() {

	fmt.Println("Scrape Loop Running")
	fmt.Println(len(mm.Channel))

	for NewUrl := range mm.Channel {
		fmt.Println("Recieved URL")
		mm.WaitGroup.Add(1)
		go func(url string) {
			defer mm.WaitGroup.Done()
			fmt.Println("Entering GoScrape")
			mm.GoScrape(url)
		}(NewUrl)
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

func (mm *MasterMap) SetToRefs(Url string, Refs []string) {
	//fmt.Println("attemping to hold lock in setToRefs")
	mm.Refs.Mut.Lock()
	mm.Refs.ToRefs[Url] = Refs
	mm.Refs.Mut.Unlock()
	//fmt.Println("letting go hold lock in setToRefs")

	for _, url := range Refs {
		mm.AddNewUrl(url)
	}

}

func (mm *MasterMap) AddNewUrl(NewUrl string) {

	mm.Refs.Mut.Lock()

	defer mm.Refs.Mut.Unlock()

	if _, exists := mm.Refs.ToRefs[NewUrl]; exists {
		fmt.Println("????")
		return
	}

	if mm.Size >= mm.SizeLimit {
		return
	}

	mm.Size++
	mm.Refs.ToRefs[NewUrl] = []string{}
	fmt.Println("Sending url to channel...")
	mm.Channel <- NewUrl

}
