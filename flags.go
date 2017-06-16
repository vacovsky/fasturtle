package main

import (
	"flag"
)

func flagInit() {
	flag.String("input", "", "Path the tokenized input file.")
	flag.String("output", "", "Destination path and file name for the detokenized file.  If not set, detokenized file is printed to stdout.")
	flag.String("tokens", "", "Path to the JSON key-value pair set to be used for detokenization of the input file.")
}
