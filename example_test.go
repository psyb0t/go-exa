package exa_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	exa "github.com/psyb0t/go-exa"
)

// Example is the README usage example. A local server stands in for
// https://api.exa.ai so the example runs without network access or a key.
func Example() {
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(
				`{"results":[{"id":"1","url":"https://example.com"}]}`,
			))
		},
	))
	defer srv.Close()

	auth := exa.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("x-api-key", os.Getenv("EXA_API_KEY"))

			return nil
		},
	)

	client, err := exa.NewClientWithResponses(srv.URL, auth)
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

	// Output: https://example.com
}
