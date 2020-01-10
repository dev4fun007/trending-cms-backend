package qod

import (
	"encoding/json"
	"log"
	"net/http"
)

const(
	qodUrl = "http://quotes.rest/qod.json"
)

var (
	defaultQuote = Quotation{
		Quote:  "A fool's brain digests philosophy into folly, science into superstition, and art into pedantry. Hence University education.",
		Author: "George Bernard Shaw",
	}
)

type Quotation struct {
	Quote string `json:"quote"`
	Author string `json:"author"`
}

type Response struct {
	Success struct {
		Total int `json:"total"`
	} `json:"success"`
	Contents struct {
		Quotes []struct {
			Quote      string   `json:"quote"`
			Length     string   `json:"length"`
			Author     string   `json:"author"`
			Tags       []string `json:"tags"`
			Category   string   `json:"category"`
			Date       string   `json:"date"`
			Permalink  string   `json:"permalink"`
			Title      string   `json:"title"`
			Background string   `json:"background"`
			ID         string   `json:"id"`
		} `json:"quotes"`
		Copyright string `json:"copyright"`
	} `json:"contents"`
}



func FetchQod() Quotation {
	return getQuote()
}

func getQuote() Quotation {

	response, err := http.Get(qodUrl)
	if err != nil {
		log.Printf("Error fetching quote of the day: %s", err)
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var qodResponse Response
	err = decoder.Decode(&qodResponse)
	if err != nil {
		log.Printf("error converting response to struc %s", err)
	}

	var requiredQuote Quotation
	quotes := qodResponse.Contents.Quotes
	if len(quotes) > 0 {
		requiredQuote.Author = quotes[0].Author
		requiredQuote.Quote = quotes[0].Quote
	} else {
		requiredQuote = defaultQuote
	}
	
	return requiredQuote
}

