# 🔗 GoLink Checker

> Concurrent URL health checker written in Go.

![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8?logo=go&logoColor=white)
![CLI](https://img.shields.io/badge/interface-CLI-black)
![License](https://img.shields.io/badge/license-MIT-green)

GoLink Checker checks multiple websites in parallel and reports their HTTP status, response time and availability.

## ✨ Features

- Concurrent checks using goroutines
- Configurable timeout
- HTTP status and latency
- Reads URLs from arguments or a text file
- Compact terminal output

## 🚀 Run it

```bash
go run . https://github.com https://example.com
```

Or use a file:

```bash
go run . -file urls.txt
```

## ⚙️ Options

```text
-file <path>       Read one URL per line
-timeout <seconds> Request timeout, default 5
```

## 🧠 How it works

Each URL is sent to a worker goroutine. Results travel back through a channel and are printed when they arrive. This keeps the program quick even when one server is slow.

## 🧱 Structure

```text
main.go
go.mod
urls.example.txt
.gitignore
```

## 🛠️ Requirements

Go 1.22 or newer.

## 📄 License

MIT.

## 🆕 Recent changes

### 2026-10-08

- Redirected URLs now show their final destination after the HTTP request completes.

### 2026-10-07

- Duplicate URLs are now removed before checks start, avoiding repeated requests.

### 2026-10-06

- Added validation for timeout and slow-response CLI thresholds.

### 2026-10-05

- Added `-slow <ms>` to flag responses slower than a configurable threshold.

### 2026-10-04

- Returns a nonzero exit code when any URL fails or responds with an HTTP error, making it useful in CI scripts.

### Previous update

- Added a final health summary with healthy/failed counts and average response time.
