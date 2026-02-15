# ADR-004: Recursive `.cheat` Directory Search

Date: 2026-02-15

## Status

Accepted

## Context

Previously, `cheat` only checked the current working directory for a `.cheat`
subdirectory to use as a directory-scoped cheatpath. If a user was in
`~/projects/myapp/src/handlers/` but the `.cheat` directory lived at
`~/projects/myapp/.cheat`, it would not be found. Users requested (#602) that
`cheat` walk up the directory hierarchy to find the nearest `.cheat`
directory, mirroring the discovery pattern used by `git` for `.git`
directories.

## Decision

Walk upward from the current working directory to the filesystem root, and
stop at the first `.cheat` directory found. Only directories are matched (a
file named `.cheat` is ignored).

### Stop at first `.cheat` found

Rather than collecting multiple `.cheat` directories from ancestor directories:

- Matches `.git` discovery semantics, which users already understand
- Fits the existing single-cheatpath-named-`"cwd"` code without structural
  changes
- Avoids precedence and naming complexity when multiple `.cheat` directories
  exist in the ancestor chain
- `cheat` already supports multiple cheatpaths via `conf.yml` for users who
  want that; directory-scoped `.cheat` serves the project-context use case

### Walk to filesystem root (not `$HOME`)

Rather than stopping the search at `$HOME`:

- Simpler implementation with no platform-specific home-directory detection
- Supports sysadmins working in `/etc`, `/srv`, `/var`, or other paths
  outside `$HOME`
- The boundary only matters on the failure path (no `.cheat` found anywhere),
  where the cost is a few extra `stat` calls
- Security is not a concern since cheatsheets are display-only text, not
  executable code

## Consequences

### Positive
- Users can place `.cheat` at their project root and it works from any
  subdirectory, matching their mental model
- No configuration changes needed; existing `.cheat` directories continue to
  work identically
- Minimal code change (one small helper function)

### Negative
- A `.cheat` directory in an unexpected ancestor could be picked up
  unintentionally, though this is unlikely in practice and matches how `.git`
  works

### Neutral
- The cheatpath name remains `"cwd"` regardless of which ancestor the `.cheat`
  was found in

## Alternatives Considered

### 1. Stop at `$HOME`
**Rejected**: Adds platform-specific complexity for minimal benefit. The only
downside of walking to root is a few extra `stat` calls on the failure path.

### 2. Collect multiple `.cheat` directories
**Rejected**: Introduces precedence and naming complexity. Users who want
multiple cheatpaths can configure them in `conf.yml`.

## References

- GitHub issue: #602
- Implementation: `findLocalCheatpath()` in `internal/config/config.go`
