package builder

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/umesshk/html-parser/parser"
)

func ParseLinks(link string) []string {

	page_response := GetPage(link)
	defer page_response.Body.Close()

	parsed_links := parser.ParseHtml(page_response.Body)

	reqUrl := page_response.Request.URL

	baseUrl := &url.URL{
		Host:   reqUrl.Host,
		Scheme: reqUrl.Scheme,
	}

	base := baseUrl.String()
	fomatted_links := formatLinks(parsed_links, base)

	page_links := filterLinks(fomatted_links, keepLink(base))

	fmt.Println("Parsed Links ")

	for i, l := range page_links {
		fmt.Printf("%d. %s\n", i+1, l)
	}

	return page_links

}

func formatLinks(parsed_links []parser.Link, base string) []string {

	var page_links []string

	for _, p := range parsed_links {

		switch {
		case strings.HasPrefix(p.Href, "mailto"):
			continue

		case strings.HasPrefix(p.Href, "/"):

			page_links = append(page_links, base+p.Href)

		case strings.HasPrefix(p.Href, "http"):
			page_links = append(page_links, p.Href)

		case p.Href != " " && !strings.HasPrefix(p.Href, "#"):
			page_links = append(page_links, base+"/"+p.Href)

		case strings.HasSuffix(p.Href, "/index.html"):
			trimmed_link := strings.TrimSuffix(p.Href, "/index.html") + "/"
			page_links = append(page_links, trimmed_link)
		}

	}

	return page_links

}

func filterLinks(parsed_links []string, keepLink func(string) bool) []string {

	var valid_links []string

	for _, l := range parsed_links {
		if keepLink(l) {
			valid_links = append(valid_links, l)
		}
	}
	return valid_links
}

func keepLink(prefix string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, prefix)
	}
}
