package collection

import (
	"fmt"
	"sync"
)

//going to refactor with a tree structure

type WebTree struct {
	initUrl string
}

type Node struct {
	url         string
	numChildren int
	children    []*Node
}

func CreateNode(url string, numChildren int, children []*Node) *Node {
	return &Node{url, numChildren, children}
}

type SyncMap struct {
	Mut     sync.Mutex
	Rmap    map[string]*Refs
	InitUrl string
	Depth   int
	Cntr    *Counter
}

type Refs struct {
	Mut  sync.Mutex
	From []string
	To   []string
}

type Counter struct {
	Mut   sync.Mutex
	Iters int
}

func (SMap *SyncMap) GetRefs(UrlKey string) *Refs {
	return SMap.Rmap[UrlKey]

}

func (SMap *SyncMap) Init(depth int, initUrl string) {
	SMap.Mut = sync.Mutex{}
	SMap.Rmap = make(map[string]*Refs)
	SMap.Depth = depth
	SMap.InitUrl = initUrl
	SMap.Rmap[initUrl] = &Refs{}
	SMap.Cntr = &Counter{}
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
