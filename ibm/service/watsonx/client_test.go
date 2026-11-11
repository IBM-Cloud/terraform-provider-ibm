// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package watsonx

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testClient(token string) *client {
	return &client{token: func() (string, error) { return token, nil }, http: http.DefaultClient}
}

func TestClientDoSendsAuthAndBody(t *testing.T) {
	var gotAuth, gotType, gotMethod string
	var gotBody map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth, gotType, gotMethod = r.Header.Get("Authorization"), r.Header.Get("Content-Type"), r.Method
		b, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(b, &gotBody)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"location":"/v2/projects/abc"}`))
	}))
	defer srv.Close()

	var out struct {
		Location string `json:"location"`
	}
	_, err := testClient("raw-token").do(context.Background(), http.MethodPost, srv.URL, map[string]string{"name": "p1"}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotAuth != "Bearer raw-token" {
		t.Errorf("Authorization = %q, want Bearer prefix added", gotAuth)
	}
	if gotType != "application/json" || gotMethod != http.MethodPost {
		t.Errorf("Content-Type = %q, method = %q", gotType, gotMethod)
	}
	if gotBody["name"] != "p1" {
		t.Errorf("body = %v", gotBody)
	}
	if out.Location != "/v2/projects/abc" {
		t.Errorf("location = %q", out.Location)
	}
}

func TestClientDoKeepsExistingBearerPrefix(t *testing.T) {
	var gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
	}))
	defer srv.Close()

	if _, err := testClient("Bearer abc").do(context.Background(), http.MethodGet, srv.URL, nil, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotAuth != "Bearer abc" {
		t.Errorf("Authorization = %q", gotAuth)
	}
}

func TestClientDoReturnsAPIErrorOnFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error":"gone"}`))
	}))
	defer srv.Close()

	_, err := testClient("t").do(context.Background(), http.MethodGet, srv.URL, nil, nil)
	var ae *apiError
	if !errors.As(err, &ae) || ae.StatusCode != http.StatusNotFound || !isNotFound(err) {
		t.Fatalf("want 404 apiError, got %v", err)
	}
}

func TestClientDoPropagatesTokenError(t *testing.T) {
	c := &client{token: func() (string, error) { return "", errors.New("no session") }, http: http.DefaultClient}
	if _, err := c.do(context.Background(), http.MethodGet, "http://unused", nil, nil); err == nil {
		t.Fatal("expected token error")
	}
}
