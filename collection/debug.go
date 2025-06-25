package collection

import (
	"fmt"
)

func (mm *MasterMap) PrintMasterMap() {
	fmt.Println("Initial URL: ", mm.InitUrl)
	for url, refs := range mm.Refs.ToRefs {
		fmt.Println()
		fmt.Printf("{%s} Refers To:\n", url)
		for i, u := range refs {
			fmt.Printf("\t %d: {%s}\n", i, u)
		}
	}
}
