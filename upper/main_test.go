package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// helper to run main and capture STDOUT
func runMain() string {

	// mock STDOUT
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w // main method will write to my pipe

	// ACT
	main()
	w.Close()
	os.Stdout = originalStdout

	// read from captured STDOUT
	var output bytes.Buffer
	io.Copy(&output, r)
	return output.String()
}

func TestUpperWithMultipleArgs_AreSpaceDelimitedWithNewLine(t *testing.T) {
	os.Args = []string{"upper", "hello", "world"}
	output := runMain()
	assert.Equal(t, "HELLO WORLD\n", output)
}

func TestUpperWithArg_IncludesNewLine(t *testing.T) {
	os.Args = []string{"upper", "hello"}
	captured_output := runMain()
	assert.Equal(t, "HELLO\n", captured_output)
}

func TestUpperWithMultiLineStdin_PreservesLinesAndIncludesNewLine(t *testing.T) {
	// SETUP
	os.Args = []string{"upper"} // no args
	// mock STDIN
	originalStdin := os.Stdin
	stdin_read, stdin_write, _ := os.Pipe()
	os.Stdin = stdin_read // main will read STDIN from this pipe
	// write and close STDIN
	stdin_write.WriteString("hello world\nfoo the Bar")
	stdin_write.Close()

	// ACT
	capturedOutput := runMain()
	os.Stdin = originalStdin

	assert.Equal(t, "HELLO WORLD\nFOO THE BAR\n", capturedOutput)
}
