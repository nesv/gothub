package gothub

import (
	"net/http"
	"regexp"
	"strings"
)

var (
	linkRegex = regexp.MustCompile(`^<(?P<url>https?:\/\/[^>]+)>;\s+rel="(?P<rel>next|last|first|prev)"$`)
)

type httpLink struct {
	Url string
	Rel string
}

func parseLinkHeader(response *http.Response) (links []*httpLink) {
	locations := strings.Split(response.Header.Get("Link"), ",")
	for _, v := range locations {
		matches := linkRegex.FindStringSubmatch(strings.Trim(v, " "))
		link := &httpLink{Url: matches[1], Rel: matches[2]}
		links = append(links, link)
	}
	return
}
