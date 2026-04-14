# Contributing

## Requirements

- [Go](https://go.dev/dl/) 1.25+
- [Task](https://taskfile.dev/installation/)
- [staticcheck](https://staticcheck.dev): `go install honnef.co/go/tools/cmd/staticcheck@latest`
- [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck): `go install golang.org/x/vuln/cmd/govulncheck@latest`

## Development

```sh
task build   # build binary to bin/
task check   # run fmt, vet, lint, vuln, and test
```
