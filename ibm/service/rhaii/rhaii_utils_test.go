// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/rhaii/rhaiiv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

const testProjectID = "917bc95a-fef0-4039-b936-e0b6fb17b721"

func newTestCatalog(t *testing.T) *globalcatalogv1.GlobalCatalogV1 {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/instructlab/plan":
			fmt.Fprint(w, `{"resources":[
				{"id":"plan-1","name":"instructlab-pricing-plan","kind":"plan","disabled":false,"images":{},"overview_ui":{},"provider":{"name":"ibm","email":"x@ibm.com"},"tags":[]},
				{"id":"plan-2","name":"other-plan","kind":"plan","disabled":false,"images":{},"overview_ui":{},"provider":{"name":"ibm","email":"x@ibm.com"},"tags":[]}]}`)
		case "/plan-1/deployment":
			fmt.Fprint(w, `{"resources":[
				{"id":"dep-1","name":"us-east","kind":"deployment","catalog_crn":"crn:v1:bluemix:public:globalcatalog::::deployment:plan-1%3Aus-east71658",
				 "metadata":{"rc_compatible":true,"deployment":{"location":"us-east"}},"disabled":false,"images":{},"overview_ui":{},"provider":{"name":"ibm","email":"x@ibm.com"},"tags":[]},
				{"id":"dep-2","name":"eu-de","kind":"deployment","catalog_crn":"crn:eu-de",
				 "metadata":{"rc_compatible":false,"deployment":{"location":"eu-de"}},"disabled":false,"images":{},"overview_ui":{},"provider":{"name":"ibm","email":"x@ibm.com"},"tags":[]}]}`)
		case "/plan-1":
			fmt.Fprint(w, `{"id":"plan-1","name":"instructlab-pricing-plan","kind":"plan","disabled":false,"images":{},"overview_ui":{},"provider":{"name":"ibm","email":"x@ibm.com"},"tags":[]}`)
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, `{}`)
		}
	}))
	t.Cleanup(server.Close)

	client, err := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
		URL:           server.URL,
		Authenticator: &core.NoAuthAuthenticator{},
	})
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func TestRhaiiResolvePlanID(t *testing.T) {
	gc := newTestCatalog(t)
	ctx := context.Background()

	id, err := rhaiiResolvePlanID(ctx, gc, "instructlab-pricing-plan")
	if err != nil || id != "plan-1" {
		t.Errorf("rhaiiResolvePlanID(name) = %q, %v", id, err)
	}

	// A plan ID is not accepted as a plan name; the error names the plan to use.
	_, err = rhaiiResolvePlanID(ctx, gc, "plan-1")
	if err == nil || !strings.Contains(err.Error(), `is a plan ID; set the plan name "instructlab-pricing-plan"`) {
		t.Errorf("expected a plan ID error, got %v", err)
	}

	_, err = rhaiiResolvePlanID(ctx, gc, "missing")
	if err == nil || !strings.Contains(err.Error(), `"instructlab-pricing-plan" "other-plan"`) {
		t.Errorf("expected an error listing the valid plan names, got %v", err)
	}
}

func TestRhaiiResolveTargetCRN(t *testing.T) {
	gc := newTestCatalog(t)
	ctx := context.Background()

	crn, err := rhaiiResolveTargetCRN(ctx, gc, "plan-1", "us-east")
	if err != nil || crn != "crn:v1:bluemix:public:globalcatalog::::deployment:plan-1%3Aus-east71658" {
		t.Errorf("rhaiiResolveTargetCRN(us-east) = %q, %v", crn, err)
	}

	// eu-de exists but is not resource controller compatible, so it is neither used nor listed.
	_, err = rhaiiResolveTargetCRN(ctx, gc, "plan-1", "eu-de")
	if err == nil || !strings.Contains(err.Error(), `Valid locations are: ["us-east"]`) {
		t.Errorf("expected an error listing us-east only, got %v", err)
	}
}

func TestRhaiiPlanName(t *testing.T) {
	name, err := rhaiiPlanName(context.Background(), newTestCatalog(t), "plan-1")
	if err != nil || name != "instructlab-pricing-plan" {
		t.Errorf("rhaiiPlanName() = %q, %v", name, err)
	}
}

