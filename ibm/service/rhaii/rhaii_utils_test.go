// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
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

	for _, plan := range []string{"instructlab-pricing-plan", "plan-1"} {
		id, err := rhaiiResolvePlanID(ctx, gc, plan)
		if err != nil || id != "plan-1" {
			t.Errorf("rhaiiResolvePlanID(%q) = %q, %v", plan, id, err)
		}
	}

	_, err := rhaiiResolvePlanID(ctx, gc, "missing")
	if err == nil || !strings.Contains(err.Error(), `"instructlab-pricing-plan" "other-plan"`) {
		t.Errorf("expected an error listing the valid plans, got %v", err)
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

func TestFlattenRhaiiInferenceModelMetadata(t *testing.T) {
	params := int64(70553706496)
	m := &rhaiiv1.InferenceModelMetadata{
		Name:        ptr("llama"),
		TotalParams: &params,
		Config:      map[string]interface{}{"model_type": "llama"},
		ModelCard:   ptr("# card"),
		Pricing:     &rhaiiv1.ModelPricing{InputMeasure: ptr("IN"), OutputMeasure: ptr("OUT")},
	}

	out, err := flattenRhaiiInferenceModelMetadata(m, false)
	if err != nil {
		t.Fatal(err)
	}
	if out["config_json"] != `{"model_type":"llama"}` || out["total_params"] != 70553706496 {
		t.Errorf("unexpected flatten result: %v", out)
	}
	if _, ok := out["model_card"]; ok {
		t.Error("model_card must only be set when requested")
	}
	if p := out["pricing"].([]interface{})[0].(map[string]interface{}); p["input_measure"] != "IN" {
		t.Errorf("unexpected pricing: %v", p)
	}

	out, _ = flattenRhaiiInferenceModelMetadata(m, true)
	if out["model_card"] != "# card" || out["source_json"] != "" {
		t.Errorf("unexpected flatten result with model card: %v", out)
	}

	if out, _ := flattenRhaiiInferenceModelMetadata(nil, true); len(out) != 0 {
		t.Errorf("nil metadata must flatten to an empty map, got %v", out)
	}
}

func TestRhaiiSchemas(t *testing.T) {
	r := ResourceIBMRhaiiProject()
	if err := r.InternalValidate(nil, true); err != nil {
		t.Fatalf("resource schema is not valid: %s", err)
	}
	s := r.Schema
	if !s["name"].Required || s["plan"].Default != rhaiiDefaultPlan || s["location"].Default != rhaiiDefaultLocation || !s["location"].ForceNew {
		t.Error("unexpected name, plan or location schema")
	}
	for _, k := range []string{"tags", "access_tags"} {
		if !s[k].Optional || !s[k].Computed {
			t.Errorf("%s must be optional and computed", k)
		}
	}
	for _, k := range []string{"service", "project_id", "endpoint", "crn", "guid"} {
		if s[k] == nil || s[k].Optional || !s[k].Computed {
			t.Errorf("%s must be computed only", k)
		}
	}

	if err := DataSourceIBMRhaiiProject().InternalValidate(nil, false); err != nil {
		t.Errorf("ibm_rhaii_project data source schema is not valid: %s", err)
	}
	if err := DataSourceIBMRhaiiInferenceModels().InternalValidate(nil, false); err != nil {
		t.Errorf("ibm_rhaii_inference_models data source schema is not valid: %s", err)
	}
	if err := DataSourceIBMRhaiiInferenceModel().InternalValidate(nil, false); err != nil {
		t.Errorf("ibm_rhaii_inference_model data source schema is not valid: %s", err)
	}
}
