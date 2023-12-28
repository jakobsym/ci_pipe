package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

type exceptionStep struct {
	step // using 'step' struct from step.go
}

func newExceptionStep(name, exe, message, proj string, args []string) exceptionStep {
	s := exceptionStep{}
	s.step = *NewStep(name, exe, message, proj, args)
	return s
}

func (s exceptionStep) execute() (string, error) {
	cmd := exec.Command(s.exe, s.args...)

	// copy cmd to buffer for inspection
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Dir = s.proj

	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step:  s.name,
			msg:   "failed to execute",
			cause: err,
		}
	}

	// verify size of output
	if out.Len() > 0 {
		return "", &stepErr{
			step:  s.name,
			msg:   fmt.Sprintf("Invalid Format: %s", out.String()),
			cause: nil,
		}
	}
	return s.message, nil
}
