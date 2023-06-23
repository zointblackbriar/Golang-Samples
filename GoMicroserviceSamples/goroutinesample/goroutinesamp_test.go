package goroutinesample

import (
	"bytes"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestF(t *testing.T) {
	// Redirect standard output to a buffer for testing
	var buf bytes.Buffer
	log.SetOutput(&buf)

	expected := []string{
		"direct: 0",
		"direct: 1",
		"direct: 2",
	}

	String_func("direct")

	output := buf.String()
	lines := splitLines(output)

	if len(lines) != len(expected) {
		t.Errorf("Expected %d lines, but got %d", len(expected), len(lines))
	}

	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("Expected '%s', but got '%s'", expected[i], line)
		}
	}
}

func TestMain(t *testing.T) {
	// Redirect standard output to a buffer for testing
	var buf bytes.Buffer
	log.SetOutput(&buf)

	expected := []string{
		"direct: 0",
		"direct : 1",
		"direct : 2",
		"goroutine: 0",
		"going",
		"goroutine: 1",
		"goroutine: 2",
		"done",
	}

	// Run the main function in a separate goroutine
	go Goroutinemain()

	time.Sleep(time.Second) // Wait for goroutines to finish

	output := buf.String()
	lines := splitLines(output)

	if len(lines) != len(expected) {
		t.Errorf("Expected %d lines, but got %d", len(expected), len(lines))
	}

	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("Expected '%s', but got '%s'", expected[i], line)
		}
	}
}

func splitLines(sample_string string) []string {
	lines := []string{}
	line := ""

	for _, c := range sample_string {
		if c == '\n' {
			lines = append(lines, line)
			line = ""
		} else {
			line += string(c)
		}
	}

	if line != "" {
		lines = append(lines, line)
	}

	return lines
}

func TestSplitLines1(t *testing.T) {
	input := "Hello\nWorld\nTest"
	expected := []string{"Hello", "World", "Test"}

	result := splitLines(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected result. Expected: %v, Got: %v", expected, result)
	}
}
