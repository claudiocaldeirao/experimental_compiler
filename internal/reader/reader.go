package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("error opening the file: %w", err)
	}
	defer file.Close()

	var builder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteRune('\n')
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error while reading the file: %w", err)
	}

	return builder.String(), nil
}
