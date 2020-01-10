package main

import (
	"io/ioutil"
	"log"
)

func main() {

	indexHtml, err := ioutil.ReadFile("asset/index.html")
	if err != nil {
		log.Printf("Error reading index.html file %s", err)
	}

	formattedIndexHtml := ReplaceTemplate(indexHtml)

	log.Printf("HTML\n%s", formattedIndexHtml)

}
