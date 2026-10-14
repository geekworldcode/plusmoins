package score

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

func Load(path string) (int, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return 0, nil // premier lancement
	}
	if err != nil {
		return 0, fmt.Errorf("lecture de %s : %w", path, err)
	}
	n, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return 0, fmt.Errorf("%s est corrompu : %w", path, err)
	}
	return n, nil
}

func Save(path string, attempts int) error {
	data := []byte(strconv.Itoa(attempts) + "\n")
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("écriture de %s : %w", path, err)
	}
	return nil
}
