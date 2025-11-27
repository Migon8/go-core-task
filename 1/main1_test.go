package main

import (
	"strings"
	"testing"
)

func TestConcatToString(t *testing.T) {
	numDecimal := 42
	numOctal := 052
	numHex := 0x2A
	pi := 3.14
	name := "Golang"
	isActive := true
	complexNum := complex64(1 + 2i)

	got := concatToString(numDecimal, numOctal, numHex, pi, name, isActive, complexNum)

	wantSubstrings := []string{
		"42",
		"3.140000",
		"Golang",
		"true",
		"(1+2i)",
	}

	for _, sub := range wantSubstrings {
		if !strings.Contains(got, sub) {
			t.Errorf("concatToString() = %q, expected to contain %q", got, sub)
		}
	}
}

func TestStringToRunes(t *testing.T) {
	s := "Go Привет"
	runes := stringToRunes(s)

	if len(runes) != 9 {
		t.Errorf("stringToRunes(%q) length = %d, want %d", s, len(runes), 9)
	}

	if string(runes) != s {
		t.Errorf("stringToRunes(%q) => %v (%q), want %q", s, runes, string(runes), s)
	}
}

func TestInsertSaltAndHash_Deterministic(t *testing.T) {
	runes := []rune("test string")
	salt := "go-2024"

	hash1 := insertSaltAndHash(runes, salt)
	hash2 := insertSaltAndHash(runes, salt)

	if hash1 != hash2 {
		t.Errorf("hash not deterministic: %q vs %q", hash1, hash2)
	}
}

func TestInsertSaltAndHash_DifferentSalt(t *testing.T) {
	runes := []rune("test string")
	salt1 := "go-2024"
	salt2 := "another-salt"

	hash1 := insertSaltAndHash(runes, salt1)
	hash2 := insertSaltAndHash(runes, salt2)

	if hash1 == hash2 {
		t.Errorf("expected different hashes for different salts, got %q", hash1)
	}
}
