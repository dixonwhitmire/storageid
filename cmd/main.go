// Package main provides the entrypoint for the storageid utility.
package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Merative-CTO/storageid/internal/storageid"
)

const usage = "storageid [prefix] [max length]"

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("main: invalid arguments. usage: %s", usage)
	}

	// argument values, assign prefix, default length
	prefixArg := os.Args[1]
	lengthArg := "24"

	// assign length if provided
	if len(os.Args) >= 3 {
		lengthArg = os.Args[2]
	}

	idLength, err := strconv.Atoi(lengthArg)
	if err != nil {
		log.Fatalf("main: invalid length argument %s", lengthArg)
	}

	storageId, err := storageid.New(prefixArg, uint(idLength))
	if err != nil {
		log.Fatalf("main: an error occurred %v", err)
	}

	log.Printf("main: generated storageid %s", storageId)
}
