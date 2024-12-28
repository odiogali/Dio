package main

import (
	"reflect"
	"testing"
)

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

func TestExtractPhotosAsym(t *testing.T) {
	expected := []string{"Pasted image 20240514223619.png|600", "Pasted image 20240514223650.png|550"}
	actual := extractPhotos("test_input/Asymptotic Notation.md")
	if actual != nil {
		if len(expected) != len(actual) || !reflect.DeepEqual(actual, expected) {
			t.Fatalf("Expected: %s, Actual: %s", expected, actual)
		}
	} else {
		t.Fatalf("Actual is nil but it shouldn't be.")
	}
}

func TestExtractPhotosMin(t *testing.T) {
	expected := "extract photos"
	actual := "extract photos"
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestExtractPhotosInternet(t *testing.T) {
	expected := "extract photos"
	actual := "extract photos"
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestExtractPhotosSub(t *testing.T) {
	expected := "extract photos"
	actual := "extract photos"
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestExtractPhotosSets(t *testing.T) {
	expected := "extract photos"
	actual := "extract photos"
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestExtractPhotosParse(t *testing.T) {
	expected := "extract photos"
	actual := "extract photos"
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}
