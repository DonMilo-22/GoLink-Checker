package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run tools/status-summary.go <checker-output.txt>")
		os.Exit(2)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	ok, failed, lines := 0, 0, 0
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" { continue }
		lines++
		fields := strings.Fields(line)
		for _, field := range fields {
			code, err := strconv.Atoi(strings.Trim(field, "[]()"))
			if err == nil && code >= 100 && code <= 599 {
				if code >= 200 && code < 400 { ok++ } else { failed++ }
				break
			}
		}
	}
	fmt.Printf("Non-empty lines: %d\nStatuses in 2xx/3xx: %d\nOther detected statuses: %d\n", lines, ok, failed)
}
