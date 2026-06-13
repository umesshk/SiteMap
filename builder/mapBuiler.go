package builder

func BuildMap(page_link string, maxDepth int) map[string][]string {

	linksMap := make(map[string][]string)
	linksSeen := make(map[string]struct{})

	linkDiscoverd := map[string]struct{}{
		page_link: struct{}{},
	}

	q := make(map[string]struct{})
	nq := map[string]struct{}{
		page_link: struct{}{},
	}

	for i := 0; i < maxDepth; i++ {

		q, nq = nq, make(map[string]struct{})
		for link, _ := range q {

			if _, ok := linksSeen[link]; ok {
				continue
			}

			linksSeen[link] = struct{}{}
			curr_link := link

			for _, l := range ParseLinks(link) {
				linksMap[curr_link] = append(linksMap[curr_link], l)
				if _, ok := linkDiscoverd[l]; ok {
					continue
				}
				linkDiscoverd[l] = struct{}{}
				nq[l] = struct{}{}
			}

		}
	}

	return linksMap
}
