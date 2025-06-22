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

	webData.SetNewKey("Test")
	webData.SetAllRefs("Test", []string{"to", "to"}, []string{"from", "from"})

	println(webData.Rmap["Test"].From[0])
	println(webData.Rmap["Test"].To[0])

	println(webData.Rmap["Test"].From[1])
	println(webData.Rmap["Test"].To[1])

}
