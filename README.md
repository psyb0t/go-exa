# go-exa

[![Go Reference](https://pkg.go.dev/badge/github.com/psyb0t/go-exa.svg)](https://pkg.go.dev/github.com/psyb0t/go-exa)
[![CI](https://github.com/psyb0t/go-exa/actions/workflows/pipeline.yml/badge.svg?branch=main)](https://github.com/psyb0t/go-exa/actions/workflows/pipeline.yml)
[![version](https://raw.githubusercontent.com/psyb0t/go-exa/badges/version.svg)](https://github.com/psyb0t/go-exa/tags)
[![license](https://raw.githubusercontent.com/psyb0t/go-exa/badges/license.svg)](LICENSE)
[![imported by](https://raw.githubusercontent.com/psyb0t/go-exa/badges/importers.svg)](https://github.com/psyb0t/go-exa/blob/badges/importers.md)

Unofficial Go client for the [Exa API](https://exa.ai), generated with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from the [OpenAPI spec](https://github.com/exa-labs/openapi-spec) Exa publishes.

The package `exa` lives in the repository root and talks to `https://api.exa.ai`. `client.gen.go` is generated. Don't edit it by hand, regenerate it.

## Install

```sh
go get github.com/psyb0t/go-exa
```

## Usage

Every endpoint in the spec has a typed method on `ClientWithResponses`: `Search`, `FindSimilar`, `GetContents`, `Answer`, `ResearchTasksCreate`, `ResearchTasksList` and `ResearchControllerV0GetResearchTask`.

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	exa "github.com/psyb0t/go-exa"
)

func main() {
	auth := exa.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		req.Header.Set("x-api-key", os.Getenv("EXA_API_KEY"))

		return nil
	})

	client, err := exa.NewClientWithResponses("https://api.exa.ai", auth)
	if err != nil {
		panic(err)
	}

	res, err := client.SearchWithResponse(
		context.Background(),
		exa.SearchJSONRequestBody{Query: "latest research in LLMs"},
	)
	if err != nil {
		panic(err)
	}

	if res.JSON200 == nil || res.JSON200.Results == nil {
		return
	}

	for _, result := range *res.JSON200.Results {
		fmt.Println(*result.Url)
	}
}
```

`example_test.go` runs the same code against a local test server, so `make test` fails if the example stops compiling.

The spec reuses some names for both a schema and an operation's response wrapper. The schemas keep the plain name (`SearchResponse`, `AnswerResponse`, `FindSimilarResponse`) and the wrappers returned by the `...WithResponse` methods get a `2` suffix (`SearchResponse2`).

## Auth

Exa authenticates with an API key from the [dashboard](https://dashboard.exa.ai), sent in the `x-api-key` header. The spec also accepts `Authorization: Bearer <key>`. The generated client sets neither, so pass a `RequestEditorFn` as shown above, either to the constructor with `WithRequestEditorFn` or per call as the last argument.

## Regenerating

Exa's spec is not committed. `make generate` downloads the current OpenAPI document into the gitignored `api/` directory and regenerates the client:

```sh
make generate
```

`gen.go` declares the steps, so `make generate` is just `go generate ./...`:

- `scripts/fetch-spec.sh` downloads `exa-openapi-spec.yaml` from [exa-labs/openapi-spec](https://github.com/exa-labs/openapi-spec) to `api/openapi.yaml`.
- oapi-codegen generates `client.gen.go` with `config.yaml`.

## Development

Every command runs in the `Dockerfile.dev` image. The host needs only Docker and Make.

```sh
make help      # list targets
make dep       # go mod tidy + vendor
make lint      # shellcheck, shfmt, go fix, golangci-lint
make test      # go test -race
make sec       # gitleaks, govulncheck, semgrep
make generate  # download the spec and regenerate
```

## License

[MIT](LICENSE). The generated code comes from Exa's published API specification. The specification itself is not redistributed here.
