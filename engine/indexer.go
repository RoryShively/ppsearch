package engine

import (
	"fmt"
	"sync"
)

type Indexer struct {
	pages []*PageParser
	links []string
}

func NewIndexer() *Indexer {
	return &Indexer{}
}

func (v *Indexer) Start(url string) {
	v.parsePage(url, 1)
	for i, p := range v.pages {
		fmt.Println(i, p.url, p.depth)
	}
	fmt.Println(len(v.pages))
}

func (v *Indexer) parsePage(url string, depth int) {
	// var wg sync.WaitGroup
	// linkChannel := getLinkChannel()

	// for {
	// 	select {
	// 	case l := <-linkChannel:
	// 		pp := NewPageParser(url, depth)
	// 		v.pages = append(v.pages, pp)
	// 		wg.Add(1)
	// 		go func(l string) {}
	// 		pp.Start()

	// 	}
	// }

	pp := NewPageParser(url, depth)
	v.pages = append(v.pages, pp)
	pp.Start()

	var wg sync.WaitGroup
	for _, link := range pp.links {
		wg.Add(1)
		go func(l string) {
			v.parsePage(l, depth+1)
			wg.Done()
		}(link)
	}
	wg.Wait()
}
