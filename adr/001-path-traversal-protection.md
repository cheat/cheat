# ADR-001: Path Traversal Protection for Cheatsheet Names

Date: 2025-01-21

## Status

Accepted

## Context

The `cheat` tool allows users to create, edit, and remove cheatsheets using commands like:
- `cheat --edit <name>`
- `cheat --rm <name>`

Without validation, a user could potentially provide malicious names like:
- `../../../etc/passwd` (directory traversal)
- `/etc/passwd` (absolute path)
- `~/.ssh/authorized_keys` (home directory expansion)

While `cheat` is a local tool run by the user themselves (not a network service), path traversal could still lead to:
1. Accidental file overwrites outside cheatsheet directories
2. Confusion about where files are being created
3. Potential security issues in shared environments

## Decision

We implemented input validation for cheatsheet names to prevent directory traversal attacks. The validation rejects names that:

1. Contain `..` (parent directory references)
2. Are absolute paths (start with `/` on Unix)
3. Start with `~` (home directory expansion)
4. Are empty
5. Start with `.` (hidden files - these are not displayed by cheat)

The validation is performed at the application layer before any file operations occur.

## Implementation Details

### Validation Function

The validation is implemented in `internal/sheet/validate.go`:

```go
func Validate(name string) error {
    // Reject empty names
    if name == "" {
        return fmt.Errorf("cheatsheet name cannot be empty")
    }

    // Reject names containing directory traversal
    if strings.Contains(name, "..") {
        return fmt.Errorf("cheatsheet name cannot contain '..'")
    }

    // Reject absolute paths
    if filepath.IsAbs(name) {
        return fmt.Errorf("cheatsheet name cannot be an absolute path")
    }

    // Reject names that start with ~ (home directory expansion)
    if strings.HasPrefix(name, "~") {
        return fmt.Errorf("cheatsheet name cannot start with '~'")
    }

    // Reject hidden files (files that start with a dot)
    filename := filepath.Base(name)
    if strings.HasPrefix(filename, ".") {
        return fmt.Errorf("cheatsheet name cannot start with '.' (hidden files are not supported)")
    }

    return nil
}
```

### Integration Points

The validation is called in:
- `cmd/cheat/cmd_edit.go` - before creating or editing a cheatsheet
- `cmd/cheat/cmd_remove.go` - before removing a cheatsheet

### Allowed Patterns

The following patterns are explicitly allowed:
- Simple names: `docker`, `git`
- Nested paths: `docker/compose`, `lang/go/slice`
- Current directory references: `./mysheet`

## Consequences

### Positive

1. **Safety**: Prevents accidental or intentional file operations outside cheatsheet directories
2. **Simplicity**: Validation happens early, before any file operations
3. **User-friendly**: Clear error messages explain why a name was rejected
4. **Performance**: Minimal overhead - simple string checks
5. **Compatibility**: Doesn't break existing valid cheatsheet names

### Negative

1. **Limitation**: Users cannot use `..` in cheatsheet names even if legitimate
2. **No symlink support**: Cannot create cheatsheets through symlinks outside the cheatpath

### Neutral

1. Uses Go's `filepath.IsAbs()` which handles platform differences (Windows vs Unix)
2. No attempt to resolve or canonicalize paths - validation is purely syntactic

## Security Considerations

### Threat Model

`cheat` is a local command-line tool, not a network service. The primary threats are:
- User error (accidentally overwriting important files)
- Malicious scripts that invoke `cheat` with crafted arguments
- Shared system scenarios where cheatsheets might be shared

### What This Protects Against

- Directory traversal using `../`
- Absolute path access to system files
- Shell expansion of `~` to home directory
- Empty names that might cause unexpected behavior
- Hidden files that wouldn't be displayed anyway

### What This Does NOT Protect Against

- Users with filesystem permissions can still directly edit any file
- Symbolic links within the cheatpath pointing outside
- Race conditions (TOCTOU) - though minimal risk for a local tool
- Malicious content within cheatsheets themselves

## Testing

Comprehensive tests ensure the validation works correctly:

1. **Unit tests** (`internal/sheet/validate_test.go`) verify the validation logic
2. **Integration tests** verify the actual binary blocks malicious inputs
3. **No system files are accessed** during testing - all tests use isolated directories

Example test cases:
```bash
# These are blocked:
cheat --edit "../../../etc/passwd"
cheat --edit "/etc/passwd"
cheat --edit "~/.ssh/config"
cheat --rm ".."

# These are allowed:
cheat --edit "docker"
cheat --edit "docker/compose"
cheat --edit "./local"
```

## Alternative Approaches Considered

1. **Path resolution and verification**: Resolve the final path and check if it's within the cheatpath
   - Rejected: More complex, potential race conditions, platform-specific edge cases

2. **Chroot/sandbox**: Run file operations in a restricted environment
   - Rejected: Overkill for a local tool, platform compatibility issues

3. **Filename allowlist**: Only allow alphanumeric characters and specific symbols
   - Rejected: Too restrictive, would break existing cheatsheets with valid special characters

## References

- OWASP Path Traversal: https://owasp.org/www-community/attacks/Path_Traversal
- CWE-22: Improper Limitation of a Pathname to a Restricted Directory
- Go filepath package documentation: https://pkg.go.dev/path/filepath