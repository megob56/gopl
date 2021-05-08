package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type Queue []string

func (q *Queue) Pop() (string, Queue) {
	elem := (*q)[1]
	*q = (*q)[1:]
	return elem, *q
}

const URI_STUDENT_HOUSING_BASE_URL = "https://commuters.apps.uri.edu"

// urlQueue keeps track of found urls we have to visit
var urlQueue = make(Queue, 0)

// seenUrls keeps track of what urls we've already visited so we don't visit one multiple times
var seenUrls = make(map[string]bool)

func Crawl() {
	urlQueue = append(urlQueue, URI_STUDENT_HOUSING_BASE_URL)
	for {
		// All paths in the website have been explored
		if len(urlQueue) == 0 {
			break
		}

		// Explore next url
		var url string
		url, urlQueue = urlQueue.Pop()
		if !seenUrls[url] {
			err := Visit(url)
			seenUrls[url] = true
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func Visit(url string) error {
	// Fetch page and parse html
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed GET for url: %s", url)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse html: %s", url)
	}

	// Recursively look for a tags and add their hrefs to the queue
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					extractedUrl := attr.Val
					if len(extractedUrl) > 0 && extractedUrl[0] == '/' {
						extractedUrl = URI_STUDENT_HOUSING_BASE_URL + extractedUrl
					}
					urlQueue = append(urlQueue, extractedUrl)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	for _, value := range urlQueue {
		if strings.Contains(value, "property/listings/info") {
			fmt.Printf("%v\n", value)
		}
	}

	return nil
}

func main() {
	Visit("https://commuters.apps.uri.edu/property/home/myList?rental_period=Academic%20Year")
}
