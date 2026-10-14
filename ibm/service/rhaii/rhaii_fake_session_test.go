// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/rhaii/rhaiiv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/globalsearchv2"
	"github.com/IBM/platform-services-go-sdk/globaltaggingv1"
)

// fakeSession is a conns.ClientSession for unit tests. It embeds the interface
// (nil) and overrides only the clients the tests use; any other call panics,
// which shows a test reaching an unexpected client.
type fakeSession struct {
	conns.ClientSession
	tagging *globaltaggingv1.GlobalTaggingV1
	search  *globalsearchv2.GlobalSearchV2
	rhaii   *rhaiiv1.RhaiiV1
}

func (f fakeSession) GlobalSearchAPIV2() (globalsearchv2.GlobalSearchV2, error) {
	return *f.search, nil
}

func (f fakeSession) GlobalTaggingAPIv1() (globaltaggingv1.GlobalTaggingV1, error) {
	return *f.tagging, nil
}

func (f fakeSession) BluemixUserDetails() (*conns.UserConfig, error) {
	return &conns.UserConfig{UserAccount: "acc"}, nil
}

func (f fakeSession) RhaiiV1() (*rhaiiv1.RhaiiV1, error) {
	return f.rhaii.Clone(), nil
}

// recordedCall is one request received by the mock Global Tagging server.
type recordedCall struct {
	Method, Path, TagType string
	TagNames              []string
}

// newTaggingServer returns Global Tagging and Global Search clients backed by a
// mock server. It records every tagging call, answers attach and detach with
// success, and answers search (used by flex to wait until the attached tags are
// visible) with searchTags as both the user and the access tags of the CRN.
func newTaggingServer(t *testing.T, searchTags []string) (*globaltaggingv1.GlobalTaggingV1, *globalsearchv2.GlobalSearchV2, func() []recordedCall) {
	t.Helper()
	var mu sync.Mutex
	var calls []recordedCall
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/resources/search") {
			w.Header().Set("Content-Type", "application/json")
			b, _ := json.Marshal(map[string]interface{}{"items": []map[string]interface{}{{"crn": "crn:v1:project", "tags": searchTags, "access_tags": searchTags}}})
			w.Write(b)
			return
		}
		body, _ := io.ReadAll(r.Body)
		var req struct {
			TagNames []string `json:"tag_names"`
		}
		_ = json.Unmarshal(body, &req)
		mu.Lock()
		calls = append(calls, recordedCall{Method: r.Method, Path: r.URL.Path, TagType: r.URL.Query().Get("tag_type"), TagNames: req.TagNames})
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"results":[]}`)
	}))
	t.Cleanup(server.Close)

	client, err := globaltaggingv1.NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{URL: server.URL, Authenticator: &core.NoAuthAuthenticator{}})
	if err != nil {
		t.Fatal(err)
	}
	search, err := globalsearchv2.NewGlobalSearchV2(&globalsearchv2.GlobalSearchV2Options{URL: server.URL, Authenticator: &core.NoAuthAuthenticator{}})
	if err != nil {
		t.Fatal(err)
	}
	return client, search, func() []recordedCall {
		mu.Lock()
		defer mu.Unlock()
		return append([]recordedCall(nil), calls...)
	}
}

// TestRhaiiUpdateTagsDoesNotDeleteTags checks that changing tags only attaches
// and detaches them on the project. flex.UpdateTagsUsingCRN, used before, also
// called DELETE /v3/tags/{name}, which deletes the tag from the whole account.
func TestRhaiiUpdateTagsDoesNotDeleteTags(t *testing.T) {
	t.Setenv("IC_ENV_TAGS", "")
	for _, tagType := range []string{rhaiiUserTagType, rhaiiAccessTagType} {
		client, search, calls := newTaggingServer(t, []string{"keep:1", "new:1"})
		meta := fakeSession{tagging: client, search: search}

		err := rhaiiUpdateTags(meta, "crn:v1:project", tagType, []string{"keep:1", "drop:1"}, []string{"keep:1", "new:1"})
		if err != nil {
			t.Fatalf("%s: rhaiiUpdateTags() error: %s", tagType, err)
		}

		var attached, detached []string
		for _, c := range calls() {
			switch {
			case c.Method == http.MethodDelete:
				t.Errorf("%s: unexpected tag delete call %s %s", tagType, c.Method, c.Path)
			case strings.HasSuffix(c.Path, "/v3/tags/attach"):
				attached = append(attached, c.TagNames...)
				if c.TagType != tagType {
					t.Errorf("%s: attach tag_type = %q", tagType, c.TagType)
				}
			case strings.HasSuffix(c.Path, "/v3/tags/detach"):
				detached = append(detached, c.TagNames...)
				if c.TagType != tagType {
					t.Errorf("%s: detach tag_type = %q", tagType, c.TagType)
				}
			default:
				t.Errorf("%s: unexpected call %s %s", tagType, c.Method, c.Path)
			}
		}
		if strings.Join(attached, ",") != "new:1" || strings.Join(detached, ",") != "drop:1" {
			t.Errorf("%s: attached %v, detached %v; want [new:1] and [drop:1]", tagType, attached, detached)
		}
	}
}
