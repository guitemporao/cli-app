package todo

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

func GetInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if text == "" {
		return "", errors.New("empty task")
	}

	return text, nil
}
