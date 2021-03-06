package main

import (
	"encoding/json"
	"regexp"

	"bytes"
)

func extractTokens(input []byte, buffer []string) [][]byte {
	bufferBuilder := buffer[0] + "(.*?)" + buffer[1]
	rex := regexp.MustCompile(bufferBuilder)
	return rex.FindAll(input, -1)
}

func convertToJSON(data [][]byte, buffer []string) []byte {
	dmap := map[string]string{}
	for _, d := range data {
		cleanD := bytes.Replace(d, []byte(buffer[0]), []byte{}, -1)
		cleanD = bytes.Replace(cleanD, []byte(buffer[1]), []byte{}, -1)
		dmap[string(cleanD)] = ""
	}
	jdata, err := json.Marshal(dmap)
	checkError(err)
	return jdata
}
