# Hacking Guide

This document provides a comprehensive guide for developing `cheat`, including setup, architecture overview, and code patterns.

## Quick Start

### 1. Install system dependencies

The following are required and must be available on your `PATH`:
- `git`
- `go` (>= 1.19 is recommended)
- `make`

Optional dependencies:
- `docker`
- `pandoc` (necessary to generate a `man` page)

### 2. Install utility applications
Run `make setup` to install `scc` and `revive`, which are used by various `make` targets.

### 3. Development workflow

1. Make changes to the `cheat` source code
2. Run `make test` to run unit-tests
3. Fix compiler errors and failing tests as necessary
4. Run `make build`. A `cheat` executable will be written to the `dist` directory
5. Use the new executable by running `dist/cheat <command>`
6. Run `make install` to install `cheat` to your `PATH`
7. Run `make build-release` to build cross-platform binaries in `dist`
8. Run `make clean` to clean the `dist` directory when desired

You may run `make help` to see a list of available `make` commands.

### 4. Testing

#### Unit Tests
Run unit tests with:
```bash
make test
```

#### Integration Tests
Integration tests that require network access are separated using build tags. Run them with:
```bash
make test-integration
```

To run all tests (unit and integration):
```bash
make test-all
```

#### Test Coverage
Generate a coverage report with:
```bash
make coverage        # HTML report
make coverage-text   # Terminal output
```

## Architecture Overview

### Package Structure

The `cheat` application follows a clean architecture with well-separated concerns:

- **`cmd/cheat/`**: Command layer (cobra-based CLI, flag registration, command routing, shell completions)
- **`internal/config`**: Configuration management (YAML loading, validation, paths)
- **`internal/cheatpath`**: Cheatsheet path management (collections, filtering)
- **`internal/sheet`**: Individual cheatsheet handling (parsing, search, highlighting)  
- **`internal/sheets`**: Collection operations (loading, consolidation, filtering)
- **`internal/display`**: Output formatting (pager integration, colorization)
- **`internal/repo`**: Git repository management for community sheets

### Key Design Patterns

- **Filesystem-based storage**: Cheatsheets are plain text files
- **Override mechanism**: Local sheets override community sheets with same name
- **Tag system**: Sheets can be categorized with tags in frontmatter
- **Multiple cheatpaths**: Supports personal, community, and directory-scoped sheets

## Core Types and Functions

### Config (`internal/config`)

The main configuration structure:

```go
type Config struct {
    Colorize   bool           `yaml:"colorize"`
    Editor     string         `yaml:"editor"`
    Cheatpaths []cp.Path      `yaml:"cheatpaths"`
    Style      string         `yaml:"style"`
    Formatter  string         `yaml:"formatter"`
    Pager      string         `yaml:"pager"`
    Path       string
}
```

Key functions:
- `New(confPath, resolve)` - Load config from file
- `Validate()` - Validate configuration values
- `Editor()` - Get editor from environment or defaults (package-level function)
- `Pager()` - Get pager from environment or defaults (package-level function)

### Cheatpath (`internal/cheatpath`)

Represents a directory containing cheatsheets:

```go
type Path struct {
    Name     string   // Friendly name (e.g., "personal")
    Path     string   // Filesystem path
    Tags     []string // Tags applied to all sheets in this path
    ReadOnly bool     // Whether sheets can be modified
}
```

### Sheet (`internal/sheet`)

Represents an individual cheatsheet:

```go
type Sheet struct {
    Title     string   // Sheet name (from filename)
    CheatPath string   // Name of the cheatpath this sheet belongs to
    Path      string   // Full filesystem path
    Text      string   // Content (without frontmatter)
    Tags      []string // Combined tags (from frontmatter + cheatpath)
    Syntax    string   // Syntax for highlighting
    ReadOnly  bool     // Whether sheet can be edited
}
```

Key methods:
- `New(title, cheatpath, path, tags, readOnly)` - Load from file
- `Search(reg)` - Search content with a compiled regexp
- `Colorize(conf)` - Apply syntax highlighting (modifies sheet in place)
- `Tagged(needle)` - Check if sheet has the given tag

## Common Operations

### Loading and Displaying a Sheet

```go
// Load sheet
s, err := sheet.New("tar", "personal", "/path/to/tar", []string{"personal"}, false)
if err != nil {
    log.Fatal(err)
}

// Apply syntax highlighting (modifies sheet in place)
s.Colorize(conf)

// Display with pager
display.Write(s.Text, conf)
```

### Working with Sheet Collections

```go
// Load all sheets from cheatpaths (returns a slice of maps, one per cheatpath)
allSheets, err := sheets.Load(conf.Cheatpaths)
if err != nil {
    log.Fatal(err)
}

// Consolidate to handle duplicates (later cheatpaths take precedence)
consolidated := sheets.Consolidate(allSheets)

// Filter by tag (operates on the slice of maps)
filtered := sheets.Filter(allSheets, []string{"networking"})

// Sort alphabetically (returns a sorted slice)
sorted := sheets.Sort(consolidated)
```

### Sheet Format

Cheatsheets are plain text files that may begin with YAML frontmatter:

```yaml
---
syntax: bash
tags: [networking, linux, ssh]
---
# Connect to remote server
ssh user@hostname

# Copy files over SSH
scp local_file user@hostname:/remote/path
```

## Testing

Run tests with:
```bash
make test           # Run all tests
make coverage       # Generate coverage report
go test ./...       # Go test directly
```

Test files follow Go conventions:
- `*_test.go` files in same package
- Table-driven tests for multiple scenarios
- Mock data in `mocks` package

## Error Handling

The codebase follows consistent error handling patterns:
- Functions return explicit errors
- Errors are wrapped with context using `fmt.Errorf`
- User-facing errors are written to stderr

Example:
```go
s, err := sheet.New(title, cheatpath, path, tags, false)
if err != nil {
    return fmt.Errorf("failed to load sheet: %w", err)
}
```

## Developing with Docker

It may be useful to test your changes within a pristine environment. An Alpine-based docker container has been provided for that purpose.

Build the docker container:
```bash
make docker-setup
```

Shell into the container:
```bash
make docker-sh
```

The `cheat` source code will be mounted at `/app` within the container.

To destroy the container:
```bash
make distclean
```
