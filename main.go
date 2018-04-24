package main

import (
	"flag"
	"fmt"
	"github.com/cbrgm/mcdc/matrix"
)

func main() {
	filePtr := flag.String("file", "./table.txt", "The file containing the matrix")
	flag.Parse()
	grid, err := matrix.FromFile(*filePtr)
	if err != nil {
		panic(err)
	}

	for i, vector := range matrix.GetRelevant(grid) {
		fmt.Printf("Testcase #%v : %v\n", i, vector)
	}
}
