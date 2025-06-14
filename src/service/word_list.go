package service

import (
	"bufio"
	"os"
	"strings"
)

func ReadWordFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, strings.TrimSpace(scanner.Text()))
	}
	return words, scanner.Err()
}
