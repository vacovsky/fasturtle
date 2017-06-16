package main

import (
	"encoding/json"
	"log"

	"io/ioutil"

	"fmt"

	"bytes"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// Parse the command line flag values into variables for later use
	args := flagInit()

	// convert JSON file into a map[string]string for later use
	tokens := mapKeyPairs(*args.tokensPath, *args.bufferChars)

	// load tokenized document into memory
	input := loadTokenizedDocument(*args.inputPath)

	// store final product for later use
	output := detokenize(input, tokens)
	if *args.outputPath != "" {
		outputDetokenizedDocumentToFile(*args.outputPath, output)
	} else {
		outputDetokenizedDocumentToStdout(output)
	}
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func mapKeyPairs(path, buffer string) map[string][]byte {
	tokenMap := map[string]*json.RawMessage{}
	tokenMapS := map[string][]byte{}
	databytes := loadFile(path)
	err := json.Unmarshal(databytes, &tokenMap)
	checkError(err)

	for k, v := range tokenMap {
		j, err := json.Marshal(&v)
		checkError(err)
		// fmt.Println(k, string(j))
		// j = append(j, []byte(buffer)...)
		// j = append([]byte(buffer), j...)
		tokenMapS[buffer+k+buffer] = j
	}
	spew.Dump(tokenMapS)
	return tokenMapS
}

func detokenize(input []byte, tokenMap map[string][]byte) []byte {
	for k, v := range tokenMap {
		input = bytes.Replace(input, []byte(k), v, -1)
		spew.Dump(input)
	}
	return input
}

func loadFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	checkError(err)
	return file
}

func loadTokenizedDocument(path string) []byte {
	result := loadFile(path)
	return result
}

func outputDetokenizedDocumentToFile(path string, data []byte) {
	err := ioutil.WriteFile(path, data, 0644)
	checkError(err)
}

func outputDetokenizedDocumentToStdout(data []byte) {
	fmt.Print(string(data))
}
