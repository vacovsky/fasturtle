package main

import (
	"log"
	"strings"
)

func main() {
	// Parse the command line flag values into variables for later use
	args := flagInit()

	// load tokenized document into memory
	input := loadFile(*args.inputPath)

	var output []byte
	if *args.extract {
		tokens := extractTokens(input, *args.bufferChars)

		if strings.HasSuffix(*args.outputPath, ".json") {
			// if the output file ends in json, format it as a json file with { "key": "" }
			output = convertToJSON(tokens, *args.bufferChars)
		} else {
			// just spit out the keys
			for i, t := range tokens {
				if i > 0 {
					output = append(output, []byte("\n")...)
				}
				output = append(output, t...)
			}
		}
	} else {
		// parse tokens from json file
		tokens := mapKeyPairs(*args.tokensPath, *args.bufferChars)

		// store final product for later use
		output = detokenize(input, tokens)
	}

	if *args.outputPath != "" {
		outputToFile(*args.outputPath, output)
	} else {
		outputToStdout(output)
	}
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
