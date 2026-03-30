package patterns

import (
	"errors"
	"strings"
)

// validateSize ensures size is a positive integer.
func validateSize(size int) error {
	if size <= 0 {
		return errors.New("size must be a positive integer")
	}
	return nil
}

// RightTriangle generates a right-aligned growing triangle of stars.
// For size N, it returns N rows where row i contains i stars.
func RightTriangle(size int) ([]string, error) {
	if err := validateSize(size); err != nil {
		return nil, err
	}
	out := make([]string, 0, size)
	for i := 1; i <= size; i++ {
		out = append(out, strings.Repeat("*", i))
	}
	return out, nil
}

// CenteredPyramid generates a centered pyramid of stars.
// For size N, it returns N rows; row i has (2*i - 1) stars centered with spaces.
func CenteredPyramid(size int) ([]string, error) {
	if err := validateSize(size); err != nil {
		return nil, err
	}
	out := make([]string, 0, size)
	for i := 1; i <= size; i++ {
		spaces := strings.Repeat(" ", size-i)
		stars := strings.Repeat("*", 2*i-1)
		out = append(out, spaces+stars)
	}
	return out, nil
}

// Diamond generates a centered diamond of stars.
// Size represents the half-height at the widest point.
// For size N, it returns 2*N - 1 rows total.
func Diamond(size int) ([]string, error) {
	if err := validateSize(size); err != nil {
		return nil, err
	}
	out := make([]string, 0, 2*size-1)
	// Top half including middle row.
	for i := 1; i <= size; i++ {
		spaces := strings.Repeat(" ", size-i)
		stars := strings.Repeat("*", 2*i-1)
		out = append(out, spaces+stars)
	}
	// Bottom half.
	for i := size - 1; i >= 1; i-- {
		spaces := strings.Repeat(" ", size-i)
		stars := strings.Repeat("*", 2*i-1)
		out = append(out, spaces+stars)
	}
	return out, nil
}

// HollowSquare generates a hollow square border of stars with the given side length.
// For size N:
//   - If N == 1, it returns a single "*"
//   - If N >= 2, top and bottom rows are full "*", middle rows have "*" borders and inner spaces.
func HollowSquare(size int) ([]string, error) {
	if err := validateSize(size); err != nil {
		return nil, err
	}
	if size == 1 {
		return []string{"*"}, nil
	}

	out := make([]string, 0, size)
	border := strings.Repeat("*", size)
	middle := "*" + strings.Repeat(" ", size-2) + "*"

	out = append(out, border)
	for i := 0; i < size-2; i++ {
		out = append(out, middle)
	}
	out = append(out, border)
	return out, nil
}

// RightPascal generates a right pascal triangle of stars.
// For size N, it returns 2*N - 1 rows: an increasing right triangle then a decreasing one.
func RightPascal(size int) ([]string, error) {
	if err := validateSize(size); err != nil {
		return nil, err
	}
	total := 2*size - 1
	out := make([]string, 0, total)
	// Increasing part.
	for i := 1; i <= size; i++ {
		out = append(out, strings.Repeat("*", i))
	}
	// Decreasing part.
	for i := size - 1; i >= 1; i-- {
		out = append(out, strings.Repeat("*", i))
	}
	return out, nil
}
