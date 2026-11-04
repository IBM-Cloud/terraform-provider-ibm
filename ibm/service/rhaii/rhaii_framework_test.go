// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii_test

import (
	"context"
	"sort"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/rhaii"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// TestRhaiiFrameworkSchemas validates the framework schemas of the resource and data sources.
func TestRhaiiFrameworkSchemas(t *testing.T) {
	ctx := context.Background()

	var rs resource.SchemaResponse
	rhaii.NewRhaiiProjectResource().Schema(ctx, resource.SchemaRequest{}, &rs)
	if rs.Diagnostics.HasError() {
		t.Fatalf("resource schema: %v", rs.Diagnostics)
	}
	if d := rs.Schema.ValidateImplementation(ctx); d.HasError() {
		t.Fatalf("resource schema implementation: %v", d)
	}

	for name, ds := range map[string]datasource.DataSource{
		"ibm_rhaii_project":          rhaii.NewRhaiiProjectDataSource(),
		"ibm_rhaii_inference_models": rhaii.NewRhaiiInferenceModelsDataSource(),
		"ibm_rhaii_inference_model":  rhaii.NewRhaiiInferenceModelDataSource(),
	} {
		var resp datasource.SchemaResponse
		ds.Schema(ctx, datasource.SchemaRequest{}, &resp)
		if resp.Diagnostics.HasError() {
			t.Fatalf("%s schema: %v", name, resp.Diagnostics)
		}
		if d := resp.Schema.ValidateImplementation(ctx); d.HasError() {
			t.Fatalf("%s schema implementation: %v", name, d)
		}
	}
}

// muxServer returns the production mux (SDKv2 upgraded to protocol 6 plus the
// framework provider), as built by the acceptance test factories and main.go.
func muxServer(t *testing.T) tfprotov6.ProviderServer {
	t.Helper()
	server, err := acc.TestAccProtoV6ProviderFactories()["ibm"]()
	if err != nil {
		t.Fatalf("creating the mux server: %s", err)
	}
	return server
}

// TestRhaiiMuxRegistration checks that the mux server starts (a type name in
// both providers would fail here) and serves the RHAII types from the framework.
func TestRhaiiMuxRegistration(t *testing.T) {
	resp, err := muxServer(t).GetProviderSchema(context.Background(), &tfprotov6.GetProviderSchemaRequest{})
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range resp.Diagnostics {
		if d.Severity == tfprotov6.DiagnosticSeverityError {
			t.Fatalf("GetProviderSchema: %s: %s", d.Summary, d.Detail)
		}
	}
	if resp.ResourceSchemas["ibm_rhaii_project"] == nil {
		t.Error("resource ibm_rhaii_project not served")
	}
	for _, n := range []string{"ibm_rhaii_project", "ibm_rhaii_inference_models", "ibm_rhaii_inference_model"} {
		if resp.DataSourceSchemas[n] == nil {
			t.Errorf("data source %s not served", n)
		}
	}
	// A type that stays in SDKv2 is still served by the mux.
	if resp.ResourceSchemas["ibm_resource_instance"] == nil {
		t.Error("SDKv2 resource ibm_resource_instance not served")
	}
}

// objectValue builds a value of the resource type with the given attributes; the
// other attributes are null. Strings are strings, []string are sets of strings.
func objectValue(t *testing.T, typ tftypes.Object, values map[string]interface{}) tftypes.Value {
	t.Helper()
	attrs := map[string]tftypes.Value{}
	for name, attrType := range typ.AttributeTypes {
		v, ok := values[name]
		if !ok {
			attrs[name] = tftypes.NewValue(attrType, nil)
			continue
		}
		switch x := v.(type) {
		case string:
			attrs[name] = tftypes.NewValue(tftypes.String, x)
		case []string:
			elems := []tftypes.Value{}
			for _, e := range x {
				elems = append(elems, tftypes.NewValue(tftypes.String, e))
			}
			attrs[name] = tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, elems)
		default:
			t.Fatalf("unsupported test value for %s: %T", name, v)
		}
	}
	return tftypes.NewValue(typ, attrs)
}

func dynamic(t *testing.T, typ tftypes.Type, v tftypes.Value) *tfprotov6.DynamicValue {
	t.Helper()
	dv, err := tfprotov6.NewDynamicValue(typ, v)
	if err != nil {
		t.Fatal(err)
	}
	return &dv
}

type planResult struct {
	values          map[string]tftypes.Value
	requiresReplace map[string]bool
}

