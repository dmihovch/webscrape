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
