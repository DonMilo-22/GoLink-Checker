package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run tools/url-dedupe.go <input-file>")
		os.Exit(2)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot open input:", err)
		os.Exit(1)
	}
	defer file.Close()

	seen := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key := strings.ToLower(line)
		if _, exists := seen[key]; exists {
			continue
		}
		seen[key] = struct{}{}
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Cannot read input:", err)
		os.Exit(1)
	}
}
