// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/rhaii/rhaiiv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

const testModelProject = "917bc95a-fef0-4039-b936-e0b6fb17b721"

// newModelServer returns an RHAII client backed by a mock server that serves
// the list and get model responses of the API spec examples.
func newModelServer(t *testing.T) *rhaiiv1.RhaiiV1 {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/v1/projects/" + testModelProject + "/inference/models":
			fmt.Fprint(w, `{"object":"list","data":[
				{"id":"llama-3-3-70b-instruct","object":"model","created":1723215000,"owned_by":"ibm",
				 "custom_metadata":{"name":"llama-3-3-70b-instruct","display_name":"llama-3-3-70b-instruct","state":"model_loaded","status":"loaded",
				   "total_params":70553706496,"config":{"max_position_embeddings":131072},
				   "pricing":{"input_measure":"LLAMA-IN","output_measure":"LLAMA-OUT"}}},
				{"id":"granite-4-1-30b","object":"model","created":1750000000,"owned_by":"ibm"}]}`)
		case "/v1/projects/" + testModelProject + "/inference/models/llama-3-3-70b-instruct":
			fmt.Fprint(w, `{"identifier":"llama-3-3-70b-instruct","provider_id":"ibm","provider_resource_id":"llama-3-3-70b-instruct",
				"type":"model","model_type":"llm",
				"metadata":{"name":"llama-3-3-70b-instruct","state":"model_loaded","status":"loaded","total_params":70553706496,
				  "config":{"max_position_embeddings":131072},"model_card":"# Llama",
				  "pricing":{"input_measure":"LLAMA-IN","output_measure":"LLAMA-OUT"},
				  "validation":{"status":"validated"},
				  "source":{"huggingface":{"image":"meta-llama/Llama-3.3-70B-Instruct"}}}}`)
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, `{"errors":[{"id":"E00404","code":"invalid_request","message":"Not found."}],"status_code":404}`)
		}
	}))
	t.Cleanup(server.Close)

	client, err := rhaiiv1.NewRhaiiV1(&rhaiiv1.RhaiiV1Options{URL: server.URL + "/v1", Authenticator: &core.NoAuthAuthenticator{}})
	if err != nil {
		t.Fatal(err)
	}
	return client
}

// readDataSource runs Configure and Read of a data source with the given
// configuration and returns the state, or the diagnostics on error.
func readDataSource(t *testing.T, ds datasource.DataSource, meta fakeSession, config map[string]tftypes.Value) (tftypes.Value, datasource.ReadResponse) {
	t.Helper()
	ctx := context.Background()

	var schemaResp datasource.SchemaResponse
	ds.Schema(ctx, datasource.SchemaRequest{}, &schemaResp)
	typ := schemaResp.Schema.Type().TerraformType(ctx).(tftypes.Object)

	attrs := map[string]tftypes.Value{}
	for name, attrType := range typ.AttributeTypes {
		if v, ok := config[name]; ok {
			attrs[name] = v
		} else {
			attrs[name] = tftypes.NewValue(attrType, nil)
		}
	}
	raw := tftypes.NewValue(typ, attrs)

	if c, ok := ds.(datasource.DataSourceWithConfigure); ok {
		var cr datasource.ConfigureResponse
		c.Configure(ctx, datasource.ConfigureRequest{ProviderData: meta}, &cr)
		if cr.Diagnostics.HasError() {
			t.Fatalf("Configure: %v", cr.Diagnostics)
		}
	}
	resp := datasource.ReadResponse{State: tfsdk.State{Schema: schemaResp.Schema, Raw: tftypes.NewValue(typ, nil)}}
	ds.Read(ctx, datasource.ReadRequest{Config: tfsdk.Config{Schema: schemaResp.Schema, Raw: raw}}, &resp)
	return resp.State.Raw, resp
}

// attrOf returns a top level attribute of an object value.
func attrOf(t *testing.T, v tftypes.Value, name string) tftypes.Value {
	t.Helper()
	m := map[string]tftypes.Value{}
	if err := v.As(&m); err != nil {
		t.Fatal(err)
	}
	return m[name]
}

func asString(t *testing.T, v tftypes.Value) string {
	t.Helper()
	var s string
	if err := v.As(&s); err != nil {
		t.Fatal(err)
	}
	return s
}

func asInt(t *testing.T, v tftypes.Value) int64 {
	t.Helper()
	n := new(big.Float)
	if err := v.As(&n); err != nil {
		t.Fatal(err)
	}
	i, _ := n.Int64()
	return i
}

