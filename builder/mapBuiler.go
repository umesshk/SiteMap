package builder

func BuildMap(page_link string, maxDepth int) []string {

	linksSeen := make(map[string]struct{})

	q := make(map[string]struct{})
	nq := map[string]struct{}{
		page_link: {},
	}

	for i := 0; i <= maxDepth; i++ {

		q, nq = nq, make(map[string]struct{})

		if len(q) == 0 {
			break
		}

		for link := range q {

			if _, ok := linksSeen[link]; ok {
				continue
			}

			linksSeen[link] = struct{}{}

			for _, l := range ParseLinks(link) {
				nq[l] = struct{}{}
			}
		}

	}

	seen_links := make([]string, 0, len(linksSeen))
	for l := range linksSeen {
		seen_links = append(seen_links, l)
	}

	return seen_links
}
