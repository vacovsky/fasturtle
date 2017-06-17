package main

import (
	"regexp"
)

// __(.*?)__
// replace __ with buffers

func extractTokens(input []byte, buffer string) [][]byte {
	bufferBuilder := buffer + "(.*?)" + buffer
	rex := regexp.MustCompile(bufferBuilder)
	return rex.FindAll(input, -1)
}
