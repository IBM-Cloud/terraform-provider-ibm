// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcecontroller

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestJSONEqual(t *testing.T) {
	if !jsonEqual(map[string]any{"a": 1, "b": []any{true}}, map[string]any{"b": []any{true}, "a": 1.0}) {
		t.Error("equal values with different numeric types should match")
	}
	if jsonEqual(map[string]any{"enable_cspm": true}, map[string]any{"enable_cspm": false}) {
		t.Error("different values must not match")
	}
}

func TestParseParametersJSON(t *testing.T) {
	for raw, ok := range map[string]bool{
		`{"enable_cspm":false,"target_accounts":[]}`: true,
		`{}`:       true,
		`not-json`: false,
		`[1,2]`:    false,
		`null`:     false,
		`"text"`:   false,
	} {
		if _, err := parseParametersJSON(raw); (err == nil) != ok {
			t.Errorf("%s: err=%v, want ok=%v", raw, err, ok)
		}
	}
}

func TestIsInstanceGone(t *testing.T) {
	for code, want := range map[int]bool{404: true, 410: true, 403: false, 500: false} {
		if got := isInstanceGone(&core.DetailedResponse{StatusCode: code}); got != want {
			t.Errorf("%d: got %v", code, got)
		}
	}
	if isInstanceGone(nil) {
		t.Error("nil response is not gone")
	}
}

// fakeRC answers PATCH with the given status/async flag, then serves the GET
// responses in order (the last one repeats).
type fakeRC struct {
	mu        sync.Mutex
	patchCode int
	async     bool
	gets      []string // "<instance state>/<last operation state>"
	patchBody string
	getCount  int
}

func (f *fakeRC) handler(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	instance := func(state, opState string, async bool) string {
		return fmt.Sprintf(`{"id":"crn:i","guid":"i","state":%q,"last_operation":{"type":"update","state":%q,"async":%t,"description":"op %s","cancelable":false,"poll":true}}`,
			state, opState, async, opState)
	}
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPatch:
		b, _ := io.ReadAll(r.Body)
		f.patchBody = string(b)
		if f.patchCode >= 400 {
			w.WriteHeader(f.patchCode)
			_, _ = w.Write([]byte(`{"message":"gone"}`))
			return
		}
		opState := "succeeded"
		if f.async {
			opState = "in progress"
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(instance("active", opState, f.async)))
	case http.MethodGet:
		g := f.gets[min(f.getCount, len(f.gets)-1)]
		f.getCount++
		parts := strings.SplitN(g, "/", 2)
		_, _ = w.Write([]byte(instance(parts[0], parts[1], true)))
	}
}

func newTestResource(t *testing.T, f *fakeRC) *resourceInstanceParametersResource {
	t.Helper()
	resourceInstanceParametersPollDelay, resourceInstanceParametersPollInterval = 0, time.Millisecond
	srv := httptest.NewServer(http.HandlerFunc(f.handler))
	t.Cleanup(srv.Close)
	client, err := rc.NewResourceControllerV2(&rc.ResourceControllerV2Options{URL: srv.URL, Authenticator: &core.NoAuthAuthenticator{}})
	if err != nil {
		t.Fatal(err)
	}
	return &resourceInstanceParametersResource{client: client}
}

func model() *resourceInstanceParametersModel {
	return &resourceInstanceParametersModel{ResourceInstanceID: types.StringValue("i")}
}

func TestPatchSync(t *testing.T) {
	f := &fakeRC{gets: []string{"active/succeeded"}}
	m := model()
	if err := newTestResource(t, f).patch(context.Background(), m, `{"enable_cspm":true}`, time.Minute); err != nil {
		t.Fatal(err)
	}
	if f.getCount != 0 || m.State.ValueString() != "active" || !strings.Contains(f.patchBody, `"enable_cspm":true`) {
		t.Errorf("gets=%d state=%s body=%s", f.getCount, m.State, f.patchBody)
	}
}

func TestPatchAsyncWaitsForActive(t *testing.T) {
	f := &fakeRC{async: true, gets: []string{"active/in progress", "inactive/succeeded", "active/succeeded"}}
	m := model()
	if err := newTestResource(t, f).patch(context.Background(), m, `{"enable_cspm":true}`, time.Minute); err != nil {
		t.Fatal(err)
	}
	if f.getCount < 3 || m.State.ValueString() != "active" {
		t.Errorf("gets=%d state=%s", f.getCount, m.State)
	}
}

func TestPatchAsyncFailed(t *testing.T) {
	f := &fakeRC{async: true, gets: []string{"active/in progress", "active/failed"}}
	err := newTestResource(t, f).patch(context.Background(), model(), `{}`, time.Minute)
	if err == nil || !strings.Contains(err.Error(), "op failed") {
		t.Fatalf("want failed op error, got %v", err)
	}
}

func TestPatchGoneInstance(t *testing.T) {
	for _, code := range []int{http.StatusNotFound, http.StatusGone} {
		f := &fakeRC{patchCode: code, gets: []string{"active/succeeded"}}
		err := newTestResource(t, f).patch(context.Background(), model(), `{"enable_cspm":false}`, time.Minute)
		if !errors.Is(err, errInstanceGone) {
			t.Errorf("%d: want errInstanceGone, got %v", code, err)
		}
	}
	f := &fakeRC{patchCode: http.StatusForbidden}
	if err := newTestResource(t, f).patch(context.Background(), model(), `{}`, time.Minute); err == nil || errors.Is(err, errInstanceGone) {
		t.Errorf("403 must be a real error, got %v", err)
	}
}

func TestSchemaValid(t *testing.T) {
	r := NewResourceInstanceParametersResource()
	resp := &resource.SchemaResponse{}
	r.Schema(context.Background(), resource.SchemaRequest{}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatal(resp.Diagnostics)
	}
	if diags := resp.Schema.ValidateImplementation(context.Background()); diags.HasError() {
		t.Fatal(diags)
	}
}
