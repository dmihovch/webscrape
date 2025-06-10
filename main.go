package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func initialScape(url string) ([]string, error) {
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

	strarr := strings.Split(str, " ")
	return strarr, nil

}

func recursiveScrape(url string, parents []int) {
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("please provide a url")
		return
	}

	tokens, err := initialScape(os.Args[1])
	if err != nil {
		panic(err)
	}

	for i, tok := range tokens {
		fmt.Printf("%d: %s\n", i, tok)
	}
}
