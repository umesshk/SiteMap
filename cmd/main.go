package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/umesshk/SiteMap/builder"
)

func main() {

	page_url := flag.String("link", "", "Provide the url of page you want to build site map")
	max_depth := flag.Int("depth", 3, "provide max depth of links to follow from a  page (max 10)")

	flag.Parse()

	if *page_url == "" {
		fmt.Println("Please Provide a URL")
		os.Exit(-1)

	}

	if *max_depth >= 10 {
		fmt.Println("Max depth limited to 10")
		os.Exit(-1)
	}

	links := builder.BuildMap(*page_url, *max_depth)

	builder.GenerateXML(links)

}
