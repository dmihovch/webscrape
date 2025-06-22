package collection

import (
	"io"
)

func CollectData(url string, r io.Reader) ([]string, error) {
	links, err := ParseWikiLinks(r)
	if err != nil {
		return links, err
	}
	return links, err

}
