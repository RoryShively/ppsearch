package engine

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	pageFile      = "data/pages"
	countDir      = "data/count"
	countFileTmpl = countDir + "/%v"
	pageRowTmpl   = "%v %v\n"
	countRowTmpl  = "%v %v\n"

	maxResults = 20
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
	pageFile := pageFile
	pageRow := fmt.Sprintf(pageRowTmpl, p.url, p.title)
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
		countFile := fmt.Sprintf(countFileTmpl, k)
		row := fmt.Sprintf(countRowTmpl, p.url, v)

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
	_ = os.Remove(pageFile)

	_ = filepath.Walk(countDir, func(path string, info os.FileInfo, err error) error {
		_ = os.Remove(path)
		return nil
	})
}

func (v *DB) Find(term string) *Results {
	res := NewResults(maxResults)
	termFile := fmt.Sprintf(countFileTmpl, term)
	if _, err := os.Stat(termFile); os.IsNotExist(err) {
		return res
	}

	f, err := os.Open(termFile)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		col := strings.Split(row, " ")
		url := col[0]
		countStr := col[1]
		countI32, err := strconv.ParseInt(countStr, 10, 32)
		if err != nil {
			panic(err)
		}
		count := int(countI32)

		page, found := v.FindPage(url)
		if found {
			r := &Result{
				url:   url,
				title: page["title"],
				count: count,
			}
			res.Add(r)
		} else {
			panic("url was found in countDir but not in pageFile")
		}
	}

	return res
}

func (v *DB) FindPage(url string) (map[string]string, bool) {
	f, err := os.Open(pageFile)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		col := strings.Split(row, " ")
		item := map[string]string{
			"url":   string(col[0]),
			"title": string(col[1]),
		}

		if url == item["url"] {
			return item, true
		}
	}
	return nil, false
}
