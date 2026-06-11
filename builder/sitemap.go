package builder

import (
	"fmt"

	"github.com/umesshk/html-parser/parser"
)

func BuildSiteMap(link string) {
	page_response := GetPage(link)
	defer page_response.Close()

	links := parser.ParseHtml(page_response)

	for _, p := range links {
		fmt.Println(" links: ", p.Href)
		fmt.Println(" to : ", p.Text)
	}
}
