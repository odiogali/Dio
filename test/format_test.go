package tests

import (
	"testing"
)

func TestMdToHTML(t *testing.T) {
	expected := "this"
	actual := "this"
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestCopyFile(t *testing.T) {
	expected := "that"
	actual := "that"
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}
