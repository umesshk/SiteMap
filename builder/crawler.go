package builder

import (
	"net/http"
	"strings"
)

func GetPage(link string) *http.Response {

	if !strings.HasPrefix(link, "https://") {
		link = "https://" + link
	}
	resp, err := http.Get(link)

	if err != nil {
		panic(err)
	}

	return resp

}
