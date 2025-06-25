package collection

import (
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
	CloseFlag bool
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
		Channel:   make(chan string, SizeLimit),
		CloseFlag: false,
	}
}

func (mm *MasterMap) AddNewUrl(NewUrl string) {

	if (mm.Size + 1) == mm.SizeLimit {
		mm.CloseFlag = true
		return
	}

	mm.Size = mm.Size + 1
	mm.Refs.ToRefs[NewUrl] = []string{}

}

func (mm *MasterMap) SetToRefs(Url string, Refs []string) {

	mm.Refs.ToRefs[Url] = Refs

	mm.WaitGroup.Add(1)
	go func(refs []string) {
		defer mm.WaitGroup.Done()
		for _, url := range Refs {
			mm.Refs.Mut.Lock()
			if _, exists := mm.Refs.ToRefs[url]; !exists {
				mm.AddNewUrl(url)
				mm.Refs.Mut.Unlock()
				mm.Channel <- url
			} else {
				mm.Refs.Mut.Unlock()
			}
		}
	}(Refs)

}
