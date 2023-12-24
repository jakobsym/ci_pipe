package main

import (
	"bytes"
	"errors"
	"testing"
)

var testCases = []struct {
	name   string
	proj   string
	out    string
	expErr error
}{
	{name: "success", proj: "/testData/tool/add.go",
		out:    "Go Build: SUCCESS\n",
		expErr: nil},
	{name: "fail", proj: "/testData/toolErr/add.go",
		out:    "",
		expErr: &stepErr{step: "go build"}},
}

/* Using Table-Driven Testing */
func testRun(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer
			err := run(tc.proj, &out)

			if tc.expErr != nil {
				if err == nil {
					t.Errorf("Expected error: %q\nGot 'nil' instead.", tc.expErr)
					return
				}
				if !errors.Is(err, tc.expErr) {
					t.Errorf("Expected error: %q\nGot %q instead.", tc.expErr, err)
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error: %q", err)
			}
			if out.String() != tc.out {
				t.Errorf("Expected output: %q\nGot %q isntead.", tc.out, out.String())
			}
		})
	}
}
