package parser

import (
	"errors"
	"testing"
)

func TestParseGuess(t *testing.T) {
	cases := []struct {
		input string
		want  int
		ok    bool
	}{
		{"42", 42, true},
		{"  42 \r\n", 42, true},
		{"1", 1, true},
		{"100", 100, true},
		{"", 0, false},
		{"   ", 0, false},
		{"abc", 0, false},
		{"3.5", 0, false},
		{"42 50", 0, false},
		{"0", 0, false},
		{"150", 0, false},
		{"-3", 0, false},
	}
	for _, c := range cases {
		got, err := ParseGuess(c.input, 100)
		if (err == nil) != c.ok {
			t.Errorf("ParseGuess(%q) : err = %v", c.input, err)
		}
		if err == nil && got != c.want {
			t.Errorf("ParseGuess(%q) = %d, attendu %d", c.input, got, c.want)
		}
	}
}

func TestParseGuessEmpty(t *testing.T) {
	if _, err := ParseGuess("  ", 100); !errors.Is(err, ErrEmpty) {
		t.Errorf("err = %v, attendu ErrEmpty", err)
	}
}

func TestParseGuessQuit(t *testing.T) {
	for _, in := range []string{"q", "Q", "quit", "QUITTER", " Quitter "} {
		if _, err := ParseGuess(in, 100); !errors.Is(err, ErrQuit) {
			t.Errorf("ParseGuess(%q) : err = %v, attendu ErrQuit", in, err)
		}
	}
}

func TestParseGuessRange(t *testing.T) {
	_, err := ParseGuess("150", 100)
	var re *RangeError
	if !errors.As(err, &re) {
		t.Fatalf("err = %v, attendu *RangeError", err)
	}
	if re.Value != 150 || re.Max != 100 {
		t.Errorf("RangeError = %+v", re)
	}
}
