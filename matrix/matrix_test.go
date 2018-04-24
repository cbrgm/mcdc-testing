package matrix

import (
	"testing"
)

func TestGetRelevant(t *testing.T) {
	input := [][]string{
		{"1", "0", "0"},
		{"1", "1", "0"},
		{"0", "1", "1"}, // not neighbour
		{"1", "0", "1"},
	}

	got, err := createMatrix(input)
	if err != nil {
		t.Error(err)
	}

	want := 3
	neighbours := GetRelevant(got)
	if len(neighbours) != want {
		t.Errorf("Number of neighbours expected is invalid, want %v, got %v", want, len(neighbours))
	}
}

func TestIsRelevant(t *testing.T) {
	n1 := Vector{
		Values: []bool{true, false, false},
	}
	n2 := Vector{
		Values: []bool{true, false, true},
	}
	if want, got := false, isRelevant(n1, n2); got != want {
		t.Errorf("Want data %v, got %v", want, got)
	}

	n2 = Vector{
		Values: []bool{true, true, true},
	}
	if want, got := true, isRelevant(n1, n2); got != want {
		t.Errorf("Want data %v, got %v", want, got)
	}
}

func TestIsTruthValueDifferent(t *testing.T) {
	n1 := Vector{
		Values: []bool{true, false, false},
	}
	n2 := Vector{
		Values: []bool{false, true, false},
	}
	if want, got := false, isTruthValueDifferent(n1, n2); got != want {
		t.Errorf("Want data %v, got %v", want, got)
	}

	n2 = Vector{
		Values: []bool{false, true, true},
	}
	if want, got := true, isTruthValueDifferent(n1, n2); got != want {
		t.Errorf("Want data %v, got %v", want, got)
	}
}

func TestIsNeighbour(t *testing.T) {
	n1 := Vector{
		Values: []bool{true, false, false},
	}
	n2 := Vector{
		Values: []bool{false, true, false},
	}
	if want, got := false, isNeighbour(n1, n2); got != want {
		t.Errorf("Want data %v, got %v", want, got)
	}

	n2 = Vector{
		Values: []bool{true, true, false},
	}
	if want, got := true, isNeighbour(n1, n2); got != want {
		t.Errorf("Want data %v, got %v", want, got)
	}

}

func TestGetNeighbours(t *testing.T) {
	input := [][]string{
		{"1", "0", "0"},
		{"1", "1", "0"}, // is neighbour
		{"0", "0", "1"},
		{"1", "0", "1"}, // is neighbour
	}

	got, err := createMatrix(input)
	if err != nil {
		t.Error(err)
	}

	neighbours := got.GetNeighbours(got.Vectors[0])
	want := 2
	if len(neighbours) != want {
		t.Errorf("Number of neighbours expected is invalid, want %v, got %v", want, len(neighbours))
	}
}

func TestCreateMatrix(t *testing.T) {
	input := [][]string{
		{"1", "0", "0"},
		{"0", "1", "0"},
		{"0", "0", "1"},
	}

	want := [3][3]bool{
		{true, false, false},
		{false, true, false},
		{false, false, true},
	}

	got, err := createMatrix(input)
	if err != nil {
		t.Error(err)
	}

	for x := 0; x < len(got.Vectors); x++ {
		for y := 0; y < len(got.Vectors[x].Values); y++ {
			if got.Vectors[x].Values[y] != want[x][y] {
				t.Errorf("Want data %v, got %v", got.Vectors[x].Values[y], want[x][y])
			}
		}
	}
}
