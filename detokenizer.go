package main

import (
	"encoding/json"

	"io/ioutil"

	"fmt"

	"bytes"
)

func mapKeyPairs(path, buffer string) map[string][]byte {
	tokenMap := map[string]*json.RawMessage{}
	tokenMapS := map[string][]byte{}
	if path == "" {
		return tokenMapS
	}
	databytes := loadFile(path)
	err := json.Unmarshal(databytes, &tokenMap)
	checkError(err)

	for k, v := range tokenMap {
		j, err := json.Marshal(&v)
		checkError(err)
		tokenMapS[buffer+k+buffer] = j
	}
	return tokenMapS
}

func detokenize(input []byte, tokenMap map[string][]byte) []byte {
	for k, v := range tokenMap {
		input = bytes.Replace(input, []byte(k), v, -1)
	}
	return input
}

func loadFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	checkError(err)
	return file
}

func outputToFile(path string, data []byte) {
	err := ioutil.WriteFile(path, data, 0644)
	checkError(err)
}

func outputToStdout(data []byte) {
	fmt.Print(string(data))
}
