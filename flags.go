package main

import (
	"flag"
)

type flagsModel struct {
	bufferChars *string
	inputPath   *string
	outputPath  *string
	tokensPath  *string
}

func flagInit() flagsModel {
	model := flagsModel{}
	model.bufferChars = flag.String("buffer", "__", "Characters used to buffer the keys within the input file.  Example: __mykey__.")
	model.inputPath = flag.String("input", "", "Path the tokenized input file.")
	model.outputPath = flag.String("output", "", "Destination path and file name for the detokenized file.  If not set, detokenized file is printed to stdout.")
	model.tokensPath = flag.String("tokens", "", "Path to the JSON key-value pair set to be used for detokenization of the input file.")

	flag.Parse()
	return model
}
