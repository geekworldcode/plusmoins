package parser

import (
	"errors"
	"testing"
)

func TestParseGuess(t *testing.T) {
	cases := []struct {
		input string
		ok    bool
	}{
		{"42", true},       // nombre accepté
		{"  42 ", true},    // les espaces sont ignorés
		{" 7 \r\n", true},  // fin de ligne Windows
		{"", false},        // saisie vide
		{"abc", false},     // pas un nombre
		{"3.5", false},     // pas un nombre entier
		{"42 50", false},   // deux nombres d'un coup
		{"150", false},     // hors limites
		{"0", false},       // hors limites
		{"-3", false},      // hors limites
		{"q", false},       // quitter
		{"Q", false},       // quitter, autre casse
		{"quitter", false}, // quitter
		{"QUITTER", false}, // quitter, autre casse
	}
	for _, c := range cases {
		_, err := ParseGuess(c.input, 100)
		if (err == nil) != c.ok {
			t.Errorf("ParseGuess(%q) : err = %v", c.input, err)
		}
	}
}

func TestParseGuessErrEmpty(t *testing.T) {
	if _, err := ParseGuess("   ", 100); !errors.Is(err, ErrEmpty) {
		t.Errorf("ParseGuess(\"   \") : err = %v, attendu ErrEmpty", err)
	}
}

func TestParseGuessErrQuit(t *testing.T) {
	for _, in := range []string{"q", "Q", "quitter", "QUITTER", " Quitter "} {
		if _, err := ParseGuess(in, 100); !errors.Is(err, ErrQuit) {
			t.Errorf("ParseGuess(%q) : err = %v, attendu ErrQuit", in, err)
		}
	}
}
