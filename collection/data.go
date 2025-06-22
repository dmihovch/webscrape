package collection

import (
	"io"
)

func CollectData(url string, r io.Reader) (*Refs, error) {
	page := &Refs{}
	links, err := ParseWikiLinks(r)
	if err != nil {
		return page, err
	}
	page.To = links

	return page, err

}
