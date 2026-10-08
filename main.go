package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Result struct {
	URL string
	Status int
	Duration time.Duration
	Err error
}

func check(client *http.Client, url string, out chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	start := time.Now()
	resp, err := client.Get(url)
	if err != nil { out <- Result{URL:url, Err:err, Duration:time.Since(start)}; return }
	defer resp.Body.Close()
	out <- Result{URL:url, Status:resp.StatusCode, Duration:time.Since(start)}
}

func fromFile(path string) ([]string, error) {
	f, err := os.Open(path); if err != nil { return nil, err }; defer f.Close()
	var urls []string
	s := bufio.NewScanner(f)
	for s.Scan() { if v := strings.TrimSpace(s.Text()); v != "" && !strings.HasPrefix(v, "#") { urls = append(urls, v) } }
	return urls, s.Err()
}

func main() {
	file := flag.String("file", "", "file containing one URL per line")
	timeout := flag.Int("timeout", 5, "request timeout in seconds")
	slow := flag.Int("slow", 1000, "mark responses slower than this many milliseconds")
	flag.Parse()
	if *timeout <= 0 || *slow < 0 {
		fmt.Fprintln(os.Stderr, "timeout must be positive and slow must not be negative")
		os.Exit(2)
	}
	urls := flag.Args()
	if *file != "" {
		fu, err := fromFile(*file); if err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(1) }
		urls = append(urls, fu...)
	}
	if len(urls) == 0 { fmt.Println("Usage: go run . [-file urls.txt] [-timeout 5] <url...>"); return }
	seen := map[string]bool{}
	unique := make([]string, 0, len(urls))
	for _, u := range urls {
		if !seen[u] { seen[u] = true; unique = append(unique, u) }
	}
	urls = unique

	client := &http.Client{Timeout: time.Duration(*timeout)*time.Second}
	out := make(chan Result)
	var wg sync.WaitGroup
	for _, u := range urls { wg.Add(1); go check(client, u, out, &wg) }
	go func(){ wg.Wait(); close(out) }()

	success, failed, slowCount := 0, 0, 0
	var totalDuration time.Duration
	for r := range out {
		totalDuration += r.Duration
		if r.Err != nil {
			failed++
			fmt.Printf("✗ %-35s ERROR (%s)\n", r.URL, r.Duration.Round(time.Millisecond))
			continue
		}
		if r.Status >= 200 && r.Status < 400 { success++ } else { failed++ }
		label := ""
		if r.Duration > time.Duration(*slow)*time.Millisecond { slowCount++; label = "  SLOW" }
		fmt.Printf("✓ %-35s %3d  %s%s\n", r.URL, r.Status, r.Duration.Round(time.Millisecond), label)
	}
	average := totalDuration / time.Duration(len(urls))
	fmt.Printf("\nSummary: %d healthy | %d failed | %d slow | avg %s\n", success, failed, slowCount, average.Round(time.Millisecond))
	if failed > 0 { os.Exit(1) }
}
