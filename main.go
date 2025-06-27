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

	Data := c.CreateMasterMap(os.Args[1], 1000)
	Data.Collect()
	Data.PrintMasterMapStats()
}
