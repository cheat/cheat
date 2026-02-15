# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Development Commands

### Building
```bash
# Build for your architecture
make build

# Build release binaries for all platforms
make build-release

# Install cheat to your PATH
make install
```

### Testing and Quality Checks
```bash
# Run all tests
make test
go test ./...

# Run a single test
go test -run TestFunctionName ./internal/package_name

# Generate test coverage report
make coverage

# Run linter (revive)
make lint

# Run go vet
make vet

# Format code
make fmt

# Run all checks (vendor, fmt, lint, vet, test)
make check
```

### Development Setup
```bash
# Install development dependencies (revive linter, scc)
make setup

# Update and verify vendored dependencies
make vendor-update
```

## Architecture Overview

The `cheat` command-line tool is organized into several key packages:

### Command Layer (`cmd/cheat/`)
- `main.go`: Entry point, cobra command definition, flag registration, command routing
- `cmd_*.go`: Individual command implementations (view, edit, list, search, etc.)
- `completions.go`: Dynamic shell completion functions for cheatsheet names, tags, and paths
- Commands are routed via a `switch` block in the cobra `RunE` handler

### Core Internal Packages

1. **`internal/config`**: Configuration management
   - Loads YAML config from platform-specific paths
   - Manages editor, pager, colorization settings
   - Validates and expands cheatpath configurations

2. **`internal/cheatpath`**: Cheatsheet path management
   - Represents collections of cheatsheets on filesystem
   - Handles read-only vs writable paths
   - Supports filtering and validation

3. **`internal/sheet`**: Individual cheatsheet handling
   - Parses YAML frontmatter for tags and syntax
   - Implements syntax highlighting via Chroma
   - Provides search functionality within sheets

4. **`internal/sheets`**: Collection operations
   - Loads sheets from multiple cheatpaths
   - Consolidates duplicates (local overrides global)
   - Filters by tags and sorts results

5. **`internal/display`**: Output formatting
   - Writes to stdout or pager
   - Handles text formatting and indentation

6. **`internal/installer`**: First-run installer
   - Prompts user for initial configuration choices
   - Generates default `conf.yml` and downloads community cheatsheets

7. **`internal/repo`**: Git repository management
   - Clones community cheatsheet repositories
   - Updates existing repositories

### Key Design Patterns

- **Filesystem-based storage**: Cheatsheets are plain text files
- **Override mechanism**: Local sheets override community sheets with same name
- **Tag system**: Sheets can be categorized with tags in frontmatter
- **Multiple cheatpaths**: Supports personal, community, and directory-scoped sheets
- **Directory-scoped discovery**: Walks up from cwd to find the nearest `.cheat` directory (like `.git` discovery)

### Sheet Format

Cheatsheets are plain text files optionally prefixed with YAML frontmatter:
```
---
syntax: bash
tags: [ networking, ssh ]
---
# SSH tunneling example
ssh -L 8080:localhost:80 user@remote
```

### Working with the Codebase

- Always check for `.git` directories and skip them during filesystem walks
- Use `go-git` for repository operations, not exec'ing git commands
- Platform-specific paths are handled in `internal/config/paths.go`
- Color output uses ANSI codes via the Chroma library
- Test files use the `mocks` package for test data