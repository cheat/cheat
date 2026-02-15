# ADR-003: No Parallelization for Search Operations

Date: 2025-01-22

## Status

Accepted

## Context

We investigated optimizing cheat's search performance through parallelization. Initial assumptions suggested that I/O operations (reading multiple cheatsheet files) would be the primary bottleneck, making parallel processing beneficial.

Performance benchmarks were implemented to measure search operations, and a parallel search implementation using goroutines was created and tested.

## Decision

We will **not** implement parallel search. The sequential implementation will remain unchanged.

## Rationale

### Performance Profile Analysis

CPU profiling revealed that search performance is dominated by:
- **Process creation overhead** (~30% in `os/exec.(*Cmd).Run`)
- **System calls** (~30% in `syscall.Syscall6`)
- **Process management** (fork, exec, pipe setup)

The actual search logic (regex matching, file reading) was negligible in the profile, indicating our optimization efforts were targeting the wrong bottleneck.

### Benchmark Results

Parallel implementation showed minimal improvements:
- Simple search: 17ms → 15.3ms (10% improvement)
- Regex search: 15ms → 14.9ms (minimal improvement)  
- Colorized search: 19.5ms → 16.8ms (14% improvement)
- Complex regex: 20ms → 15.3ms (24% improvement)

The best case saved only ~5ms in absolute terms.

### Cost-Benefit Analysis

**Costs of parallelization:**
- Added complexity with goroutines, channels, and synchronization
- Increased maintenance burden
- More difficult debugging and testing
- Potential race conditions

**Benefits:**
- 5-15% performance improvement (5ms in real terms)
- Imperceptible to users in interactive use

### User Experience Perspective

For a command-line tool:
- Current 15-20ms response time is excellent
- Users cannot perceive 5ms differences
- Sub-50ms is considered "instant" in HCI research

## Consequences

### Positive
- Simpler, more maintainable codebase
- Easier to debug and reason about
- No synchronization bugs or race conditions
- Focus remains on code clarity

### Negative
- Missed opportunity for ~5ms performance gain
- Search remains single-threaded

### Neutral
- Performance remains excellent for intended use case
- Follows Go philosophy of preferring simplicity

## Alternatives Considered

### 1. Keep Parallel Implementation
**Rejected**: Complexity outweighs negligible performance gains.

### 2. Optimize Process Startup
**Rejected**: Process creation overhead is inherent to CLI tools and cannot be avoided without fundamental architecture changes.

### 3. Future Optimizations
If performance becomes critical, consider:
- **Long-running daemon**: Eliminate process startup overhead entirely
- **Shell function**: Reduce fork/exec overhead
- **Compiled-in cheatsheets**: Eliminate file I/O

However, these would fundamentally change the tool's architecture and usage model.

## Notes

This decision reinforces important principles:
1. Always profile before optimizing
2. Consider the full execution context
3. Measure what matters to users
4. Complexity has a real cost

The parallelization attempt was valuable as a learning exercise and definitively answered whether this optimization path was worthwhile.

## References

- Benchmark implementation: cmd/cheat/search_bench_test.go
- Reverted parallel implementation: see git history (commit 82eb918)