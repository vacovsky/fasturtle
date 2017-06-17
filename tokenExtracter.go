package main

import (
	"encoding/json"
	"regexp"

	"bytes"

	"github.com/davecgh/go-spew/spew"
)

// __(.*?)__
// replace __ with buffers

func extractTokens(input []byte, buffer string) [][]byte {
	bufferBuilder := buffer + "(.*?)" + buffer
	rex := regexp.MustCompile(bufferBuilder)
	return rex.FindAll(input, -1)
}

func convertToJSON(data [][]byte, buffer string) []byte {
	dmap := map[string]string{}
	for _, d := range data {
		cleanD := bytes.Replace(d, []byte(buffer), []byte{}, -1)
		dmap[string(cleanD)] = ""
	}
	jdata, err := json.Marshal(dmap)
	checkError(err)
	spew.Dump(jdata)
	return jdata
}
