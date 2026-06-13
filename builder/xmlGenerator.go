package builder

import (
	"encoding/xml"
	"fmt"
	"os"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlSet struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func GenerateXML(links []string) {

	xml_urls := urlSet{
		Xmlns: xmlns,
	}

	for _, l := range links {
		xml_urls.Urls = append(xml_urls.Urls, loc{l})
	}

	fmt.Printf(xml.Header)
	ec := xml.NewEncoder(os.Stdout)
	ec.Indent("", "   ")
	if err := ec.Encode(xml_urls); err != nil {
		panic(err)
	}

	fmt.Println()

}
