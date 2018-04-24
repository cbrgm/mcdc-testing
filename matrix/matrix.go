package matrix

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Matrix struct {
	Vectors []Vector
}

type Vector struct {
	Values []bool
}

func GetRelevant(matrix *Matrix) []Vector {
	result := make(map[string]Vector)
	for _, vector := range matrix.Vectors {
		for _, toCompare := range matrix.Vectors {
			if isRelevant(vector, toCompare) {
				// Use map as a data structure
				// instead of a slice for removing dublicates
				result[fmt.Sprintf("%v", toCompare)] = toCompare
			}
		}
	}
	return asList(&result)
}

// Checks wether the two vectors are relevant to be tested or not
func isRelevant(n1 Vector, n2 Vector) bool {
	return isNeighbour(n1, n2) && isTruthValueDifferent(n1, n2)
}

// Checks wether the truth value is different or not
func isTruthValueDifferent(n1 Vector, n2 Vector) bool {
	return n1.Values[len(n1.Values)-1] != n2.Values[len(n2.Values)-1]
}

// Returns true when two vectors are neighbours, false if not
func isNeighbour(n1 Vector, n2 Vector) bool {
	differences := 0

	// Iterate over atomic conditions, NOT result (hence -1).
	for i := 0; i < len(n1.Values)-1; i++ {
		if n1.Values[i] != n2.Values[i] {
			differences++
		}
	}
	return differences == 1
}

// Returns true when two vectors are neighbours, false if not
func (m *Matrix) GetNeighbours(n1 Vector) []Vector {
	var neighbours []Vector

	for _, toCompare := range m.Vectors {
		if isNeighbour(n1, toCompare) {
			neighbours = append(neighbours, toCompare)
		}
	}

	return neighbours
}

// Creates a new Matrix from file
func FromFile(path string) (*Matrix, error) {

	file, err := os.Open(path)
	if err != nil {
		return new(Matrix), err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines [][]string
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ","))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if result, err := createMatrix(lines); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// Converts a string matrix to a boolean based matrix
func createMatrix(lines [][]string) (*Matrix, error) {
	var result Matrix

	// Iterate over each line in file
	for x := 0; x < len(lines); x++ {

		// Create a new vector
		// and convert strings to bool
		var vec Vector
		for y := 0; y < len(lines[x]); y++ {
			bool, err := strconv.ParseBool(lines[x][y])
			if err != nil {
				return nil, err
			}
			vec.Values = append(vec.Values, bool)
		}
		result.Vectors = append(result.Vectors, vec)
	}

	return &result, nil
}

// Helper function to convert maps to slice
func asList(set *map[string]Vector) []Vector {
	var result []Vector
	for _, value := range *set {
		result = append(result, value)
	}
	return result
}
