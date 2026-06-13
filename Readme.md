# SiteMap

A simple sitemap generator written in Go.

This project crawls a website starting from a given URL, extracts all internal links, and builds a sitemap by traversing pages up to a specified depth.

The project uses my custom HTML parser to extract links from web pages:

* HTML Parser: https://github.com/umesshk/html-parser

## Features

* Crawl a website starting from a root URL.
* Extract links from HTML pages.
* Traverse pages using breadth-first search (BFS).
* Limit crawling depth to avoid infinite traversal.
* Skip already visited pages.
* Build a graph of discovered pages and their links.
* Supports both relative and absolute URLs.

## How It Works

1. Fetch a web page.
2. Parse all anchor (`<a>`) tags.
3. Extract and normalize URLs.
4. Visit internal links that have not been visited before.
5. Continue crawling until the specified depth is reached.
6. Generate a sitemap structure containing all discovered pages.

## Installation

Clone the repository:

```bash
git clone https://github.com/umesshk/SiteMap
cd SiteMap
```

Install dependencies:

```bash
go mod tidy
```

## Usage

Run the application and provide:

* `-link` : Starting URL to crawl.
* `-depth` : Maximum crawl depth.

Example:

```bash
go run cmd/main.go \
  -link=https://quotes.toscrape.com \
  -depth=3
```

## Example Output

```text
https://example.com
├── https://example.com/about
├── https://example.com/contact
└── https://example.com/blog

https://example.com/blog
├── https://example.com/post1
└── https://example.com/post2
```

## Project Structure

```text
SiteMap/
├── cmd/
│   └── main.go
├── builder/
│   ├── crawler.go
│   ├── sitemap.go
│   └── mapBuilder.go
└── README.md
```

## Future Improvements

* XML sitemap generation.
* Better URL normalization.
