# Go Microservices

Twenty self-contained HTTP services built with Go.

## Services

### Core Services (1-10)
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

### Extra Services (11-20)
| Service | Port | Endpoint | Description |
|---------|------|----------|-------------|
| auth-service | 9001 | `/login`, `/validate` | JWT auth mock |
| cache-service | 9002 | `/set`, `/get` | In-memory key/value cache |
| config-service | 9003 | `/config`, `/configs` | App config store |
| database-service | 9004 | `/users` | Simple REST database |
| email-service | 9005 | `/send`, `/inbox`, `/template` | Mock email service |
| file-service | 9006 | `/upload`, `/download`, `/list` | Simple file storage |
| proxy-service | 9007 | `/forward`, `/routes`, `/headers` | HTTP proxy mock |
| queue-service | 9008 | `/enqueue`, `/dequeue` | Priority queue |
| rate-service | 9009 | `/check` | Token bucket rate limiter |
| storage-service | 9010 | `/put`, `/get`, `/keys` | Key/blob store |

## Run

```bash
cd go-services
go mod tidy
go run ./cmd/hello-service
```

Or build all binaries:
```bash
go build -o bin/ ./cmd/...

# Then run any service:
bin\hello-service.exe
bin\auth-service.exe
...
```

## Docker

Each service has its own multi-stage Dockerfile.

Build an image:
```bash
docker build -t go-services/hello -f cmd/hello-service/Dockerfile .
docker run -p 8080:8080 go-services/hello
```

Or use Docker Compose:
```yaml
services:
  hello:
    build:
      context: .
      dockerfile: cmd/hello-service/Dockerfile
    ports:
      - "8080:8080"
  auth:
    build:
      context: .
      dockerfile: cmd/auth-service/Dockerfile
    ports:
      - "9001:9001"
```
