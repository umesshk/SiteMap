package builder

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

	fmt.Println("Generating XML")

	xml_urls := urlSet{
		Xmlns: xmlns,
	}

	for _, l := range links {
		xml_urls.Urls = append(xml_urls.Urls, loc{l})
	}

	rootDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	file_path := filepath.Join(rootDir, "sitemap.xml")

	file, err := os.Create(file_path)

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(xml.Header)
	ec := xml.NewEncoder(file)

	ec.Indent("", "   ")
	if err := ec.Encode(xml_urls); err != nil {
		panic(err)
	}

	fmt.Println(" XML File Generated...")

}
