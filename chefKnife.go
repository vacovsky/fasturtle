package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"

	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// lets you substitute data bags for token files!  So cool!
// knife data bag show {data_bag} {bag_entry.json} -F json

// TODO: build support for encryptoed data bags
func collectEncrytpedDataBagJSON(bag, blob, secret string) []byte {
	spew.Dump(bag, blob, secret)
	cmd := exec.Command("knife", "data", "bag", "show", "--secret-file", secret, bag, blob, "-F", "json")
	fmt.Println("knife", "data", "bag", "show", "--secret-file", secret, bag, blob, "-F", "json")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.Bytes()
}

func listDataBagEntries(bag string) []string {
	cmd := exec.Command("knife", "data", "bag", "show", bag)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(out.String(), "\n")
}

func collectDataBagJSON(bag, blob string) []byte {
	cmd := exec.Command("knife", "data", "bag", "show", bag, blob, "-F", "json")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.Bytes()
}
