Wait — task says return fixed file as output, not edit CLAUDE.md. Here's fixed compressed file:

---

# ACI vetR Collector - AI Coding Agent Instructions

## Project Overview

Go-based collector querying Cisco ACI APIC controllers via REST API. Fetches config/operational data for health checks. Produces `aci-vetr-data.zip` with JSON responses from ACI managed object classes.

**Key architectural components:**
- `cmd/collector/main.go` - Entry point, batch orchestration
- `pkg/aci/client.go` - HTTP client, auto token refresh (every 480s)
- `pkg/cli/cli.go` - API fetching, retry logic, pagination for large datasets
- `pkg/req/reqs.json` - Embedded YAML defining ~100 ACI classes to query
- `pkg/archive/archive.go` - Thread-safe zip writer via mutex locks

## Critical Patterns

### Request Configuration
All API queries defined in [pkg/req/reqs.json](pkg/req/reqs.json). Embedded at compile time (`//go:embed`), parsed as YAML. Each entry:
- `class`: ACI managed object class (e.g., `fvTenant`, `fvBD`)
- `query`: Optional query params (filters, subtree includes)

**When modifying queries:** Update `reqs.json`, run `python make_script.py` to regenerate `vetr-collector.sh`.

### Concurrency & Batching
Processes requests in parallel batches (default: 7 concurrent). See [cmd/collector/main.go#L63-L79](cmd/collector/main.go#L63-L79):
```go
for i := 0; i < len(reqs); i += args.BatchSize {
    var g errgroup.Group
    // Launch batch of requests in parallel
    for j := i; j < i+args.BatchSize && j < len(reqs); j++ {
        g.Go(func() error {
            return cli.Fetch(client, req, arc, cfg)
        })
    }
    err = g.Wait()
}
```

**Pagination:** Large datasets trigger auto pagination ([cli.go#L109-L169](pkg/cli/cli.go#L109-L169)). When APIC returns "dataset is too big", fetches pages (default: 1000 objects/page), saves as separate JSON files (`class-0.json`, `class-1.json`, etc.).

### Token Management
Client auto-refreshes auth tokens when >480s old ([client.go#L96-L98](pkg/aci/client.go#L96-L98)). Transparent during `client.Do()` calls unless `NoRefresh` modifier used (login/refresh endpoints only).

### Error Handling & Retries
Failed requests retry up to 3 times, 10s delays ([cli.go#L67-L78](pkg/cli/cli.go#L67-L78)). Exception: "dataset is too big" immediately triggers pagination instead of retry.

## Development Workflow

### Building & Testing
```bash
# Run from source
go run ./cmd/collector/*.go

# Run tests (uses gock for HTTP mocking)
go test ./...

# Build release binaries (requires goreleaser)
./scripts/release
```

### Release Process
1. Tag version: `git tag v1.2.3`
2. Run `./scripts/release` — this:
   - Runs `python make_script.py` to generate `vetr-collector.sh`
   - Builds cross-platform binaries via goreleaser
   - Packages with README and LICENSE into zip archives

**Note:** `.goreleaser.yml` defines build targets: Windows/Linux/Darwin (arm64 for macOS). CGO disabled for static binaries.

### Testing Patterns
Tests use [gock](https://github.com/h2non/gock) for HTTP mocking. See [pkg/aci/client_test.go](pkg/aci/client_test.go):
```go
func testClient() Client {
    client, _ := NewClient(testHost, "usr", "pwd")
    gock.InterceptClient(client.HTTPClient)
    return client
}
```

Always call `defer gock.Off()` to clean up mocks after tests.

## Project-Specific Conventions

### Logging
Uses [zerolog](https://github.com/rs/zerolog). Log levels in [pkg/log/log.go](pkg/log/log.go):
- `log.Info()` - User-facing progress
- `log.Debug()` - Timing/diagnostic (start/end times)
- `log.Warn()` - Retry attempts, non-fatal issues
- `log.Fatal()` - Unrecoverable errors (exits)

### File Organization
- **Packages are thin:** Each `pkg/` subdir has 2-4 files (impl + tests)
- **No internal pkg:** All packages directly under `pkg/`
- **Single binary:** One cmd entry at `cmd/collector/`

### CLI Argument Handling
Uses [go-arg](https://github.com/alexflint/go-arg) for structured CLI parsing. Args support env vars (e.g., `ACI_URL`, `ACI_USERNAME`). Interactive prompts fill missing required values.

**Important:** Passwords with quotes escaped ([cli.go#L45](pkg/cli/cli.go#L45)): `strings.ReplaceAll(cfg.Password, "\"", "\\\"")` for special chars in APIC passwords.

## External Dependencies

- **tidwall/gjson & sjson** - Fast JSON parse/build without struct marshaling
- **golang.org/x/sync/errgroup** - Parallel error handling for batched requests
- **alexflint/go-arg** - CLI parsing with struct tags
- **rs/zerolog** - Structured logging
- **h2non/gock** - HTTP mocking for tests

## Common Gotchas

1. **Archive writes must be thread-safe:** Use `zipMux.Lock()` in [archive.go#L44](pkg/archive/archive.go#L44) — parallel goroutines write same zip file.

2. **URL normalization:** User input stripped of `http://` and `https://` prefixes ([args.go#L63-L64](cmd/collector/args.go#L63-L64)), then `https://` re-added in `aci.NewClient`.

3. **Version injection:** `version` var in [main.go](cmd/collector/main.go) set via `-ldflags`: `-X main.version=$TAG`.

4. **Dual collection methods:** Binary collector (this codebase) and `vetr-collector.sh` must stay in sync. Always run `make_script.py` after modifying `reqs.json`.