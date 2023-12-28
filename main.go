package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// Parse cl flags, before calling 'run()'
func main() {
	proj := flag.String("p", "", "Project Directory")
	flag.Parse()

	if err := run(*proj, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Execute 'go build' on provided `proj“
func run(proj string, out io.Writer) error {
	// wrapping error
	if proj == "" {
		return fmt.Errorf("Project directory required: %w", ErrValidation)
	}
	pipeline := make([]step, 1) // make a slice w/ cap. of 1
	pipeline[0] = *NewStep(
		"go build",
		"go",
		"go build: SUCCESS",
		proj,
		[]string{"build", ".", "errors"},
	)
	// Loop through pipeline, printing the
	for _, s := range pipeline {
		msg, err := s.execute()
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(out, msg)
		if err != nil {
			return err
		}
	}
	return nil
}
