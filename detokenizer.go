package main

import (
	"encoding/json"
	"log"
	"os"

	"fmt"
)

func main() {
	// Parse the command line flag values into variables for later use
	args := flagInit()

	// convert JSON file into a map[string]string for later use
	tokens := mapKeyPairs(*args.tokensPath)

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

func mapKeyPairs(path string) map[string]*json.RawMessage {
	tokenMap := map[string]*json.RawMessage{}

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Panic(err)
	}

	dbytes := []byte{}
	_, err = file.Read(dbytes)
	err = json.Unmarshal(dbytes, &tokenMap)

	// spew.Dump(dbytes, tokenMap)

	return tokenMap
}

func detokenize(input string, tokenMap map[string]*json.RawMessage) string {
	result := ""

	for k, v := range tokenMap {
		fmt.Println(k, v)
	}

	return result
}

func loadTokenizedDocument(path string) string {
	return ""
}

func outputDetokenizedDocumentToFile(path string, data string) {

}

func outputDetokenizedDocumentToStdout(data string) {

}
