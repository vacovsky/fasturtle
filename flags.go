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
	dataBag     *string
}

func flagInit() flagsModel {
	model := flagsModel{}
	model.bufferChars = flag.String("buffer", "", "Characters used to buffer the keys within the input file.  The default value is an empty string.  Example: __mykey__.")
	model.inputPath = flag.String("input", "", "Path the tokenized input file.")
	model.outputPath = flag.String("output", "", "Destination path and file name for the detokenized file.  If not set, detokenized file is printed to stdout.")
	model.tokensPath = flag.String("tokens", "", "Path to the JSON key-value pair set(s) to be used for detokenization of the input file.  For multiple files, separate file paths with a comma (,).")
	model.extract = flag.Bool("extract", false, "If true, enters extract mode.  In extract mode, the output file or stdout becomes a list of the tokens found within the input file.")
	model.dataBag = flag.String("databag", "", "Name of the Chef data bag containing the tokenized values.  Under the hood, this relies on your environment having a properly configured knife.rb and necessary certs in place to connect to the Chef server.  Alternately, use --tokens to specify a json file.")
	flag.Parse()
	if *model.inputPath == "" {
		fmt.Println("Error: At least Input (--input) or Data Bag (--databag) must be provided.  Please provide at least one of them.  See --help for details.")
		os.Exit(1)
	}

	if !*model.extract && (*model.dataBag == "" && *model.tokensPath == "") {
		fmt.Println(`Error: To detokenize, at least --tokens or --databags must have a value.`)
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
