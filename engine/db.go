package engine

import (
	"fmt"
	"os"
	"path/filepath"
)

func init() {
	os.MkdirAll("data/count", os.ModePerm)
}

type DB struct{}

func GetDB() *DB {
	return &DB{}
}

func (v *DB) SaveIndex(idx *Indexer) {
	for _, p := range idx.pages {
		v.savePage(p)
	}
}

func (v *DB) savePage(p *PageParser) {
	pageFile := "data/pages"
	pageRow := fmt.Sprintf("%v %v\n", p.url, p.title)
	f, err := os.OpenFile(pageFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	if _, err := f.WriteString(pageRow); err != nil {
		panic(err)
	}
	f.Close()

	for k, v := range p.data {
		countFile := fmt.Sprintf("data/count/%v", k)
		row := fmt.Sprintf("%v %v\n", p.url, v)

		f, err := os.OpenFile(countFile,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		if _, err := f.WriteString(row); err != nil {
			panic(err)
		}
		f.Close()
	}
}

func (v *DB) Clear() {
	_ = os.Remove("data/pages")

	_ = filepath.Walk("data/count", func(path string, info os.FileInfo, err error) error {
		_ = os.Remove(path)
		return nil
	})
}
