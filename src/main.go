package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	tokens := stringToTokens(test1)

	elapsed := time.Since(startTime)

	printTokens(tokens)

	fmt.Printf("\nExecution time: %s\n", elapsed)
}
