package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("please provide a url")
		return
	}

	resp, err := http.Get(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ParseWikiLinks(resp.Body)
	if err != nil {
		panic(err)
	}

	for _, d := range data {
		fmt.Printf("%s\n", d)
	}
}