func TestRhaiiFindProjectByName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("name") != "proj" || q.Get("resource_id") != "instructlab" {
			t.Errorf("unexpected query %s", r.URL.RawQuery)
		}
		w.Header().Set("Content-Type", "application/json")
		if q.Get("start") == "" {
			fmt.Fprint(w, `{"rows_count":2,"next_url":"/v2/resource_instances?start=page2&name=proj","resources":[
				{"id":"crn-a","guid":"a","name":"proj","region_id":"us-east","resource_id":"instructlab","state":"active"},
				{"id":"crn-old","guid":"old","name":"proj","region_id":"us-east","resource_id":"instructlab","state":"pending_reclamation"}]}`)
			return
		}
		fmt.Fprint(w, `{"rows_count":1,"next_url":null,"resources":[
			{"id":"crn-b","guid":"b","name":"proj","region_id":"eu-de","resource_id":"instructlab","state":"active"}]}`)
	}))
	defer server.Close()

	rcClient, err := rc.NewResourceControllerV2(&rc.ResourceControllerV2Options{URL: server.URL, Authenticator: &core.NoAuthAuthenticator{}})
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	// Two live matches across both pages, so the lookup is ambiguous.
	if _, err := rhaiiFindProjectByName(ctx, rcClient, "proj", "", ""); err == nil || !strings.Contains(err.Error(), "found 2") {
		t.Errorf("expected an ambiguity error, got %v", err)
	}

	// The location filter makes it unique; the pending_reclamation instance is ignored.
	instance, err := rhaiiFindProjectByName(ctx, rcClient, "proj", "", "us-east")
	if err != nil || *instance.GUID != "a" {
		t.Errorf("rhaiiFindProjectByName(us-east) = %+v, %v", instance, err)
	}

	if _, err := rhaiiFindProjectByName(ctx, rcClient, "proj", "", "jp-tok"); err == nil || !strings.Contains(err.Error(), "no Red Hat AI Inference project") {
		t.Errorf("expected a not found error, got %v", err)
	}
}

func TestRhaiiProjectEndpoint(t *testing.T) {
	if got, want := rhaiiProjectEndpoint("us-east", testProjectID), "https://us-east.rhai.ibm.com/v1/projects/"+testProjectID; got != want {
		t.Errorf("rhaiiProjectEndpoint(us-east) = %q, want %q", got, want)
	}
	if got, want := rhaiiProjectEndpoint("eu-de", testProjectID), "https://eu-de.rhai.ibm.com/v1/projects/"+testProjectID; got != want {
		t.Errorf("rhaiiProjectEndpoint(eu-de) = %q, want %q", got, want)
	}
	if got, want := rhaiiProjectPrivateEndpoint("us-east", testProjectID), "https://private.us-east.rhai.ibm.com/v1/projects/"+testProjectID; got != want {
		t.Errorf("rhaiiProjectPrivateEndpoint(us-east) = %q, want %q", got, want)
	}
	if got, want := rhaiiProjectPrivateEndpoint("eu-de", testProjectID), "https://private.eu-de.rhai.ibm.com/v1/projects/"+testProjectID; got != want {
		t.Errorf("rhaiiProjectPrivateEndpoint(eu-de) = %q, want %q", got, want)
	}
}

func TestRhaiiRefreshFunctions(t *testing.T) {
	inst := func(state, opState string) *rc.ResourceInstance {
		i := &rc.ResourceInstance{ID: ptr("crn"), State: ptr(state)}
		if opState != "" {
			i.LastOperation = &rc.ResourceInstanceLastOperation{State: ptr(opState), Description: ptr("desc")}
		}
		return i
	}

	cases := []struct {
		name    string
		refresh func(*rc.ResourceInstance) (string, error)
		in      *rc.ResourceInstance
		want    string
		wantErr bool
	}{
		{"create provisioning", rhaiiCreateRefresh, inst("provisioning", ""), rhaiiWaitCreating, false},
		{"create unknown state", rhaiiCreateRefresh, inst("something_new", ""), rhaiiWaitCreating, false},
		{"create active", rhaiiCreateRefresh, inst("active", ""), rhaiiWaitDone, false},
		{"create failed", rhaiiCreateRefresh, inst("failed", ""), "", true},
		{"update in progress", rhaiiUpdateRefresh, inst("active", "in progress"), rhaiiWaitUpdating, false},
		{"update succeeded", rhaiiUpdateRefresh, inst("active", "succeeded"), rhaiiWaitDone, false},
		{"update failed", rhaiiUpdateRefresh, inst("active", "failed"), "", true},
		{"update no operation", rhaiiUpdateRefresh, inst("active", ""), rhaiiWaitDone, false},
		{"delete active", rhaiiDeleteRefresh, inst("active", ""), rhaiiWaitDeleting, false},
		{"delete removed", rhaiiDeleteRefresh, inst("removed", ""), rhaiiWaitDone, false},
		{"delete pending reclamation", rhaiiDeleteRefresh, inst("pending_reclamation", ""), rhaiiWaitDone, false},
		{"delete failed", rhaiiDeleteRefresh, inst("failed", ""), "", true},
	}
	for _, c := range cases {
		got, err := c.refresh(c.in)
		if (err != nil) != c.wantErr || got != c.want {
			t.Errorf("%s: got %q, %v; want %q, error %v", c.name, got, err, c.want, c.wantErr)
		}
	}
}

