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

func TestSmartSelect(t *testing.T) {
	expected := "smart select"
	actual := "smart selected"
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestChooseFile(t *testing.T) {
	expected := "choose file"
	actual := "choosen file"
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}
