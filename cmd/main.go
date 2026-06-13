package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/umesshk/SiteMap/builder"
)

func main() {

	page_url := flag.String("link", "google.com", "Provide the url of page you want to build site map")
	max_depth := flag.Int("depth", 1, "provide max depth of links to follow from a  page ")

	flag.Parse()

	if *page_url == "" {
		fmt.Println("Please Provide a URL")
		os.Exit(-1)

	}
	fmt.Println("Building Site Map...")

	linksMap := builder.BuildMap(*page_url, *max_depth)

	fmt.Printf("\nSite Map of depth %d for %s\n", *max_depth, *page_url)

	for k, v := range linksMap {
		fmt.Println("Link : ", k)
		for i, l := range v {
			fmt.Printf("\t %d.  %s\n", i+1, l)
		}

	}

}
