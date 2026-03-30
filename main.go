package main

import (
	"flag"
	"fmt"
	"os"

	"printing_patterns/internal/patterns"
)

func generators() map[string]func(int) ([]string, error) {
	return map[string]func(int) ([]string, error){
		"right-triangle":   patterns.RightTriangle,
		"centered-pyramid": patterns.CenteredPyramid,
		"diamond":          patterns.Diamond,
		"hollow-square":    patterns.HollowSquare,
		"right-pascal":     patterns.RightPascal,
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s -pattern PATTERN -size N\n\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "Supported patterns:")
	fmt.Fprintln(os.Stderr, "  - right-triangle")
	fmt.Fprintln(os.Stderr, "  - centered-pyramid")
	fmt.Fprintln(os.Stderr, "  - diamond")
	fmt.Fprintln(os.Stderr, "  - hollow-square")
	fmt.Fprintln(os.Stderr, "  - right-pascal")
	fmt.Fprintln(os.Stderr, "\nExamples:")
	fmt.Fprintln(os.Stderr, "  go run . -pattern diamond -size 5")
}

func printLines(lines []string) {
	printLines(lines)
}

func main() {
	patternFlag := flag.String("pattern", "", "pattern name: right-triangle, centered-pyramid, diamond, hollow-square, right-pascal")
	sizeFlag := flag.Int("size", 0, "positive integer size parameter")

	flag.Usage = usage
	flag.Parse()

	genMap := generators()

	if *patternFlag == "" {
		fmt.Fprintln(os.Stderr, "error: -pattern is required")
		usage()
		os.Exit(2)
	}

	gen, ok := genMap[*patternFlag]
	if !ok {
		fmt.Fprintf(os.Stderr, "error: unknown pattern %q\n", *patternFlag)
		usage()
		os.Exit(2)
	}

	if *sizeFlag <= 0 {
		fmt.Fprintln(os.Stderr, "error: -size must be a positive integer")
		usage()
		os.Exit(2)
	}

	lines, err := gen(*sizeFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
