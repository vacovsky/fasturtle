package main

import (
	"regexp"
)

func fixAssemblyBindings(output []byte, inputPath string) []byte {
	rex := "<assemblies>(?s)(.*)</assemblies>"
	return extractNodes(output, inputPath, rex)
}

func extractNodes(output []byte, inputPath, regexstring string) []byte {
	// find allemblies block in XML and store it in a variable
	realBindings := findXMLBlock(inputPath, regexstring)

	// replace assemblies block in detokenized output with the assemblies block from the default web.config.
	// We are making the assumption that the destination output file is the source of the updated assemblly bindings.
	return replaceXMLBlock(regexstring, output, realBindings)
}

func findXMLBlock(input, rexinp string) []byte {
	rex := regexp.MustCompile(rexinp)
	return rex.FindAll(loadFile(input), -1)[0]
}

func replaceXMLBlock(rexinp string, source, vals []byte) []byte {
	rex := regexp.MustCompile(rexinp)
	return rex.ReplaceAll(source, vals)
}
