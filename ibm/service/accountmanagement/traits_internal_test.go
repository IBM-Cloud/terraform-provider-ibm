// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package accountmanagement

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/accountmanagementv4"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func testAccountClient(t *testing.T, h http.HandlerFunc) *accountmanagementv4.AccountManagementV4 {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	c, err := accountmanagementv4.NewAccountManagementV4(&accountmanagementv4.AccountManagementV4Options{
		URL: srv.URL, Authenticator: &core.NoAuthAuthenticator{},
	})
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestPatchTraitsSendsOnlyManagedTraits(t *testing.T) {
	var gotMethod, gotPath string
	var body map[string]any
	c := testAccountClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath = r.Method, r.URL.Path
		b, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(b, &body)
		w.WriteHeader(http.StatusNoContent)
	})
	f := false
	if err := patchTraits(context.Background(), c, "acct1", &accountTraits{TeamDirectoryEnabled: &f}); err != nil {
		t.Fatal(err)
	}
	if gotMethod != http.MethodPatch || gotPath != "/v1/accounts/acct1/traits" {
		t.Errorf("method=%s path=%s", gotMethod, gotPath)
	}
	if _, ok := body["fs_ready"]; ok || body["team_directory_enabled"] != false {
		t.Errorf("body = %v, want only team_directory_enabled=false", body)
	}
}

func TestGetTraits(t *testing.T) {
	c := testAccountClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"fs_ready":true,"team_directory_enabled":false,"eu_supported":false}`))
	})
	traits, _, err := getTraits(context.Background(), c, "acct1")
	if err != nil || traits.FsReady == nil || !*traits.FsReady || traits.TeamDirectoryEnabled == nil || *traits.TeamDirectoryEnabled {
		t.Fatalf("traits = %+v, err = %v", traits, err)
	}
}

func TestApplyManagedTraits(t *testing.T) {
	tr, fa := true, false
	m := accountTraitsModel{FsReady: types.BoolValue(true), TeamDirectoryEnabled: types.BoolNull()}
	applyManagedTraits(&m, &accountTraits{FsReady: &fa, TeamDirectoryEnabled: &tr})
	if m.FsReady.ValueBool() != false {
		t.Error("managed trait changed outside Terraform must show as drift")
	}
	if !m.TeamDirectoryEnabled.IsNull() {
		t.Error("unmanaged trait must stay null")
	}
}

func TestForEndpointAndAccountID(t *testing.T) {
	base, _ := accountmanagementv4.NewAccountManagementV4(&accountmanagementv4.AccountManagementV4Options{
		URL: "https://accounts.cloud.ibm.com", Authenticator: &core.NoAuthAuthenticator{},
	})
	tc := &traitsClient{client: base, account: "provider-acct"}
	c, err := tc.forEndpoint(types.StringValue("private.us-south.accounts.cloud.ibm.com"))
	if err != nil || c.Service.GetServiceURL() != "https://private.us-south.accounts.cloud.ibm.com" {
		t.Errorf("forEndpoint = %v, %v", c.Service.GetServiceURL(), err)
	}
	if base.Service.GetServiceURL() != "https://accounts.cloud.ibm.com" {
		t.Error("forEndpoint must not change the shared client")
	}
	if id, _ := tc.accountID(types.StringNull()); id != "provider-acct" {
		t.Errorf("accountID default = %q", id)
	}
	if id, _ := tc.accountID(types.StringValue("other")); id != "other" {
		t.Errorf("accountID configured = %q", id)
	}
	if _, err := (&traitsClient{client: base}).accountID(types.StringNull()); err == nil {
		t.Error("expected error when the account is unknown")
	}
}
