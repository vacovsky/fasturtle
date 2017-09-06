package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type flagsModel struct {
	bufferChars            *string
	inputPath              *string
	outputPath             *string
	tokensPath             *string
	extract                *bool
	unsafe                 *bool
	dataBag                *string
	dataBagSecret          *string
	bufferCharsLeft        *string
	bufferCharsRight       *string
	assemblyBindingsSource *string
	unquoted               *bool
	version                *bool
}

func flagInit() flagsModel {
	model := flagsModel{}

	model.extract = flag.Bool("extract", false, "If true, enters extract mode.  In extract mode, the output file or stdout becomes a list of the tokens found within the input file.")

	model.bufferChars = flag.String("buffer", "", "Characters used to buffer the keys within the input file.  The default value is an empty string.  Example: \"__mykey__\" (not used if --buffer-left or --buffer-right are provided).")

	model.bufferCharsLeft = flag.String("buffer-left", "", "Characters used to buffer the keys within the input file on the left side of a token key.  The default value is an empty string.  Example: \"{{mykey\" (usually used in conjunction with --buffer-right).")

	model.bufferCharsRight = flag.String("buffer-right", "", "Characters used to buffer the keys within the input file on the right side of a token key.  The default value is an empty string.  Example: \"mykey}}\" (usually used in conjunction with --buffer-left).")

	model.inputPath = flag.String("input", "", "Path the tokenized input file.")

	model.outputPath = flag.String("output", "", "Destination path and file name for the detokenized file.  If not set, detokenized file is printed to stdout.")

	model.tokensPath = flag.String("tokens", "", "Path to the JSON key-value pair set(s) to be used for detokenization of the input file.  For multiple files, separate file paths with a comma (,).  If the same key exists in one or more of the files, then the file furthest to the right takes precendece (overrides) the previous.")

	model.dataBag = flag.String("databag", "", "Name of the Chef data bag containing the tokenized values.  Under the hood, this relies on your environment having a properly configured knife.rb and necessary certs in place to connect to the Chef server.  Alternately, use --tokens to specify a json file.")

	model.dataBagSecret = flag.String("databag-secret", "", "Path to the data bag secret.  Only necessary if you use encrypted data bags.")

	model.unsafe = flag.Bool("unsafe", false, "If true, will not throw error if all tokens are not replaced.  Default is false, and if a token still exists after detokenization, an error will be thrown.")

	model.unquoted = flag.Bool("unquoted", false, "If true, will remove loose typing from replaced values.  This allows for the token to be replaced inside of quotes instead of relying on the detokenization to apple quotes.  For example, unquoted=true: <value=\"__this_works__\">,  vs unquoted=false: <value=__this_works__>")

	model.version = flag.Bool("version", false, "Prints the version number to stdout.")

	model.assemblyBindingsSource = flag.String("assembly-bindings-source", "", "The path to a configuration file containing the correct assembly bindings for the project.  This was added to solve for an issue where bindings set in a base config didn't match thoseset in the token config.  Default is an empty string.")

	flag.Parse()

	if *model.version {
		fmt.Println("Fasturtle", "v"+version)
		os.Exit(0)
	}
	if *model.inputPath == "" {
		fmt.Fprintf(os.Stderr, "Error: At least Input (--input) must be provided.  See --help for details.")
		os.Exit(1)
	}

	if !*model.extract && (*model.dataBag == "" && *model.tokensPath == "") {
		fmt.Fprintf(os.Stderr, `Error: To detokenize, at least --tokens or --databags must have a value.`)
		os.Exit(1)
	}
	if strings.ContainsAny(*model.bufferChars, "$") || strings.ContainsAny(*model.bufferChars, "*") {
		fmt.Fprintf(os.Stderr, `Error: Buffer characters (--buffer) may not contain * nor $.  There are probably 
		other illegal characters I didn't think of, but if you are getting weird errors, maybe try a 
		difference buffer character set.`)
		os.Exit(1)
	}
	if *model.bufferChars == "" && (*model.bufferCharsLeft == "" && *model.bufferCharsRight == "") {
		fmt.Println("Did you forget to set the buffer characters?")
	}
	return model
}
