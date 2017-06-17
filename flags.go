package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type flagsModel struct {
	bufferChars *string
	inputPath   *string
	outputPath  *string
	tokensPath  *string
	extract     *bool
}

func flagInit() flagsModel {
	model := flagsModel{}
	model.bufferChars = flag.String("buffer", "__", "Characters used to buffer the keys within the input file.  Example: __mykey__.")
	model.inputPath = flag.String("input", "", "Required.  Path the tokenized input file.")
	model.outputPath = flag.String("output", "", "Destination path and file name for the detokenized file.  If not set, detokenized file is printed to stdout.")
	model.tokensPath = flag.String("tokens", "", "Path to the JSON key-value pair set to be used for detokenization of the input file.")
	model.extract = flag.Bool("extract", false, "If true, enters extract mode.  In extract mode, the output file or stdout becomes a list of the tokens found within the input file.")

	flag.Parse()
	if *model.inputPath == "" {
		fmt.Println("Error: Input (--input) cannot be empty.  Please provide a file path.  See --help for details.")
		os.Exit(1)
	}

	if strings.ContainsAny(*model.bufferChars, "$") || strings.ContainsAny(*model.bufferChars, "*") {
		fmt.Println(`Error: Buffer characters (--buffer) may not contain * nor $.  There are probably 
		other illegal characters I didn't think of, but if you are getting weird errors, maybe try a 
		difference buffer character set.`)
		os.Exit(1)
	}

	return model
}
