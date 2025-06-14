package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func initialScrape(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	str := string(body)

	strarr := strings.Split(str, "\n")
	return strarr, nil

}

func ParseWikiLinks(r io.Reader) ([]string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	var wikiLinks []string

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {

			for _, atr := range n.Attr {
				if atr.Key == "href" && strings.HasPrefix(atr.Val, "/wiki/") {
					wikiLinks = append(wikiLinks, "https://en.wikipedia.org"+atr.Val)
				}
			}

		}
	}
	return wikiLinks, nil
}

func scrapeToken(line string, target string) {
	if strings.Contains(line, "href=\"https://") {
		fmt.Println(line)
	}
}

func recursiveScrape(url string, parents []int) {
}
