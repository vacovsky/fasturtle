package main

import (
	"encoding/json"

	"io/ioutil"

	"fmt"

	"bytes"
	"strings"
)

func mapKeyPairs(input []string, buffer string) []map[string][]byte {
	tokenMapS := []map[string][]byte{}
	if path == "" {
		return tokenMapS
	}
	paths := strings.Split(path, ",")

	for _, pv := range paths {
		tokenMap := map[string]*json.RawMessage{}

		err := json.Unmarshal(input, &tokenMap)
		checkError(err)
		tempMap := map[string][]byte{}

		for k, v := range tokenMap {
			j, err := json.Marshal(&v)
			checkError(err)
			tempMap[buffer+k+buffer] = j
		}
		tokenMapS = append(tokenMapS, tempMap)
	}
	return tokenMapS
}

func detokenize(input []byte, tokenMap []map[string][]byte) []byte {
	for _, v := range tokenMap {
		for mk, mv := range v {
			input = bytes.Replace(input, []byte(mk), mv, -1)
		}
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
