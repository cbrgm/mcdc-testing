package main

import (
	"flag"
	"fmt"
	"github.com/cbrgm/mcdc/matrix"
)

func main() {

	// parse cli commands
	filePtr := flag.String("file", "./table.txt", "The file containing the matrix")
	flag.Parse()

	// create matrix struct from file
	grid, err := matrix.FromFile(*filePtr)
	if err != nil {
		panic(err)
	}

	// print the output
	for i, vector := range matrix.GetRelevant(grid) {
		fmt.Printf("Testcase #%v : %v\n", i, vector)
	}
}
