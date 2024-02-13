package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/markalanboyd/gokenizer/gokenizer"
)

func main() {
	prettyPrintFlag := flag.Bool("pp", false, "pretty-prints token output")
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: gokenizer <source.lua> [destination.json] [flags]")
		os.Exit(1)
	}

	sourceFilePath := args[0]
	var destFilePath string

	fileContents, err := readFileAsString(sourceFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(args) >= 2 && !strings.HasPrefix(args[1], "-") {
		destFilePath = args[1]
		err = validateExtension(destFilePath, ".json")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		sourceBase := strings.TrimSuffix(sourceFilePath, ".lua")
		destFilePath = sourceBase + "_tokens.json"
	}

	fmt.Printf("\nOutput will be saved to: %s\n", destFilePath)
	fmt.Printf("\n-- Tokens --\n")

	startTime := time.Now()
	tokens := gokenizer.Tokenize(fileContents)
	tokenizingTime := time.Since(startTime)

	var tokensJSON []byte
	if *prettyPrintFlag {
		tokensJSON, err = json.MarshalIndent(tokens, "", "    ")
	} else {
		tokensJSON, err = json.Marshal(tokens)
	}

	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}

	err = os.WriteFile(destFilePath, tokensJSON, 0o644)
	if err != nil {
		log.Fatalf("Error writing JSON to file: %v", err)
	}

	totalTime := time.Since(startTime)
	writingTime := totalTime - tokenizingTime

	gokenizer.PrintTokens(tokens)

	fmt.Println("\n-- Execution Time --")
	fmt.Printf("Tokenizing........%s\n", tokenizingTime)
	fmt.Printf("Writing to file...%s\n", writingTime)
	fmt.Printf("Total.............%s\n", totalTime)
}
