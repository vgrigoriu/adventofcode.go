package main

import "testing"

func TestHasLetterTwice(t *testing.T) {
	result := hasLetterTwice("abcdef")
	if result != false {
		t.Error("Expected false, got", result)
	}

	result = hasLetterTwice("bababc")
	if result != true {
		t.Error("Expected true, got", result)
	}

	result = hasLetterTwice("abbcde")
	if result != true {
		t.Error("Expected true, got", result)
	}

	result = hasLetterTwice("abcccd")
	if result != false {
		t.Error("Expected false, got", result)
	}
}

func TestHasLetterThrice(t *testing.T) {
	result := hasLetterThrice("abcdef")
	if result != false {
		t.Error("Expected false, got", result)
	}

	result = hasLetterThrice("bababc")
	if result != true {
		t.Error("Expected true, got", result)
	}

	result = hasLetterThrice("abbcde")
	if result != false {
		t.Error("Expected false, got", result)
	}

	result = hasLetterThrice("abcccd")
	if result != true {
		t.Error("Expected true, got", result)
	}
}
