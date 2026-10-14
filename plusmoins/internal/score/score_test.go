package score

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadMissing(t *testing.T) {
	n, err := Load(filepath.Join(t.TempDir(), "record.txt"))
	if err != nil || n != 0 {
		t.Errorf("Load = %d, %v ; attendu 0, nil", n, err)
	}
}

func TestSaveThenLoad(t *testing.T) {
	path := filepath.Join(t.TempDir(), "record.txt")
	if err := Save(path, 7); err != nil {
		t.Fatal(err)
	}
	n, err := Load(path)
	if err != nil || n != 7 {
		t.Errorf("Load = %d, %v ; attendu 7, nil", n, err)
	}
}

func TestLoadCorrupted(t *testing.T) {
	for _, content := range []string{"abc", "", "0", "-4"} {
		path := filepath.Join(t.TempDir(), "record.txt")
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
		if _, err := Load(path); err == nil {
			t.Errorf("Load(%q) : erreur attendue", content)
		}
	}
}

func TestSaveUnwritableDir(t *testing.T) {
	path := filepath.Join(t.TempDir(), "absent", "record.txt")
	if err := Save(path, 3); err == nil {
		t.Error("erreur attendue pour un dossier inexistant")
	}
}
