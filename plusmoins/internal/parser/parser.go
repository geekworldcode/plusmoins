package parser

import (
	"errors"
	"strings"
	"fmt"
	"strconv"

)

var (
	ErrEmpty = errors.New("saisie vide")
	ErrQuit  = errors.New("le joueur quitte la partie")
)

func ParseGuess(input string, max int) (int, error) {
	s := strings.TrimSpace(input)
	if s == "" {
		return 0, ErrEmpty
	}
	switch strings.ToLower(s) {
	case "q", "Q", "quitter":
		return 0, ErrQuit
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("%q n'est pas un nombre", s)
	}
	if n < 1 || n > max {
		return 0, fmt.Errorf("%d hors limites (1-%d)", n, max)
	}
	return n, nil
}