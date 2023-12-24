package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
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

// Execute 'go buisld' on provided `proj“
func run(proj string, out io.Writer) error {
	if proj == "" {
		return fmt.Errorf("Project directory required: %w", ErrValidation)
	}

	args := []string{"build", ".", "errors"} // prevents creating a file, that would need to be cleaned
	cmd := exec.Command("go", args...)       // `args...` expands our slice into list of strings

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("'go build' failed: %s", err)
	}

	_, err := fmt.Fprintln(out, "Go Build: SUCCESS")

	return err
}
