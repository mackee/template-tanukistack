package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/mackee/template-tanukistack/client"
	"github.com/mackee/template-tanukistack/client/database/databasetest"
	"github.com/mackee/template-tanukistack/handler"
	"github.com/mackee/template-tanukistack/record"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newClient(t *testing.T) (*client.Client, func()) {
	t.Helper()
	db, closer := databasetest.Initialize(t, "../sql/pg.sql")
	return &client.Client{
		Database: db,
	}, closer
}

type testCase[Req any, Resp any] struct {
	name        string
	setup       func(*testing.T, *client.Client)
	method      string
	requestPath string
	query       url.Values
	body        *Req
	statusCode  int
	expect      *Resp
	teardown    func(*testing.T, *client.Client)
	opts        []cmp.Option
}

type asserter interface {
	Name() string
	assert(*testing.T)
}

func (tc testCase[Req, Resp]) Name() string {
	return tc.name
}

func (tc testCase[Req, Resp]) assert(t *testing.T) {
	cli, closer := newClient(t)
	defer closer()
	if tc.setup != nil {
		tc.setup(t, cli)
	}

	server := httptest.NewServer(handler.NewHandler(cli))

	u, err := url.Parse(server.URL + tc.requestPath)
	require.NoError(t, err)
	u.RawQuery = tc.query.Encode()

	var req *http.Request
	if tc.body != nil {
		buf := &bytes.Buffer{}
		require.NoError(t, json.NewEncoder(buf).Encode(tc.body))
		req, err = http.NewRequest(tc.method, u.String(), buf)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/json")
		require.NoError(t, err)
	} else {
		req, err = http.NewRequest(tc.method, u.String(), nil)
		req.Header.Add("Accept", "application/json")
		require.NoError(t, err)
	}

	resp, err := server.Client().Do(req)
	require.NoError(t, err)
	func() {
		assert.Equal(t, tc.statusCode, resp.StatusCode)
		if tc.expect == nil {
			return
		}
		defer resp.Body.Close()
		var got Resp
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&got))
		assert.Empty(t, cmp.Diff(tc.expect, &got, tc.opts...))
	}()

	if tc.teardown != nil {
		tc.teardown(t, cli)
	}
}

func TestHandler(t *testing.T) {
	tcs := []asserter{
		testCase[handler.HelloRequest, handler.HelloResponse]{
			name:        "hello",
			method:      http.MethodGet,
			requestPath: "/api/hello",
			query: url.Values{
				"name": []string{"world"},
			},
			statusCode: http.StatusOK,
			expect: &handler.HelloResponse{
				Message: "Hello, world",
			},
		},
		testCase[handler.ListMessagesRequest, handler.ListMessagesResponse]{
			name: "list messages",
			setup: func(t *testing.T, cli *client.Client) {
				if _, err := record.NewMessageSQL().Insert().
					ValueID(42).
					ValueAuthor("author").
					ValueText("text").
					ExecContextWithoutSelect(t.Context(), cli.Database); err != nil {
					t.Fatal(err)
				}
			},
			method:      http.MethodGet,
			requestPath: "/api/messages",
			query:       url.Values{},
			statusCode:  http.StatusOK,
			expect: &handler.ListMessagesResponse{
				Messages: record.Messages{
					{
						ID:     42,
						Author: "author",
						Text:   "text",
					},
				},
			},
			opts: cmp.Options{
				cmpopts.IgnoreFields(record.Message{}, "CreatedAt", "UpdatedAt"),
			},
		},
		testCase[handler.PostMessageRequest, handler.PostMessageResponse]{
			name:        "post message",
			method:      http.MethodPost,
			requestPath: "/api/messages",
			query:       url.Values{},
			body: &handler.PostMessageRequest{
				Author: "author",
				Text:   "text",
			},
			statusCode: http.StatusOK,
			expect: &handler.PostMessageResponse{
				Message: &record.Message{
					Author: "author",
					Text:   "text",
				},
			},
			opts: cmp.Options{
				cmpopts.IgnoreFields(record.Message{}, "ID", "CreatedAt", "UpdatedAt"),
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.Name(), tc.assert)
	}
}
