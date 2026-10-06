// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaiiv1_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/rhaii/rhaiiv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

const projectID = "917bc95a-fef0-4039-b936-e0b6fb17b721"

func newTestClient(t *testing.T, handler http.HandlerFunc) *rhaiiv1.RhaiiV1 {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	client, err := rhaiiv1.NewRhaiiV1(&rhaiiv1.RhaiiV1Options{
		URL:           server.URL + "/v1",
		Authenticator: &core.BearerTokenAuthenticator{BearerToken: "token"},
	})
	if err != nil {
		t.Fatalf("NewRhaiiV1() error: %s", err)
	}
	return client
}

func TestListInferenceModels(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != fmt.Sprintf("/v1/projects/%s/inference/models", projectID) {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer token" {
			t.Errorf("Authorization header = %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"object":"list","data":[{"id":"llama-3-3-70b-instruct","object":"model","created":1723215000,"owned_by":"ibm",
			"custom_metadata":{"name":"llama-3-3-70b-instruct","state":"model_loaded","status":"loaded","total_params":70553706496,
			"config":{"model_type":"llama"},"pricing":{"input_measure":"IN","output_measure":"OUT"}}}]}`)
	})

	result, _, err := client.ListInferenceModels(client.NewListInferenceModelsOptions(projectID))
	if err != nil {
		t.Fatalf("ListInferenceModels() error: %s", err)
	}
	if len(result.Data) != 1 {
		t.Fatalf("got %d models, want 1", len(result.Data))
	}
	m := result.Data[0]
	if *m.ID != "llama-3-3-70b-instruct" || *m.OwnedBy != "ibm" || *m.Created != 1723215000 {
		t.Errorf("unexpected model: %+v", m)
	}
	if *m.CustomMetadata.TotalParams != 70553706496 || *m.CustomMetadata.Pricing.InputMeasure != "IN" || m.CustomMetadata.Config["model_type"] != "llama" {
		t.Errorf("unexpected metadata: %+v", m.CustomMetadata)
	}
}

func TestGetInferenceModel(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("/v1/projects/%s/inference/models/granite-4-1-30b", projectID) {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"identifier":"granite-4-1-30b","provider_id":"ibm","model_type":"llm","type":"model",
			"metadata":{"display_name":"Granite","model_card":"# card","validation":{"status":"validated"}}}`)
	})

	result, _, err := client.GetInferenceModel(client.NewGetInferenceModelOptions(projectID, "granite-4-1-30b"))
	if err != nil {
		t.Fatalf("GetInferenceModel() error: %s", err)
	}
	if *result.Identifier != "granite-4-1-30b" || *result.ModelType != "llm" {
		t.Errorf("unexpected model: %+v", result)
	}
	if *result.Metadata.ModelCard != "# card" || *result.Metadata.Validation.Status != "validated" {
		t.Errorf("unexpected metadata: %+v", result.Metadata)
	}
}

func TestGetInferenceModelError(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, `{"errors":[{"id":"EDKHR1","code":"forbidden","message":"Forbidden."}],"status_code":403}`)
	})

	result, resp, err := client.GetInferenceModel(client.NewGetInferenceModelOptions(projectID, "x"))
	if err == nil || result != nil {
		t.Fatalf("expected an error, got result %+v", result)
	}
	if resp == nil || resp.StatusCode != http.StatusForbidden {
		t.Fatalf("expected a 403 response, got %+v", resp)
	}
}

func TestOptionsValidation(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		t.Error("no request expected")
	})
	if _, _, err := client.ListInferenceModels(nil); err == nil {
		t.Error("expected an error for nil options")
	}
	if _, _, err := client.GetInferenceModel(client.NewGetInferenceModelOptions(projectID, "")); err == nil {
		t.Error("expected an error for an empty model")
	}
}

func TestServiceURLs(t *testing.T) {
	url, err := rhaiiv1.GetServiceURLForRegion("us-east")
	if err != nil || url != rhaiiv1.DefaultServiceURL {
		t.Errorf("GetServiceURLForRegion(us-east) = %q, %v", url, err)
	}
	if _, err := rhaiiv1.GetServiceURLForRegion("mars-1"); err == nil {
		t.Error("expected an error for an unknown region")
	}
	if got := rhaiiv1.ProjectServiceURL(rhaiiv1.DefaultServiceURL+"/", projectID); got != "https://us-east.rhai.ibm.com/v1/projects/"+projectID {
		t.Errorf("ProjectServiceURL() = %q", got)
	}
}

func TestClone(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {})
	clone := client.Clone()
	if err := clone.SetServiceURL("https://example.com/v1"); err != nil {
		t.Fatal(err)
	}
	if client.GetServiceURL() == clone.GetServiceURL() {
		t.Error("changing the clone URL must not change the original client")
	}
}
