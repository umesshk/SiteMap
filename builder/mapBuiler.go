package builder

import "fmt"

func BuildMap(page_link string, maxDepth int) {

	linksMap := make(map[string][]string)
	linksSeen := make(map[string]struct{})

	q := make(map[string]struct{})
	nq := map[string]struct{}{
		page_link: struct{}{},
	}

	for i := 0; i < maxDepth; i++ {

		q, nq = nq, make(map[string]struct{})
		for link, _ := range q {

			if _, ok := linksSeen[link]; !ok {

				linksSeen[link] = struct{}{}
				curr_link := link

				for _, l := range ParseLinks(link) {
					linksMap[curr_link] = append(linksMap[curr_link], l)
					nq[l] = struct{}{}
				}

			}
		}
	}

	for k, v := range linksMap {
		fmt.Println(k)
		for _, l := range v {
			fmt.Println("-->", l)
		}
	}

}
