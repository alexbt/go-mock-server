package api

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

func HandleRequests(writer http.ResponseWriter, reader *http.Request) {
	var output string

	var content, _ = ioutil.ReadAll(reader.Body)

	var out bytes.Buffer
	json.Indent(&out, content, "", "\t")

	output += "URI: '" + reader.RequestURI + "'\n"
	output += "Method: '" + reader.Method + "'\n"
	output += "Body: '" + out.String() + "'\n"

	output+="Header:\n"
	for name, values := range reader.Header {
		for _, value := range values {
			output += "  " + name + ": '" + value + "'\n"
		}
	}

	println(output)
	writer.Write([]byte(output))
}
