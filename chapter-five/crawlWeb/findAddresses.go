package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Make HTTP request
	response, err := http.Get("https://commuters.apps.uri.edu/property/home/myList?rental_period=Academic%20Year")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find("h3").Each(func(index int, element *goquery.Selection) {
		address := element.Text()
		fmt.Println(address)
	})
}

// 	for _, url := range os.Args[1:] {
// 		// if strings.HasPrefix(url, "https://") == false || strings.HasPrefix(url, "http://") == false {
// 		// 	url = "https://" + url
// 		// }
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 			os.Exit(1)
// 		}
// 		b, err := ioutil.ReadAll(resp.Body)
// 		resp.Body.Close()
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
// 			os.Exit(1)
// 		}
// 		str := string(b)
// 		// var listOfAddresses []string

// 		// listOfAddresses.append(GetStringInBetween(str, "<h3>", "</h3"))

// 		fmt.Print(GetStringInBetween(str, "<h3>", "</h3"))
// 		// 	for i := 0; i < len(str); i++ {
// 		// 		// if string(str[i]) == "h" && string(str[i+1]) == "3" {
// 		// 		// 	fmt.Println(string(str[i+2]))
// 		// 		// }
// 		// 		// strings.TrimLeft(strings.TrimRight(initial,"</h3>"),"<h3>")
// 		// 	}
// 		// 	// fmt.Printf("%s", b)
// 		// }
// 	}
// }

// func GetStringInBetween(str string, start string, end string) (result string) {
// 	s := strings.Index(str, start)
// 	if s == -1 {
// 		return
// 	}
// 	s += len(start)
// 	e := strings.Index(str[s:], end)
// 	if e == -1 {
// 		return
// 	}
// 	return str[s:e]
// }
