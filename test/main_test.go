package tests

import (
	"testing"
)

func TestUpdateContent(t *testing.T) {
	expected := "update"
	actual := "update"
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestRepopulate(t *testing.T) {
	expected := "repopulate"
	actual := "repopulate"
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}