func asList(t *testing.T, v tftypes.Value) []tftypes.Value {
	t.Helper()
	var l []tftypes.Value
	if err := v.As(&l); err != nil {
		t.Fatal(err)
	}
	return l
}

func TestRhaiiInferenceModelsDataSourceRead(t *testing.T) {
	meta := fakeSession{rhaii: newModelServer(t)}
	state, resp := readDataSource(t, NewRhaiiInferenceModelsDataSource(), meta, map[string]tftypes.Value{
		"project_id": tftypes.NewValue(tftypes.String, testModelProject),
	})
	if resp.Diagnostics.HasError() {
		t.Fatalf("Read: %v", resp.Diagnostics)
	}

	if got := asString(t, attrOf(t, state, "id")); got != testModelProject {
		t.Errorf("id = %q", got)
	}
	models := asList(t, attrOf(t, state, "models"))
	if len(models) != 2 {
		t.Fatalf("got %d models, want 2", len(models))
	}
	llama := models[0]
	if got := asString(t, attrOf(t, llama, "id")); got != "llama-3-3-70b-instruct" {
		t.Errorf("models.0.id = %q", got)
	}
	if got := asInt(t, attrOf(t, llama, "total_params")); got != 70553706496 {
		t.Errorf("models.0.total_params = %d", got)
	}
	if got := asInt(t, attrOf(t, llama, "created")); got != 1723215000 {
		t.Errorf("models.0.created = %d", got)
	}
	if got := asString(t, attrOf(t, llama, "config_json")); got != `{"max_position_embeddings":131072}` {
		t.Errorf("models.0.config_json = %q", got)
	}
	pricing := asList(t, attrOf(t, llama, "pricing"))
	if len(pricing) != 1 || asString(t, attrOf(t, pricing[0], "input_measure")) != "LLAMA-IN" {
		t.Errorf("models.0.pricing = %v", pricing)
	}

	// A model without custom_metadata gives empty values, not an error.
	granite := models[1]
	if !attrOf(t, granite, "total_params").IsNull() || len(asList(t, attrOf(t, granite, "pricing"))) != 0 {
		t.Errorf("models.1 without metadata: %v", granite)
	}
}

func TestRhaiiInferenceModelDataSourceRead(t *testing.T) {
	meta := fakeSession{rhaii: newModelServer(t)}
	state, resp := readDataSource(t, NewRhaiiInferenceModelDataSource(), meta, map[string]tftypes.Value{
		"project_id": tftypes.NewValue(tftypes.String, testModelProject),
		"model":      tftypes.NewValue(tftypes.String, "llama-3-3-70b-instruct"),
	})
	if resp.Diagnostics.HasError() {
		t.Fatalf("Read: %v", resp.Diagnostics)
	}

	want := map[string]string{
		"id":                   testModelProject + "/llama-3-3-70b-instruct",
		"identifier":           "llama-3-3-70b-instruct",
		"provider_id":          "ibm",
		"provider_resource_id": "llama-3-3-70b-instruct",
		"type":                 "model",
		"model_type":           "llm",
		"status":               "loaded",
		"model_card":           "# Llama",
		"config_json":          `{"max_position_embeddings":131072}`,
		"source_json":          `{"huggingface":{"image":"meta-llama/Llama-3.3-70B-Instruct"}}`,
	}
	for k, v := range want {
		if got := asString(t, attrOf(t, state, k)); got != v {
			t.Errorf("%s = %q, want %q", k, got, v)
		}
	}
	if got := asInt(t, attrOf(t, state, "total_params")); got != 70553706496 {
		t.Errorf("total_params = %d", got)
	}
	validation := asList(t, attrOf(t, state, "validation"))
	if len(validation) != 1 || asString(t, attrOf(t, validation[0], "status")) != "validated" {
		t.Errorf("validation = %v", validation)
	}
}

func TestRhaiiInferenceModelDataSourceNotFound(t *testing.T) {
	meta := fakeSession{rhaii: newModelServer(t)}
	_, resp := readDataSource(t, NewRhaiiInferenceModelDataSource(), meta, map[string]tftypes.Value{
		"project_id": tftypes.NewValue(tftypes.String, testModelProject),
		"model":      tftypes.NewValue(tftypes.String, "missing-model"),
	})
	if !resp.Diagnostics.HasError() {
		t.Fatal("expected an error for a model that does not exist")
	}
}
