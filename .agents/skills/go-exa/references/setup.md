# go-exa: setup and reference

This module is unofficial and not affiliated with Exa.

## Requirements

- Go 1.26.9 or newer, per the module's `go.mod`.
- An Exa API key from the [dashboard](https://dashboard.exa.ai), exported as `EXA_API_KEY`.

## Install

```bash
go get github.com/psyb0t/go-exa
```

The module path ends in `go-exa` and the package is `exa`, so import it with an explicit name:

```go
import exa "github.com/psyb0t/go-exa"
```

## Auth

Send the key as `x-api-key` from a request editor, either for every call with `exa.WithRequestEditorFn` on the constructor or per call as the last argument. The spec also accepts `Authorization: Bearer <key>`. The generated client sets neither on its own.

## Client

- `exa.NewClientWithResponses("https://api.exa.ai", opts...)`. Options: `exa.WithRequestEditorFn`, `exa.WithHTTPClient` (your own `*http.Client`, for timeouts), `exa.WithBaseURL`.
- Each operation has `Client.X` (raw `*http.Response`), `ClientWithResponses.XWithResponse` (parsed) and `NewXRequest` (request only). Operations with a body also have `XWithBodyWithResponse` taking a raw `io.Reader`.

## Operations

| Method | Endpoint | Body or params | `JSON200` |
|---|---|---|---|
| `SearchWithResponse` | `POST /search` | `SearchJSONRequestBody` (`Query` required) | `*SearchResponse` |
| `FindSimilarWithResponse` | `POST /findSimilar` | `FindSimilarJSONRequestBody` (`Url` required) | `*FindSimilarResponse` |
| `GetContentsWithResponse` | `POST /contents` | `GetContentsJSONRequestBody` (`Urls` required) | `*ContentsResponse` |
| `AnswerWithResponse` | `POST /answer` | `AnswerJSONRequestBody` (`Query` required) | `*AnswerResponse` |
| `ResearchTasksCreateWithResponse` | `POST /research/v0/tasks` | `ResearchTasksCreateJSONRequestBody` (`Instructions` required) | nil; the created task is in `JSON201` |
| `ResearchTasksListWithResponse` | `GET /research/v0/tasks` | `*ResearchTasksListParams` (`Cursor`, `Limit`) | a page of tasks |
| `ResearchControllerV0GetResearchTaskWithResponse` | `GET /research/v0/tasks/{id}` | task ID | the task and its status |

Search options on `SearchJSONRequestBody`: `Type` (`SearchJSONBodyTypeAuto`, `...Fast`, `...Neural`, `...Deep` and more), `NumResults`, `Category`, `IncludeDomains`, `ExcludeDomains`, `IncludeText`, `ExcludeText`, `StartPublishedDate`, `EndPublishedDate`, `StartCrawlDate`, `EndCrawlDate`, `UserLocation`, and `Contents` (a `*ContentsRequest`) to get page text, highlights or summaries back with the results.

## Responses

| Field or method | Holds |
|---|---|
| `JSON200` (`JSON201` for `ResearchTasksCreateWithResponse`) | The decoded body on success, nil otherwise |
| `Body`, `StatusCode()`, `Status()`, `HTTPResponse` | The raw response |

- `SearchResponse.Results` is `*[]ResultWithContent`: `Url`, `Title`, `Author`, `PublishedDate`, `Id`, plus `Text`, `Highlights` and `Summary` when you asked for contents. `RequestId`, `SearchType` and `CostDollars` describe the call.
- Some fields are unions, a JSON value that can be one of several shapes. Read them with the generated `As...` methods: `AnswerResponse.Answer.AsAnswerResponseAnswer0()` gives the answer as a string, `AsAnswerResponseAnswer1()` gives the structured answer when the request set `OutputSchema`. Set a union field on a request with the matching `From...` method.
- Wrappers that share a name with a schema get a `2` suffix: `SearchResponse2`, `AnswerResponse2`, `FindSimilarResponse2`.
