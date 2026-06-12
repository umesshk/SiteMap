package builder

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/umesshk/html-parser/parser"
)

func BuildSiteMap(link string) {
	page_response := GetPage(link)
	defer page_response.Body.Close()

	parserd_links := parser.ParseHtml(page_response.Body)

	var links []string

	reqUrl := page_response.Request.URL

	baseUrl := &url.URL{
		Host:   reqUrl.Host,
		Scheme: reqUrl.Scheme,
	}

	base := baseUrl.String()
	for _, p := range parserd_links {

		switch {
		case strings.HasPrefix(p.Href, "/"):
			links = append(links, base+p.Href)

		case strings.HasPrefix(p.Href, "/http"):
			links = append(links, p.Href)
		}
	}

	fmt.Println("\nFound URLS ....")

	for i, u := range links {
		fmt.Printf("%d. %s \n", i+1, u)
	}

}
