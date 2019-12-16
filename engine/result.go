package engine

import "fmt"

type Results struct {
	data []*Result
	max  int
}

func NewResults(max int) *Results {
	return &Results{
		data: []*Result{},
		max:  max,
	}
}

func (v *Results) Add(res *Result) {
	// Add first elem to array
	if len(v.data) == 0 {
		v.data = append(v.data, res)
		return
	}

	var added bool
	for idx, r := range v.data {
		if res.count > r.count {
			v.data = append(v.data[:idx], append([]*Result{res}, v.data[idx:len(v.data)]...)...)
			added = true
			break
		}
	}

	if len(v.data) < v.max && !added {
		v.data = append(v.data, res)
	}
}

func (v *Results) Print() {
	for _, r := range v.data {
		r.Print()
	}
	fmt.Printf("Total Count: %v\n", len(v.data))
}

type Result struct {
	title string
	url   string
	count int
}

func (v *Result) Print() {
	fmt.Printf("Title:      %v\nUrl:        %v\nTerm Count: %v\n\n", v.title, v.url, v.count)
}
