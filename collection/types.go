package collection

import (
	"fmt"
	"sync"
)

type SyncMap struct {
	Mut     sync.Mutex
	Rmap    map[string]*Refs
	InitUrl string
}

type Refs struct {
	Mut  sync.Mutex
	From []string
	To   []string
}

func (SMap *SyncMap) GetRefs(UrlKey string) *Refs {
	return SMap.Rmap[UrlKey]

}

func (SMap *SyncMap) Init() {
	SMap.Mut = sync.Mutex{}
	SMap.Rmap = make(map[string]*Refs)

}

// returns: 0 created new field in map, 1 field already exists
func (SMap *SyncMap) SetNewKey(Url string) bool {
	SMap.Mut.Lock()
	defer SMap.Mut.Unlock()
	_, exists := SMap.Rmap[Url]
	if !exists {
		SMap.Rmap[Url] = &Refs{}
		return true
	}
	return false
}
func (SMap *SyncMap) SetToRefs(UrlKey string, Urls []string) error {

	refs := SMap.Rmap[UrlKey]
	if refs == nil {
		return fmt.Errorf("UrlKey %s does not exist", UrlKey)
	}
	SMap.Rmap[UrlKey].Mut.Lock()
	defer SMap.Rmap[UrlKey].Mut.Unlock()
	SMap.Rmap[UrlKey].To = Urls
	return nil
}
func (SMap *SyncMap) SetFromRefs(UrlKey string, Urls []string) error {

	refs := SMap.Rmap[UrlKey]
	if refs == nil {
		return fmt.Errorf("UrlKey %s does not exist", UrlKey)
	}
	SMap.Rmap[UrlKey].Mut.Lock()
	defer SMap.Rmap[UrlKey].Mut.Unlock()
	SMap.Rmap[UrlKey].From = Urls
	return nil
}

func (SMap *SyncMap) SetAllRefs(UrlKey string, ToUrls []string, FromUrls []string) error {
	refs := SMap.Rmap[UrlKey]
	if refs == nil {
		return fmt.Errorf("UrlKey %s does not exist", UrlKey)
	}
	SMap.Rmap[UrlKey].Mut.Lock()
	defer SMap.Rmap[UrlKey].Mut.Unlock()
	SMap.Rmap[UrlKey].From = FromUrls
	SMap.Rmap[UrlKey].To = ToUrls
	return nil
}
