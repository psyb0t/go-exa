# Testing

Refines `~/.claude/rule-details/go/writing-tests.md`.

- **TST1** Tests are black-box, in package `exa_test`, and use only the exported API.
- **TST2** Tests never call Exa. They run the client against an `httptest` server.
- **TST3** The README usage example lives in `example_test.go` too, run against a local test server. A README example change goes into `example_test.go` in the same commit.
- **TST4** Every spec patch or config change that alters wire behavior gets a test with a body shaped like what Exa sends.
- **TST5** Table tests use `testCases` and `tc`, call `t.Parallel()` in the test and in every subtest, and assert errors with `require.ErrorIs` against a sentinel.
- **TST6** `make test` runs `go test -race` and must pass.
