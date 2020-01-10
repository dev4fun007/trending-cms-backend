package main

import (
	"strings"
	"trending-cms-backend/qod"
)

func ReplaceQod(s string) string {

	quote := qod.FetchQod()

	//Insert Quote
	s = strings.ReplaceAll(s, "{{qod.quote}}", quote.Quote)

	//Insert Author
	s = strings.ReplaceAll(s, "{{qod.author}}", quote.Author)

	return s

}
