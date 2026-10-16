package exa_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	exa "github.com/psyb0t/go-exa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testAPIKey = "test-key"

func newTestClient(
	t *testing.T,
	handler http.HandlerFunc,
) *exa.ClientWithResponses {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	auth := exa.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("x-api-key", testAPIKey)

			return nil
		},
	)

	client, err := exa.NewClientWithResponses(srv.URL, auth)
	require.NoError(t, err)

	return client
}

func TestSearch(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/search", r.URL.Path)
		assert.Equal(t, testAPIKey, r.Header.Get("x-api-key"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		raw, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		var body map[string]any

		assert.NoError(t, json.Unmarshal(raw, &body))
		assert.Equal(t, "go clients", body["query"])

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"requestId": "req-1",
			"results": [
				{"id": "doc-1", "url": "https://example.com",
				 "title": "Example"}
			]
		}`))
	})

	res, err := client.SearchWithResponse(
		context.Background(),
		exa.SearchJSONRequestBody{Query: "go clients"},
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)
	require.NotNil(t, res.JSON200.Results)

	results := *res.JSON200.Results
	require.Len(t, results, 1)
	assert.Equal(t, "req-1", *res.JSON200.RequestId)
	assert.Equal(t, "https://example.com", *results[0].Url)
	assert.Equal(t, "Example", *results[0].Title)
}

func TestGetResearchTask(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/research/v0/tasks/task-42", r.URL.Path)
		assert.Equal(t, testAPIKey, r.Header.Get("x-api-key"))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"id": "task-42",
			"instructions": "summarise",
			"status": "completed",
			"data": {"answer": "done"}
		}`))
	})

	res, err := client.ResearchControllerV0GetResearchTaskWithResponse(
		context.Background(),
		"task-42",
	)
	require.NoError(t, err)
	require.NotNil(t, res.JSON200)
	assert.Equal(t, "task-42", res.JSON200.Id)
	assert.Equal(
		t,
		exa.ResearchTaskDtoStatusCompleted,
		res.JSON200.Status,
	)
	require.NotNil(t, res.JSON200.Data)
	assert.Equal(t, "done", (*res.JSON200.Data)["answer"])
}
