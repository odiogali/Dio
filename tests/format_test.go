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
	expected := "copy file"
	actual := "copy file"
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestCopyImages(t *testing.T) {
	expected := "copy images"
	actual := "copy images"
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestExtractPhotos(t *testing.T) {
	expected := "extract photos"
	actual := "extract photos"
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}
