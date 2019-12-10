package engine

import (
	"sync"
)

type linksSingleton struct {
	visited []string
	mu      sync.Mutex
}

func (v *linksSingleton) AddLink(link string) (added bool) {
	v.mu.Lock()
	defer v.mu.Unlock()

	_, found := Find(v.visited, link)
	if !found {
		v.visited = append(v.visited, link)
		added = true
	}
	return
}

var visitedLinksSingleton *linksSingleton
var visitedLinksOnce sync.Once

func getVisitedLinks() *linksSingleton {
	visitedLinksOnce.Do(func() {
		visitedLinksSingleton = &linksSingleton{
			visited: []string{},
		}
	})
	return visitedLinksSingleton
}

var linkChannelSingleton chan string
var linkChannelOnce sync.Once

func getLinkChannel() chan string {
	linkChannelOnce.Do(func() {
		linkChannelSingleton = make(chan string)
	})
	return linkChannelSingleton
}
