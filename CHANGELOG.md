# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.3, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.
- The skill's frontmatter is valid YAML again; an unquoted colon in the description stopped it from parsing. The setup reference says `ResearchTasksCreate` returns the task in `JSON201`, not `JSON200`.

## v0.1.2, 2026-10-17

- Documentation only. The README and `AGENTS.md` describe the daily regeneration workflow: what it checks, how it fixes failing checks, how it picks the next version and which paths a fix may not touch.

## v0.1.1, 2026-10-16

- Added the `go-exa` agent skill. The README's new Agent integrations section has the install commands for Claude Code, Codex and ClawHub.

## v0.1.0, 2026-10-16

- Added `exa`, a REST client generated with oapi-codegen from Exa's OpenAPI document (spec version 1.2.0). It covers search, find similar, contents, answer and the research task endpoints.
- `make generate` downloads the current spec from the exa-labs/openapi-spec repository and regenerates the client. The spec is not committed.
