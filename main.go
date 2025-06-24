package main

import (
	"fmt"

	"os"
	c "webscrape/collection"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("please provide a url")
		return
	}
	webData := &c.SyncMap{}
	webData.Init(1, os.Args[1])
	webData.InitialScrapeAndParse()
	webData.RecursiveScrape(webData.InitUrl)

	webData.PrintSyncMap()

}
