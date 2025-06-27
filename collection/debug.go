package collection

import (
	"fmt"
)

func (mm *MasterMap) PrintMasterMap() {
	fmt.Println("Initial URL: ", mm.InitUrl)
	fmt.Println("Size: ", mm.Size)
	for url, refs := range mm.Refs.ToRefs {
		fmt.Println()
		fmt.Printf("{%s} Refers To:\n", url)
		for i, u := range refs {
			fmt.Printf("\t %d: {%s}\n", i, u)
		}
	}
}

func (mm *MasterMap) PrintMasterMapStats() {

	refCount := 0
	for _, refs := range mm.Refs.ToRefs {

		refCount += len(refs)

	}

	fmt.Printf("Initial URL: %s\tSize: %d\t Total Refs: %d\n", mm.InitUrl, mm.Size, refCount)
}
