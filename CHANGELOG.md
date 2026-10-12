# Historial de cambios

Entradas originales conservadas en inglés. Las fechas corresponden al historial existente; la última entrada no tenía fecha registrada.

### 2026-10-09

- The final summary now shows the fastest and slowest successful response times.

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

### Actualización anterior (sin fecha registrada)

- Added a final health summary with healthy/failed counts and average response time.
