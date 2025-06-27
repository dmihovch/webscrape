package collection

import (
	"fmt"
	"sync"
)

//going to refactor with a tree structure

type MasterMap struct {
	WaitGroup sync.WaitGroup
	InitUrl   string
	SizeLimit int
	Size      int
	Refs      *SyncMap
	Channel   chan (string)
}

type SyncMap struct {
	Mut    sync.Mutex
	ToRefs map[string][]string
}

func CreateMasterMap(InitUrl string, SizeLimit int) *MasterMap {
	return &MasterMap{
		WaitGroup: sync.WaitGroup{},
		InitUrl:   InitUrl,
		SizeLimit: SizeLimit,
		Size:      1,
		Refs: &SyncMap{
			Mut:    sync.Mutex{},
			ToRefs: make(map[string][]string),
		},
		Channel: make(chan string, SizeLimit),
	}
}

func (mm *MasterMap) AddNewUrl(NewUrl string) {

	mm.Refs.Mut.Lock()

	defer mm.Refs.Mut.Unlock()

	if _, exists := mm.Refs.ToRefs[NewUrl]; exists {
		return
	}

	if mm.Size >= mm.SizeLimit {
		return
	}

	mm.Size++
	mm.Refs.ToRefs[NewUrl] = []string{}
	fmt.Println("Sending url to channel...")
	mm.Channel <- NewUrl
	fmt.Println("Done")

}

func (mm *MasterMap) SetToRefs(Url string, Refs []string) {
	fmt.Println("attemping to hold lock in setToRefs")
	mm.Refs.Mut.Lock()
	mm.Refs.ToRefs[Url] = Refs
	mm.Refs.Mut.Unlock()
	fmt.Println("letting go hold lock in setToRefs")

	for _, url := range Refs {
		mm.AddNewUrl(url)
	}

	/*

		going to leave it here because I think if done right, will speed things up, but I think its done wrong and breaking shit rn

		go func(refs []string) {
				defer mm.WaitGroup.Done()
				for _, url := range Refs {
					mm.AddNewUrl(url)
				}
			}(Refs)

	*/

}
