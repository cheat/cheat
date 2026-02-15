# ADR-002: No Defensive Checks for Environment Variable Parsing

Date: 2025-01-21

## Status

Accepted

## Context

In `cmd/cheat/main.go` lines 47-52, the code parses environment variables assuming they all contain an equals sign:

```go
for _, e := range os.Environ() {
    pair := strings.SplitN(e, "=", 2)
    if runtime.GOOS == "windows" {
        pair[0] = strings.ToUpper(pair[0])
    }
    envvars[pair[0]] = pair[1]  // Could panic if pair has < 2 elements
}
```

If `os.Environ()` returned a string without an equals sign, `strings.SplitN` would return a slice with only one element, causing a panic when accessing `pair[1]`.

## Decision

We will **not** add defensive checks for this condition. The current code that assumes all environment strings contain "=" will remain unchanged.

## Rationale

### Go Runtime Guarantees

Go's official documentation guarantees that `os.Environ()` returns environment variables in the form "key=value". This is a documented contract of the Go runtime that has been stable since Go 1.0.

### Empirical Evidence

Testing across platforms confirms:
- All environment variables returned by `os.Environ()` contain at least one "="
- Empty environment variables appear as "KEY=" (with an empty value)
- Even Windows special variables like "=C:=C:\path" maintain the format

### Cost-Benefit Analysis

Adding defensive code would:
- **Cost**: Add complexity and cognitive overhead
- **Cost**: Suggest uncertainty about Go's documented behavior
- **Cost**: Create dead code that can never execute under normal conditions
- **Benefit**: Protect against a theoretical scenario that violates Go's guarantees

The only scenarios where this could panic are:
1. A bug in Go's runtime (extremely unlikely, would affect all Go programs)
2. Corrupted OS-level environment (would cause broader system issues)
3. Breaking change in future Go version (would break many programs, unlikely)

## Consequences

### Positive
- Simpler, more readable code
- Trust in platform guarantees reduces unnecessary defensive programming
- No performance overhead from unnecessary checks

### Negative
- Theoretical panic if Go's guarantees are violated

### Neutral
- Follows Go community standards of trusting standard library contracts

## Alternatives Considered

### 1. Add Defensive Check
```go
if len(pair) < 2 {
    continue  // or pair[1] = ""
}
```
**Rejected**: Adds complexity for a condition that should never occur.

### 2. Add Panic with Clear Message
```go
if len(pair) < 2 {
    panic("os.Environ() contract violation: " + e)
}
```
**Rejected**: Would crash the program for the same theoretical issue.

### 3. Add Comment Documenting Assumption
```go
// os.Environ() guarantees "key=value" format, so pair[1] is safe
envvars[pair[0]] = pair[1]
```
**Rejected**: While documentation is good, this particular guarantee is fundamental to Go.

## Notes

If Go ever changes this behavior (extremely unlikely as it would break compatibility), it would be caught immediately in testing as the program would panic on startup. This would be a clear signal to revisit this decision.

## References

- Go os.Environ() documentation: https://pkg.go.dev/os#Environ
- Go os.Environ() source code and tests