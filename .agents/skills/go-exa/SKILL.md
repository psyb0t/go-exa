---
name: go-exa
description: Unofficial Go client for the Exa search API, generated with oapi-codegen from the OpenAPI spec Exa publishes. Package `exa` (`github.com/psyb0t/go-exa`) talks to `https://api.exa.ai` with typed methods for every endpoint, including `SearchWithResponse` (web search with filters by domain, date, category and text, optional page contents), `FindSimilarWithResponse`, `GetContentsWithResponse` (text, highlights and summaries for URLs or result IDs), `AnswerWithResponse` (an answer with citations) and the research task endpoints. Use when Go code needs web search, page contents or cited answers from Exa, or already imports go-exa.
license: MIT
homepage: https://github.com/psyb0t/go-exa
user-invocable: true
metadata:
  openclaw:
    emoji: "🔎"
    requires:
      bins:
        - go
permissions:
  filesystem: "Guides edits to the Go source of the current project. The library reads and writes no files."
  shell: "The Go toolchain (`go get`, `go build`, `go test`) to add and compile the dependency."
  network: "`go get` fetches the module from your Go module proxy. At runtime the code you write talks HTTPS to api.exa.ai only."
---

# go-exa

Unofficial, typed Go client for the Exa search API, in package `exa` at the module root.

For install, every operation, request options and response fields, see [references/setup.md](references/setup.md).

## Security & safety

- Exa authenticates with an API key in the `x-api-key` header. Read it from the `EXA_API_KEY` environment variable. Never hardcode, commit or log it, and never print a request with its headers.
- Requests go only to `https://api.exa.ai`. Search queries, URLs and answer prompts you send leave the machine for Exa, so do not put secrets or private data in them.
- Every call is billed to the key's account; responses report the cost in `CostDollars`. Research tasks run long and cost more, so start them only when the user asked.

## When to use

- Go code needs web search results, similar pages, page text or highlights, or an answer with citations.
- The project already imports `github.com/psyb0t/go-exa`.

## When NOT to use

- The code is not Go. This module ships only a Go package.
- You need a client maintained by Exa. This one is unofficial and not affiliated with Exa.
- You want to change the generated client itself. See `AGENTS.md` in the repo.

## Finding endpoints and types

The examples below cover a few calls. Every operation in the spec has a generated method, and the spec's descriptions are on them as Go doc comments. Look names and fields up instead of guessing. After `go get github.com/psyb0t/go-exa`, from inside your module:

```bash
go doc github.com/psyb0t/go-exa ClientWithResponsesInterface  # every endpoint method
go doc github.com/psyb0t/go-exa SearchJSONBody                # one method, params type or model
go doc -all github.com/psyb0t/go-exa                          # the whole package
```

The same documentation is at https://pkg.go.dev/github.com/psyb0t/go-exa.

Method names come from the spec's operation IDs. `<OperationId>WithResponse` takes the context, then the path parameters in order, then `*<OperationId>Params` when the operation has query or header parameters, then the request body when it has one. The `JSON200` field (or `JSON201` and so on, per declared status) holds the decoded body; it is nil for any other status. Response wrappers that share a name with a schema get a `2` suffix, as described under Usage.

## Usage

```go
import (
	"context"
	"fmt"
	"net/http"
	"os"

	exa "github.com/psyb0t/go-exa"
)

auth := exa.WithRequestEditorFn(
	func(_ context.Context, req *http.Request) error {
		req.Header.Set("x-api-key", os.Getenv("EXA_API_KEY"))

		return nil
	},
)

client, err := exa.NewClientWithResponses("https://api.exa.ai", auth)
if err != nil {
	return err
}

numResults := 5

search, err := client.SearchWithResponse(ctx, exa.SearchJSONRequestBody{
	Query:          "open source vector databases",
	NumResults:     &numResults,
	IncludeDomains: &[]string{"github.com"},
})
if err != nil {
	return err // transport failure
}

if search.JSON200 == nil {
	return fmt.Errorf("search: %s", search.Status())
}

if search.JSON200.Results != nil {
	for _, result := range *search.JSON200.Results {
		fmt.Println(*result.Url)
	}
}

answer, err := client.AnswerWithResponse(ctx, exa.AnswerJSONRequestBody{
	Query: "What is the boiling point of water at sea level?",
})
if err != nil {
	return err
}

if answer.JSON200 != nil && answer.JSON200.Answer != nil {
	text, err := answer.JSON200.Answer.AsAnswerResponseAnswer0() // string
	if err == nil {
		fmt.Println(text)
	}
}
```

Most response fields are pointers; check for nil before dereferencing. The error return covers only transport and decoding failures. An API error comes back with `JSON200` nil: check `StatusCode()` and `Body`.

Response wrappers that share a name with a schema get a `2` suffix: `SearchWithResponse` returns `*exa.SearchResponse2`, whose `JSON200` is the `*exa.SearchResponse` schema.