func TestRhaiiNextStart(t *testing.T) {
	start, err := rhaiiNextStart(ptr("/v2/resource_instances?start=abc&limit=100"))
	if err != nil || start != "abc" {
		t.Errorf("rhaiiNextStart() = %q, %v", start, err)
	}
	if start, _ := rhaiiNextStart(nil); start != "" {
		t.Errorf("rhaiiNextStart(nil) = %q", start)
	}
}

func TestRhaiiServiceURLForLocation(t *testing.T) {
	t.Setenv("IBMCLOUD_ENDPOINTS_FILE_PATH", "")
	t.Setenv("IC_ENDPOINTS_FILE_PATH", "")
	t.Setenv(rhaiiEndpointEnv, "")

	dir := t.TempDir()
	endpointsFile := dir + "/endpoints.json"
	if err := os.WriteFile(endpointsFile, []byte(`{"IBMCLOUD_RHAII_API_ENDPOINT":{"private":{"us-east":"https://custom-private.example.com/v1"}}}`), 0o600); err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		name, visibility, file, want string
		wantErr                      bool
	}{
		{"public", "public", "", "https://us-east.rhai.ibm.com/v1", false},
		{"private", "private", "", "https://private.us-east.rhai.ibm.com/v1", false},
		{"public and private", "public-and-private", "", "https://private.us-east.rhai.ibm.com/v1", false},
		{"private from endpoints file", "private", endpointsFile, "https://custom-private.example.com/v1", false},
		{"public not in endpoints file", "public", endpointsFile, "https://us-east.rhai.ibm.com/v1", false},
		{"public and private ignores endpoints file", "public-and-private", endpointsFile, "https://private.us-east.rhai.ibm.com/v1", false},
	}
	for _, c := range cases {
		got, err := rhaiiServiceURLForLocation("us-east", c.visibility, c.file)
		if (err != nil) != c.wantErr || got != c.want {
			t.Errorf("%s: got %q, %v; want %q", c.name, got, err, c.want)
		}
	}

	// A private request in a region without a private endpoint fails instead of going public.
	if _, err := rhaiiServiceURLForLocation("eu-de", "private", ""); err == nil {
		t.Error("expected an error for a private endpoint in an unknown region")
	}

	// The environment variable wins over everything else.
	t.Setenv(rhaiiEndpointEnv, "https://override.example.com/v1")
	if got, _ := rhaiiServiceURLForLocation("us-east", "private", endpointsFile); got != "https://override.example.com/v1" {
		t.Errorf("environment override not used, got %q", got)
	}
}

func TestFlattenRhaiiInferenceModelMetadata(t *testing.T) {
	params := int64(70553706496)
	m, diags := flattenRhaiiInferenceModelMetadata(&rhaiiv1.InferenceModelMetadata{
		Name:        ptr("llama"),
		TotalParams: &params,
		Config:      map[string]interface{}{"model_type": "llama"},
		Pricing:     &rhaiiv1.ModelPricing{InputMeasure: ptr("IN"), OutputMeasure: ptr("OUT")},
	})
	if diags.HasError() {
		t.Fatal(diags)
	}
	if m.ConfigJSON.ValueString() != `{"model_type":"llama"}` || m.TotalParams.ValueInt64() != 70553706496 || m.Name.ValueString() != "llama" {
		t.Errorf("unexpected flatten result: %+v", m)
	}
	if len(m.Pricing.Elements()) != 1 || len(m.Validation.Elements()) != 0 {
		t.Errorf("unexpected pricing or validation: %v %v", m.Pricing, m.Validation)
	}

	empty, diags := flattenRhaiiInferenceModelMetadata(nil)
	if diags.HasError() || !empty.TotalParams.IsNull() || len(empty.Pricing.Elements()) != 0 {
		t.Errorf("nil metadata must flatten to empty values, got %+v", empty)
	}
}

func TestRhaiiWithoutEnvTags(t *testing.T) {
	t.Setenv("IC_ENV_TAGS", "Schematics:abc,env:schematics")
	got := rhaiiWithoutEnvTags([]string{"Schematics:abc", "rhaii-da", "env:schematics", "test"})
	if !reflect.DeepEqual(got, []string{"rhaii-da", "test"}) {
		t.Errorf("rhaiiWithoutEnvTags() = %v", got)
	}

	t.Setenv("IC_ENV_TAGS", "")
	got = rhaiiWithoutEnvTags([]string{"Schematics:abc", "rhaii-da"})
	if !reflect.DeepEqual(got, []string{"Schematics:abc", "rhaii-da"}) {
		t.Errorf("rhaiiWithoutEnvTags() without IC_ENV_TAGS = %v", got)
	}
}
