package main

import (
	c "webscrape/collection"
)

func main() {
	/*
		if len(os.Args) < 2 {
			fmt.Println("please provide a url")
			return
		}

		resp, err := http.Get(os.Args[1])
		if err != nil {
			panic(err)
		}

		_, err = c.ParseWikiLinks(resp.Body)
		resp.Body.Close()
		if err != nil {
			panic(err)
		}
	*/
	webData := &c.SyncMap{}
	webData.Init()
	webData.InitialScrapeAndParse("url")
	webData.RecursiveScrape(10)

}
