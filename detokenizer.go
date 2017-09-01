package main

import (
	"encoding/json"
	"strconv"

	"io/ioutil"

	"fmt"

	"bytes"
)

func mapKeyPairs(input [][]byte, buffer []string) []map[string][]byte {
	tokenMapS := []map[string][]byte{}
	for _, pv := range input {
		tokenMap := map[string]*json.RawMessage{}

		err := json.Unmarshal(pv, &tokenMap)
		if err != nil {
			fmt.Println("Failed to parse one of the JSON token files.")
		}

		tempMap := map[string][]byte{}

		for k, v := range tokenMap {
			j, err := json.Marshal(&v)
			checkError(err)
			tempMap[buffer[0]+k+buffer[1]] = j
		}
		tokenMapS = append(tokenMapS, tempMap)
	}
	return tokenMapS
}

func detokenize(input []byte, tokenMap []map[string][]byte, unquoted bool) []byte {
	overrideCompiled := map[string][]byte{}
	for _, v := range tokenMap {
		for mk, mv := range v {
			if unquoted {
				v, err := strconv.Unquote(string(mv))
				mv = []byte(v)
			}
			overrideCompiled[mk] = mv
		}
	}
	for k, v := range overrideCompiled {
		input = bytes.Replace(input, []byte(k), v, -1)
		input = bytes.Replace(input, []byte("\\u003c"), []byte("<"), -1)
		input = bytes.Replace(input, []byte("\\u003e"), []byte(">"), -1)
		input = bytes.Replace(input, []byte("\\u0026"), []byte("&"), -1)
	}
	input = bytes.Replace(input, []byte(`\\`), []byte(`\`), -1)

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
