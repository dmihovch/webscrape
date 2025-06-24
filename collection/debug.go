package collection

import (
	"fmt"
)

func (sm *SyncMap) PrintSyncMap() {
	fmt.Printf("Init Url: [%s]\n\n", sm.InitUrl)

	for url, refs := range sm.Rmap {
		fmt.Printf("[%s] points to:\n", url)
		for i, tourl := range refs.To {
			fmt.Printf("\t%d: [%s]\n", i, tourl)
		}
		fmt.Printf("\n\n\n")
	}

}
