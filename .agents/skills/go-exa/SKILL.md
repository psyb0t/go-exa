---
name: go-exa
description: Unofficial Go client for the Exa search API, generated with oapi-codegen from the OpenAPI spec Exa publishes. Package `exa` (`github.com/psyb0t/go-exa`) talks to `https://api.exa.ai` with typed methods for every endpoint: `SearchWithResponse` (web search with filters by domain, date, category and text, optional page contents), `FindSimilarWithResponse`, `GetContentsWithResponse` (text, highlights and summaries for URLs or result IDs), `AnswerWithResponse` (an answer with citations) and the research task endpoints. Use when Go code needs web search, page contents or cited answers from Exa, or already imports go-exa.
homepage: https://github.com/psyb0t/go-exa
user-invocable: true
metadata:
  openclaw:
    emoji: "🔎"
    requires:
      bins:
        - go
      env:
        - EXA_API_KEY
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
