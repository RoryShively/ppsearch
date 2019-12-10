package engine

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

const (
	maxDepth int = 3
)

var blacklist = []string{
	// "https://www.youtube.com/austinpetsalive",
	// "https://www.instagram.com/",
	// "https://www.facebook.com",
}

type PageParser struct {
	data  map[string]int
	title string
	url   *url.URL
	links []string
	depth int // Depth of parser
}

func NewPageParser(uri string, depth int) *PageParser {
	pURL, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	return &PageParser{
		links: []string{},
		data:  map[string]int{},
		depth: depth,
		url:   pURL,
	}
}

func (v *PageParser) Start() {
	v.parsePage()
}

func (v *PageParser) parsePage() {
	fmt.Printf("Parsing %s ...\n", v.url)

	body, err := v.fetchPage(v.url.String())
	if err != nil {
		return // If fetch page errors out here, something outside control of program occurred
	}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	v.parseNode(doc)
}

func (v *PageParser) fetchPage(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		if strings.Contains(err.Error(), "x509: certificate") {
			return nil, err
		}
		fmt.Println("ERROR")
		fmt.Println(uri)
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	_ = ioutil.WriteFile("asdf.html", body, 0644)

	return body, nil
}

func (v *PageParser) parseNode(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					if v.depth < maxDepth {
						v.addLink(a.Val)
					}
					break
				}
			}
		} else if n.Data == "title" {
			if n.FirstChild != nil {
				v.title = n.FirstChild.Data
			} else {
				v.title = v.url.String()
			}
		}
	case html.TextNode:
		v.addData(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			if c.Data == "script" || c.Data == "style" {
				continue
			}
		}
		v.parseNode(c)
	}
}

func (v *PageParser) addLink(link string) {
	// Skip if link is not http or https
	if !strings.HasPrefix(link, "http") {
		return
	}
	for _, url := range blacklist {
		if strings.HasPrefix(link, url) {
			return
		}
	}
	newURL, err := v.url.Parse(link)
	if err != nil {
		panic(err)
	}

	vLinks := getVisitedLinks()
	added := vLinks.AddLink(newURL.String())
	if added {
		v.links = append(v.links, newURL.String())
	}

}

func (v *PageParser) addData(data string) {
	words := strings.Fields(data)

	for _, w := range words {
		w = strings.ToLower(w)
		l := rune(w[0])
		if unicode.IsLetter(l) || unicode.IsNumber(l) {
			if _, ok := v.data[w]; ok {
				v.data[w]++
			} else {
				v.data[w] = 1
			}
		}
	}
}
