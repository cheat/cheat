# Architecture Decision Records

This directory contains Architecture Decision Records (ADRs) for the cheat project.

## What is an ADR?

An Architecture Decision Record captures an important architectural decision made along with its context and consequences. ADRs help us:

- Document why decisions were made
- Understand the context and trade-offs
- Review decisions when requirements change
- Onboard new contributors

## ADR Format

Each ADR follows this template:

1. **Title**: ADR-NNN: Brief description
2. **Date**: When the decision was made
3. **Status**: Proposed, Accepted, Deprecated, Superseded
4. **Context**: What prompted this decision?
5. **Decision**: What did we decide to do?
6. **Consequences**: What are the positive, negative, and neutral outcomes?

## Index of ADRs

| ADR | Title | Status | Date |
|-----|-------|--------|------|
| [001](001-path-traversal-protection.md) | Path Traversal Protection for Cheatsheet Names | Accepted | 2025-01-21 |
| [002](002-environment-variable-parsing.md) | No Defensive Checks for Environment Variable Parsing | Accepted | 2025-01-21 |
| [003](003-search-parallelization.md) | No Parallelization for Search Operations | Accepted | 2025-01-22 |

## Creating a New ADR

1. Copy the template from an existing ADR
2. Use the next sequential number
3. Fill in all sections
4. Include the ADR alongside the commit implementing the decision