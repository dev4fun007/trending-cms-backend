package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {

	indexHtml, err := ioutil.ReadFile("asset/index.html")
	if err != nil {
		log.Printf("Error reading index.html file %s", err)
	}

	formattedIndexHtml := ReplaceTemplate(indexHtml)

	err = ioutil.WriteFile("asset/output/index.html", []byte(formattedIndexHtml), os.FileMode(os.O_CREATE|os.O_RDWR))
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}

	log.Printf("HTML\n%s", formattedIndexHtml)

}
