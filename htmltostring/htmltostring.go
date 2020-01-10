package htmltostring

import (
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func convertHtmlToString(reader io.Reader) string {

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Printf("Error reading the stream %s", err)
	}

	return sanitize(string(data))

}


func sanitize(data string) string {

	//Replace " with \"
	data = strings.ReplaceAll(data, "\"", "\\\"")

	//Replace </ with <\/
	data = strings.ReplaceAll(data, "</", "<\\/")

	//Replace /> with \/>
	data = strings.ReplaceAll(data, "/>", "\\/>")

	//Replace \n with empty character
	data = strings.ReplaceAll(data, "\r\n", "")
	data = strings.ReplaceAll(data, "\n", "")

	//data = strings.ReplaceAll(data, " ", "")

	return data
}