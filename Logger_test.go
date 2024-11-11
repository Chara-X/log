package log

import "os"

func ExampleLogger() {
	var l = New(os.Stdout, "Some prefix ")
	l.Output(2, "Some output")
	// Output:
	// Some prefix 2024/11/11 /usr/local/go/src/testing/run_example.go:63: Some output
}
