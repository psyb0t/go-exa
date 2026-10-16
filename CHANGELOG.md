# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.1, 2026-10-16

- Added the `go-exa` agent skill. The README's new Agent integrations section has the install commands for Claude Code, Codex and ClawHub.

## v0.1.0, 2026-10-16

- Added `exa`, a REST client generated with oapi-codegen from Exa's OpenAPI document (spec version 1.2.0). It covers search, find similar, contents, answer and the research task endpoints.
- `make generate` downloads the current spec from the exa-labs/openapi-spec repository and regenerates the client. The spec is not committed.
