package builder

import (
	"io"
	"net/http"
	"strings"
)

func GetPage(link string) io.ReadCloser {

	if !strings.HasPrefix(link, "https://") {
		link = "https://" + link
	}
	resp, err := http.Get(link)

	if err != nil {
		panic(err)
	}

	return resp.Body

}
