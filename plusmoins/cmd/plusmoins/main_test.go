package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestPieges(t *testing.T) {
	dir := t.TempDir()
	bin := filepath.Join(dir, "plusmoins")
	if out, err := exec.Command("go", "build", "-o", bin, ".").CombinedOutput(); err != nil {
		t.Fatalf("compilation : %v\n%s", err, out)
	}

	input, err := os.Open(filepath.Join("..", "..", "testdata", "pieges.txt"))
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()
	want, err := os.ReadFile(filepath.Join("..", "..", "testdata", "attendu.txt"))
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command(bin, "-secret", "42", "-record", filepath.Join(dir, "record.txt"))
	cmd.Stdin = input
	var got bytes.Buffer
	cmd.Stdout = &got
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got.Bytes(), want) {
		t.Errorf("sortie différente de attendu.txt\n--- obtenu\n%s\n--- attendu\n%s", got.Bytes(), want)
	}
}
