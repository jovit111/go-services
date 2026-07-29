# Go Microservices

Ten self-contained HTTP services built with Go.

## Services

| Service | Port | Endpoint | Description |
|---------|------|----------|-------------|
| hello-service | 8080 | `/?name=...` | Greeting API |
| echo-service | 8081 | `/` | Echo request details |
| health-service | 8082 | `/health`, `/ready` | Health and readiness checks |
| time-service | 8083 | `/?format=...` | Current time, RFC3339/unix formats |
| generate-service | 8084 | `/uuid`, `/api-key` | UUID/API key generation |
| reverse-service | 8085 | `/?text=...` | Reverse a string |
| json-service | 8086 | `/?text=...` | Text hashing, casing, length |
| math-service | 8087 | `/add`, `/multiply`, `/factorial` | Basic math operations |
| weather-service | 8088 | `/?location=...` | Mock weather API |
| logger-service | 8089 | `/log` | Request logger/stdout output |

## Run

```bash
go run ./cmd/hello-service
```

Or build:
```bash
go build -o bin/ ./cmd/...
```