// plan runs PlanResourceChange for ibm_rhaii_project through the mux server.
// The provider is not configured, so ModifyPlan skips the access tag API check;
// defaults, plan modifiers and the IC_ENV_TAGS handling all run.
func plan(t *testing.T, prior, config map[string]interface{}) planResult {
	t.Helper()
	ctx := context.Background()
	server := muxServer(t)
	schemaResp, err := server.GetProviderSchema(ctx, &tfprotov6.GetProviderSchemaRequest{})
	if err != nil {
		t.Fatal(err)
	}
	typ := schemaResp.ResourceSchemas["ibm_rhaii_project"].ValueType().(tftypes.Object)

	priorValue := tftypes.NewValue(typ, nil)
	if prior != nil {
		priorValue = objectValue(t, typ, prior)
	}
	configValue := objectValue(t, typ, config)

	resp, err := server.PlanResourceChange(ctx, &tfprotov6.PlanResourceChangeRequest{
		TypeName:         "ibm_rhaii_project",
		PriorState:       dynamic(t, typ, priorValue),
		ProposedNewState: dynamic(t, typ, configValue),
		Config:           dynamic(t, typ, configValue),
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range resp.Diagnostics {
		if d.Severity == tfprotov6.DiagnosticSeverityError {
			t.Fatalf("PlanResourceChange: %s: %s", d.Summary, d.Detail)
		}
	}
	planned, err := resp.PlannedState.Unmarshal(typ)
	if err != nil {
		t.Fatal(err)
	}
	values := map[string]tftypes.Value{}
	if err := planned.As(&values); err != nil {
		t.Fatal(err)
	}
	out := planResult{values: values, requiresReplace: map[string]bool{}}
	for _, p := range resp.RequiresReplace {
		out.requiresReplace[p.String()] = true
	}
	return out
}

func setStrings(t *testing.T, v tftypes.Value) []string {
	t.Helper()
	if !v.IsKnown() {
		return []string{"<unknown>"}
	}
	if v.IsNull() {
		return nil
	}
	var elems []tftypes.Value
	if err := v.As(&elems); err != nil {
		t.Fatal(err)
	}
	out := []string{}
	for _, e := range elems {
		var s string
		if err := e.As(&s); err != nil {
			t.Fatal(err)
		}
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}

func str(t *testing.T, v tftypes.Value) string {
	t.Helper()
	if !v.IsKnown() {
		return "<unknown>"
	}
	var s string
	if err := v.As(&s); err != nil {
		t.Fatal(err)
	}
	return s
}

var existingProject = map[string]interface{}{
	"id":                "crn:v1:bluemix:public:instructlab:us-east:a/acc:guid::",
	"crn":               "crn:v1:bluemix:public:instructlab:us-east:a/acc:guid::",
	"project_id":        "guid",
	"name":              "p",
	"location":          "us-east",
	"plan_name":         "instructlab-pricing-plan",
	"resource_group_id": "rg",
	"tags":              []string{"env:test"},
	"access_tags":       []string{"project:rhaii"},
}

func TestRhaiiProjectPlanCreateDefaults(t *testing.T) {
	t.Setenv("IC_ENV_TAGS", "")
	p := plan(t, nil, map[string]interface{}{"name": "p"})
	if got := str(t, p.values["location"]); got != "us-east" {
		t.Errorf("location = %q, want us-east", got)
	}
	if got := str(t, p.values["plan_name"]); got != "instructlab-pricing-plan" {
		t.Errorf("plan_name = %q, want instructlab-pricing-plan", got)
	}
	if got := str(t, p.values["project_id"]); got != "<unknown>" {
		t.Errorf("project_id = %q, want unknown before create", got)
	}
}

// Optional + Computed tags : an unset attribute keeps the
// tags of the server, an empty list detaches all of them.
func TestRhaiiProjectPlanTags(t *testing.T) {
	t.Setenv("IC_ENV_TAGS", "")
	base := map[string]interface{}{"name": "p", "resource_group_id": "rg"}

	p := plan(t, existingProject, base)
	if got := setStrings(t, p.values["tags"]); len(got) != 1 || got[0] != "env:test" {
		t.Errorf("unset tags: planned %v, want the existing [env:test]", got)
	}
	if got := setStrings(t, p.values["access_tags"]); len(got) != 1 || got[0] != "project:rhaii" {
		t.Errorf("unset access_tags: planned %v, want the existing [project:rhaii]", got)
	}
	if got := str(t, p.values["crn"]); got != existingProject["crn"] {
		t.Errorf("crn = %q, want the state value (UseStateForUnknown)", got)
	}
	if len(p.requiresReplace) != 0 {
		t.Errorf("no change must not replace the project: %v", p.requiresReplace)
	}

	empty := map[string]interface{}{"name": "p", "resource_group_id": "rg", "tags": []string{}, "access_tags": []string{}}
	p = plan(t, existingProject, empty)
	if got := setStrings(t, p.values["tags"]); len(got) != 0 {
		t.Errorf("tags = []: planned %v, want empty", got)
	}
	if got := setStrings(t, p.values["access_tags"]); len(got) != 0 {
		t.Errorf("access_tags = []: planned %v, want empty", got)
	}
	if len(p.requiresReplace) != 0 {
		t.Errorf("detaching tags must not replace the project: %v", p.requiresReplace)
	}
}

// TestRhaiiProjectPlanEnvTags checks that the IC_ENV_TAGS environment tags
// are not added to the plan: Terraform rejects a planned value that differs
// from the config value ("planned an invalid value"). The tags are attached by
// the flex helpers and dropped again on read.
func TestRhaiiProjectPlanEnvTags(t *testing.T) {
	t.Setenv("IC_ENV_TAGS", "env:ci")
	p := plan(t, existingProject, map[string]interface{}{"name": "p", "resource_group_id": "rg", "tags": []string{"env:test"}})
	if got := setStrings(t, p.values["tags"]); len(got) != 1 || got[0] != "env:test" {
		t.Errorf("planned tags %v, want [env:test]", got)
	}
}

func TestRhaiiProjectPlanReplace(t *testing.T) {
	t.Setenv("IC_ENV_TAGS", "")
	cases := map[string]struct {
		config  map[string]interface{}
		replace string
	}{
		"rename keeps the project": {map[string]interface{}{"name": "renamed", "resource_group_id": "rg"}, ""},
		"location":                 {map[string]interface{}{"name": "p", "resource_group_id": "rg", "location": "eu-de"}, `AttributeName("location")`},
		"plan_name":                {map[string]interface{}{"name": "p", "resource_group_id": "rg", "plan_name": "other"}, `AttributeName("plan_name")`},
		"resource_group_id":        {map[string]interface{}{"name": "p", "resource_group_id": "rg2"}, `AttributeName("resource_group_id")`},
		"unset resource_group_id":  {map[string]interface{}{"name": "p"}, ""},
	}
	for name, c := range cases {
		p := plan(t, existingProject, c.config)
		if c.replace == "" && len(p.requiresReplace) != 0 {
			t.Errorf("%s: unexpected replace %v", name, p.requiresReplace)
		}
		if c.replace != "" && !p.requiresReplace[c.replace] {
			t.Errorf("%s: want replace on %s, got %v", name, c.replace, p.requiresReplace)
		}
	}
}

// TestRhaiiProjectValidateConfig runs the schema validators through the mux server.
func TestRhaiiProjectValidateConfig(t *testing.T) {
	ctx := context.Background()
	server := muxServer(t)
	schemaResp, _ := server.GetProviderSchema(ctx, &tfprotov6.GetProviderSchemaRequest{})
	typ := schemaResp.ResourceSchemas["ibm_rhaii_project"].ValueType().(tftypes.Object)

	cases := map[string]struct {
		config map[string]interface{}
		valid  bool
	}{
		"valid":                  {map[string]interface{}{"name": "rhaii-nkashyap", "tags": []string{"env:dev"}, "access_tags": []string{"project:rhaii"}}, true},
		"non ascii name":         {map[string]interface{}{"name": "projet-été"}, true},
		"name with slash":        {map[string]interface{}{"name": "a/b"}, false},
		"name too long":          {map[string]interface{}{"name": strings.Repeat("a", 181)}, false},
		"name 180 chars":         {map[string]interface{}{"name": strings.Repeat("a", 180)}, true},
		"tag with =":             {map[string]interface{}{"name": "p", "tags": []string{"env=dev"}}, false},
		"access tag without key": {map[string]interface{}{"name": "p", "access_tags": []string{"rhaii"}}, false},
	}
	for name, c := range cases {
		resp, err := server.ValidateResourceConfig(ctx, &tfprotov6.ValidateResourceConfigRequest{
			TypeName: "ibm_rhaii_project",
			Config:   dynamic(t, typ, objectValue(t, typ, c.config)),
		})
		if err != nil {
			t.Fatal(err)
		}
		hasError := false
		for _, d := range resp.Diagnostics {
			if d.Severity == tfprotov6.DiagnosticSeverityError {
				hasError = true
			}
		}
		if hasError == c.valid {
			t.Errorf("%s: valid = %v, want %v (%v)", name, !hasError, c.valid, resp.Diagnostics)
		}
	}
}

// TestRhaiiInferenceModelValidateConfig runs the validators of the model data source through the mux server.
func TestRhaiiInferenceModelValidateConfig(t *testing.T) {
	ctx := context.Background()
	server := muxServer(t)
	schemaResp, _ := server.GetProviderSchema(ctx, &tfprotov6.GetProviderSchemaRequest{})
	typ := schemaResp.DataSourceSchemas["ibm_rhaii_inference_model"].ValueType().(tftypes.Object)

	cases := map[string]struct {
		model string
		valid bool
	}{
		"spec example":   {"llama-3-3-70b-instruct", true},
		"upper case":     {"Llama", false},
		"underscore":     {"llama_3", false},
		"101 characters": {strings.Repeat("a", 101), false},
		"100 characters": {strings.Repeat("a", 100), true},
		"empty":          {"", false},
	}
	for name, c := range cases {
		resp, err := server.ValidateDataResourceConfig(ctx, &tfprotov6.ValidateDataResourceConfigRequest{
			TypeName: "ibm_rhaii_inference_model",
			Config:   dynamic(t, typ, objectValue(t, typ, map[string]interface{}{"project_id": "p", "model": c.model})),
		})
		if err != nil {
			t.Fatal(err)
		}
		hasError := false
		for _, d := range resp.Diagnostics {
			if d.Severity == tfprotov6.DiagnosticSeverityError {
				hasError = true
			}
		}
		if hasError == c.valid {
			t.Errorf("%s: valid = %v, want %v (%v)", name, !hasError, c.valid, resp.Diagnostics)
		}
	}
}
