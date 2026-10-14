package score

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Load(path string) (int, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return 0, nil
	}
	if err != nil {
		return 0, fmt.Errorf("lecture de %s : %w", path, err)
	}
	n, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return 0, fmt.Errorf("%s est corrompu : %w", path, err)
	}
	if n < 1 {
		return 0, fmt.Errorf("%s est corrompu : record %d invalide", path, n)
	}
	return n, nil
}

func Save(path string, attempts int) error {
	tmp, err := os.CreateTemp(filepath.Dir(path), ".record-*")
	if err != nil {
		return fmt.Errorf("écriture de %s : %w", path, err)
	}
	defer os.Remove(tmp.Name())

	if _, err := fmt.Fprintln(tmp, attempts); err != nil {
		tmp.Close()
		return fmt.Errorf("écriture de %s : %w", path, err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("écriture de %s : %w", path, err)
	}
	if err := os.Chmod(tmp.Name(), 0644); err != nil {
		return fmt.Errorf("écriture de %s : %w", path, err)
	}
	if err := os.Rename(tmp.Name(), path); err != nil {
		return fmt.Errorf("écriture de %s : %w", path, err)
	}
	return nil
}
