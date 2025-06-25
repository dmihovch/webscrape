package collection

import (
	"sync"
)

//going to refactor with a tree structure

type MasterMap struct {
	InitUrl    string
	SizeLimit  int
	Size       int
	MasterList *Ledger
	Refs       *SyncMap
	Channel    chan (string)
}

type Ledger struct {
	Mut     sync.Mutex
	AllUrls []string
}

type SyncMap struct {
	Mut    sync.Mutex
	ToRefs map[string][]string
}

func CreateMasterMap(InitUrl string, SizeLimit int) *MasterMap {
	return &MasterMap{
		InitUrl:   InitUrl,
		SizeLimit: SizeLimit,
		Size:      1,
		Refs: &SyncMap{
			Mut:    sync.Mutex{},
			ToRefs: make(map[string][]string),
		},
		Channel: make(chan (string)),
	}
}

func (mm *MasterMap) AddNewUrl(NewUrl string) {
	mm.Channel <- NewUrl

	//this will cause big slow downs??
	mm.Refs.Mut.Lock()
	mm.Refs.ToRefs[NewUrl] = []string{}
	mm.Refs.Mut.Unlock()
}

func (mm *MasterMap) AddToRefs(Url string, Refs []string) {
	mm.Refs.Mut.Lock()
	defer mm.Refs.Mut.Unlock()
	for _, url := range Refs {
		mm.Refs.ToRefs[Url] = append(mm.Refs.ToRefs[Url], url)
	}

}
