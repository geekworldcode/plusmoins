package ui

import (
		"bufio"
		"fmt"
		"io"
		"os"
	)

var in = bufio.NewScanner(os.Stdin)

func Ask() (string, error) {
	fmt.Print("Ton nombre : ")
	if in.Scan() {
		return in.Text(), nil
	}
	if err := in.Err(); err != nil {
		return "", err
	}
	return "", io.EOF
}