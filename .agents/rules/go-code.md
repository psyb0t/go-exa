# Go code

Refines `~/.claude/rule-details/go/wrapping-errors.md` and `~/.claude/rule-details/coding-style.md`.

- **GO1** `client.gen.go` is generated. Never edit it by hand. Change `config.yaml`, or add a spec patch script under `scripts/` and call it from `gen.go` between the download and oapi-codegen, then run `make generate` and commit the result.
- **GO2** A spec patch says in its header which spec problem it fixes. Today the spec generates as published and `scripts/` holds no patch.
- **GO3** The downloaded spec in `api/` is gitignored and never committed.
- **GO4** Hand-written Go (`gen.go` and the tests) wraps errors with `ctxerrors.Wrap` against `ctxerrors/commerr` sentinels. No `fmt.Errorf`, no `errors.New` outside a sentinel declaration.
- **GO5** The exported API is whatever oapi-codegen emits. A config or spec change that renames or removes an exported name is a breaking release (see `releases.md`).
- **GO6** `make lint` passes with zero issues. Hand-written lines are at most 80 columns, the `.golangci.yml` limit.
