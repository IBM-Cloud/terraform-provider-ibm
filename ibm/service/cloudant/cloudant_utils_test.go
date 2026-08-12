// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudant

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/resourcecontroller"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TestMain seeds the provider validator registry before any tests run.
// ResourceIBMResourceInstance() (called by ResourceIBMCloudant()) invokes
// validate.InvokeValidator, which dereferences the ResourceValidatorDictionary
// pointer entries; without registration every call panics with a nil pointer
// dereference. Registering the actual RC validator is sufficient — unit tests
// here never exercise field-level validation so the exact schema content does
// not matter.
func TestMain(m *testing.M) {
	validate.SetValidatorDict(validate.ValidatorDict{
		ResourceValidatorDictionary: map[string]*validate.ResourceValidator{
			"ibm_resource_instance": resourcecontroller.ResourceIBMResourceInstanceValidator(),
		},
		DataSourceValidatorDictionary: map[string]*validate.ResourceValidator{},
	})
	os.Exit(m.Run())
}

// Realistic instance UUIDs and endpoint hostnames used across tests.
const (
	// Gen 1 — extensions.endpoints values are bare hostnames (no scheme).
	gen1PublicHost  = "00000000-0000-0000-0000-000000000000-bluemix.cloudantnosqldb.appdomain.cloud"
	gen1PrivateHost = "00000000-0000-0000-0000-000000000000-bluemix.private.cloudantnosqldb.appdomain.cloud"
	gen1PublicURL   = "https://" + gen1PublicHost
	gen1PrivateURL  = "https://" + gen1PrivateHost

	// Gen 2 — dataservices.connection values already include https://.
	gen2PublicURL  = "https://00000000-0000-0000-0000-000000000000.zzz.cloudant.us-south.dataservices.appdomain.cloud"
	gen2PrivateURL = "https://00000000-0000-0000-0000-000000000000.private.zzz.cloudant.us-south.dataservices.appdomain.cloud"
)

// gen1Ext builds a Gen 1 flattened extensions map with bare hostnames.
func gen1Ext(pub, priv string) flex.Map {
	m := flex.Map{}
	if pub != "" {
		m["endpoints.public"] = pub
	}
	if priv != "" {
		m["endpoints.private"] = priv
	}
	return m
}

// gen2Ext builds a Gen 2 flattened extensions map with full URLs.
func gen2Ext(pub, priv string) flex.Map {
	m := flex.Map{}
	if pub != "" {
		m["dataservices.connection.public_endpoint_url"] = pub
	}
	if priv != "" {
		m["dataservices.connection.vpe_url"] = priv
	}
	return m
}

// ---------------------------------------------------------------------------
// normalizeEndpoint
// ---------------------------------------------------------------------------

func TestNormalizeEndpoint_empty(t *testing.T) {
	if got := normalizeEndpoint(""); got != "" {
		t.Fatalf("expected empty string, got %q", got)
	}
}

func TestNormalizeEndpoint_alreadyHttps(t *testing.T) {
	if got := normalizeEndpoint(gen1PublicURL); got != gen1PublicURL {
		t.Fatalf("expected %q unchanged, got %q", gen1PublicURL, got)
	}
}

func TestNormalizeEndpoint_alreadyHttp(t *testing.T) {
	in := "http://" + gen1PublicHost
	if got := normalizeEndpoint(in); got != in {
		t.Fatalf("expected %q unchanged, got %q", in, got)
	}
}

func TestNormalizeEndpoint_bareHostname(t *testing.T) {
	if got := normalizeEndpoint(gen1PublicHost); got != gen1PublicURL {
		t.Fatalf("expected %q, got %q", gen1PublicURL, got)
	}
}

func TestNormalizeEndpoint_barePrivateHostname(t *testing.T) {
	if got := normalizeEndpoint(gen1PrivateHost); got != gen1PrivateURL {
		t.Fatalf("expected %q, got %q", gen1PrivateURL, got)
	}
}

// ---------------------------------------------------------------------------
// selectCloudantEndpoint
// ---------------------------------------------------------------------------

// Gen 1 — default (public) visibility selects the public endpoint
func TestSelectCloudantEndpoint_gen1_public(t *testing.T) {
	ext := gen1Ext(gen1PublicHost, gen1PrivateHost)
	url, isGen2 := selectCloudantEndpoint(ext, "public")
	if isGen2 {
		t.Fatal("expected Gen 1")
	}
	if url != gen1PublicURL {
		t.Fatalf("expected %q, got %q", gen1PublicURL, url)
	}
}

// Gen 1 — private visibility selects the private endpoint
func TestSelectCloudantEndpoint_gen1_private(t *testing.T) {
	ext := gen1Ext(gen1PublicHost, gen1PrivateHost)
	url, isGen2 := selectCloudantEndpoint(ext, "private")
	if isGen2 {
		t.Fatal("expected Gen 1")
	}
	if url != gen1PrivateURL {
		t.Fatalf("expected %q, got %q", gen1PrivateURL, url)
	}
}

// Gen 1 — public-and-private visibility prefers private when available
func TestSelectCloudantEndpoint_gen1_publicAndPrivate_prefersPrivate(t *testing.T) {
	ext := gen1Ext(gen1PublicHost, gen1PrivateHost)
	url, _ := selectCloudantEndpoint(ext, "public-and-private")
	if url != gen1PrivateURL {
		t.Fatalf("expected %q, got %q", gen1PrivateURL, url)
	}
}

// Gen 1 — public-and-private with no private endpoint falls back to public
func TestSelectCloudantEndpoint_gen1_publicAndPrivate_fallsBackToPublic(t *testing.T) {
	ext := gen1Ext(gen1PublicHost, "")
	url, _ := selectCloudantEndpoint(ext, "public-and-private")
	if url != gen1PublicURL {
		t.Fatalf("expected %q, got %q", gen1PublicURL, url)
	}
}

// Gen 2 — public visibility returns the public endpoint URL as-is (scheme preserved)
func TestSelectCloudantEndpoint_gen2_public(t *testing.T) {
	ext := gen2Ext(gen2PublicURL, gen2PrivateURL)
	url, isGen2 := selectCloudantEndpoint(ext, "public")
	if !isGen2 {
		t.Fatal("expected Gen 2")
	}
	if url != gen2PublicURL {
		t.Fatalf("expected %q, got %q", gen2PublicURL, url)
	}
}

// Gen 2 — private visibility returns the VPE URL
func TestSelectCloudantEndpoint_gen2_private(t *testing.T) {
	ext := gen2Ext(gen2PublicURL, gen2PrivateURL)
	url, isGen2 := selectCloudantEndpoint(ext, "private")
	if !isGen2 {
		t.Fatal("expected Gen 2")
	}
	if url != gen2PrivateURL {
		t.Fatalf("expected %q, got %q", gen2PrivateURL, url)
	}
}

// Gen 2 — public-and-private prefers the VPE URL
func TestSelectCloudantEndpoint_gen2_publicAndPrivate_prefersVPE(t *testing.T) {
	ext := gen2Ext(gen2PublicURL, gen2PrivateURL)
	url, _ := selectCloudantEndpoint(ext, "public-and-private")
	if url != gen2PrivateURL {
		t.Fatalf("expected %q, got %q", gen2PrivateURL, url)
	}
}

// Gen 2 — only a VPE URL present is still detected as Gen 2
func TestSelectCloudantEndpoint_gen2_onlyPrivate(t *testing.T) {
	ext := gen2Ext("", gen2PrivateURL)
	_, isGen2 := selectCloudantEndpoint(ext, "public")
	if !isGen2 {
		t.Fatal("expected Gen 2 when only VPE URL present")
	}
}

// Empty extensions return an empty URL
func TestSelectCloudantEndpoint_empty(t *testing.T) {
	url, _ := selectCloudantEndpoint(flex.Map{}, "public")
	if url != "" {
		t.Fatalf("expected empty URL for empty extensions, got %q", url)
	}
}

// ---------------------------------------------------------------------------
// getCloudantExtensions
// ---------------------------------------------------------------------------

func TestGetCloudantExtensions_valid(t *testing.T) {
	raw := map[string]interface{}{
		"endpoints": map[string]interface{}{
			"public":  gen1PublicHost,
			"private": gen1PrivateHost,
		},
	}
	result, tfErr := getCloudantExtensions(raw, "test", "read")
	if tfErr != nil {
		t.Fatalf("unexpected error: %s", tfErr)
	}
	if result["endpoints.public"] != gen1PublicHost {
		t.Fatalf("expected endpoints.public=%q, got %q", gen1PublicHost, result["endpoints.public"])
	}
}

func TestGetCloudantExtensions_nil(t *testing.T) {
	_, tfErr := getCloudantExtensions(nil, "test", "read")
	if tfErr == nil {
		t.Fatal("expected error for nil extensions")
	}
}

func TestGetCloudantExtensions_wrongType(t *testing.T) {
	_, tfErr := getCloudantExtensions("not-a-map", "test", "read")
	if tfErr == nil {
		t.Fatal("expected error for non-map extensions")
	}
}

func TestGetCloudantExtensions_empty(t *testing.T) {
	_, tfErr := getCloudantExtensions(map[string]interface{}{}, "test", "read")
	if tfErr == nil {
		t.Fatal("expected error for empty extensions map")
	}
}

// ---------------------------------------------------------------------------
// isCloudantGen2Plan / isCloudantGen2PlanFrom
// ---------------------------------------------------------------------------

func TestIsCloudantGen2Plan_gen1Plans(t *testing.T) {
	for _, plan := range []string{"lite", "standard", "dedicated-hardware", "Lite", "STANDARD", "  standard  "} {
		if isCloudantGen2Plan(plan) {
			t.Errorf("plan %q should be Gen 1, got Gen 2", plan)
		}
	}
}

func TestIsCloudantGen2Plan_gen2Plans(t *testing.T) {
	for _, plan := range []string{"standard-gen2", "another-gen2"} {
		if !isCloudantGen2Plan(plan) {
			t.Errorf("plan %q should be Gen 2, got Gen 1", plan)
		}
	}
}

func TestIsCloudantGen2Plan_empty(t *testing.T) {
	if isCloudantGen2Plan("") {
		t.Fatal("empty plan should default to Gen 1 (false)")
	}
}

func TestIsCloudantGen2Plan_whitespaceOnly(t *testing.T) {
	if isCloudantGen2Plan("   ") {
		t.Fatal("whitespace-only plan should default to Gen 1 (false)")
	}
}

// ---------------------------------------------------------------------------
// cloudantToResourceInstance / resourceInstanceToCloudant — helpers
// ---------------------------------------------------------------------------

// makeGen1CloudantRD builds a Gen 1 Cloudant ResourceData populated with shared
// RC fields plus Gen 1-specific Cloudant-only fields.
func makeGen1CloudantRD(t *testing.T) *schema.ResourceData {
	t.Helper()
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		// shared RC fields
		"name":              "my-cloudant",
		"plan":              "standard",
		"location":          "us-south",
		"resource_group_id": "rg-abc",
		"service":           "cloudantnosqldb",
		"guid":              "guid-123",
		"crn":               "crn:v1:bluemix:public:cloudantnosqldb:us-south:::",
		"status":            "active",
		// Gen 1 Cloudant-only fields
		"legacy_credentials":  true,
		"environment_crn":     "crn:v1:bluemix:public:cloudantnosqldb:us-south:a/abc:hw-instance::",
		"include_data_events": true,
		"capacity":            1,
		"enable_cors":         false,
	})
	d.SetId("instance-id-001")
	return d
}

// makeGen2CloudantRD builds a Gen 2 Cloudant ResourceData populated with shared
// RC fields plus Gen 2-specific Cloudant-only fields.
func makeGen2CloudantRD(t *testing.T) *schema.ResourceData {
	t.Helper()
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		// shared RC fields
		"name":              "my-cloudant-gen2",
		"plan":              "standard-gen2",
		"location":          "us-south",
		"resource_group_id": "rg-abc",
		"service":           "cloudantnosqldb",
		"guid":              "guid-456",
		"crn":               "crn:v1:bluemix:public:cloudantnosqldb:us-south:::",
		"status":            "active",
		// Gen 2 Cloudant-only fields
		"capacity":    3,
		"enable_cors": true,
		"cors_config": []interface{}{
			map[string]interface{}{
				"allow_credentials": true,
				"origins":           []interface{}{"https://example.com", "https://other.com"},
			},
		},
	})
	d.SetId("instance-id-002")
	return d
}

// unmarshalParamsJSON decodes the parameters_json field on d into a map.
func unmarshalParamsJSON(t *testing.T, d *schema.ResourceData) map[string]interface{} {
	t.Helper()
	raw := d.Get("parameters_json").(string)
	if raw == "" {
		t.Fatal("parameters_json is empty")
	}
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		t.Fatalf("failed to unmarshal parameters_json: %s", err)
	}
	return m
}

// ---------------------------------------------------------------------------
// cloudantToResourceInstance — Gen 1
// ---------------------------------------------------------------------------

func TestCloudantToResourceInstance_gen1_legacyCredentialsTrue(t *testing.T) {
	d := makeGen1CloudantRD(t) // legacy_credentials=true
	cloudantToResourceInstance(d)

	params := d.Get("parameters").(map[string]interface{})
	if params["legacyCredentials"] != "true" {
		t.Errorf("expected legacyCredentials=%q, got %v", "true", params["legacyCredentials"])
	}
}

func TestCloudantToResourceInstance_gen1_legacyCredentialsFalse(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan":               "standard",
		"legacy_credentials": false,
		"capacity":           1,
		"enable_cors":        false,
	})
	cloudantToResourceInstance(d)

	params := d.Get("parameters").(map[string]interface{})
	if params["legacyCredentials"] != "false" {
		t.Errorf("expected legacyCredentials=%q, got %v", "false", params["legacyCredentials"])
	}
}

func TestCloudantToResourceInstance_gen1_environmentCRN(t *testing.T) {
	d := makeGen1CloudantRD(t)
	cloudantToResourceInstance(d)

	params := d.Get("parameters").(map[string]interface{})
	if params["environment_crn"] != "crn:v1:bluemix:public:cloudantnosqldb:us-south:a/abc:hw-instance::" {
		t.Errorf("unexpected environment_crn: %v", params["environment_crn"])
	}
}

func TestCloudantToResourceInstance_gen1_noGen2Fields(t *testing.T) {
	d := makeGen1CloudantRD(t)
	cloudantToResourceInstance(d)

	params := d.Get("parameters").(map[string]interface{})
	if _, ok := params["dataservices"]; ok {
		t.Error("Gen 1 parameters should not contain dataservices key")
	}
}

func TestCloudantToResourceInstance_gen1_userParamsPassthrough(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan":               "standard",
		"legacy_credentials": false,
		"capacity":           1,
		"enable_cors":        false,
		"parameters": map[string]interface{}{
			"custom_key": "custom_value",
		},
	})
	cloudantToResourceInstance(d)

	params := d.Get("parameters").(map[string]interface{})
	// Unrelated user key must pass through.
	if params["custom_key"] != "custom_value" {
		t.Errorf("expected user parameter custom_key to pass through, got %v", params["custom_key"])
	}
	// Typed field must also be present.
	if params["legacyCredentials"] != "false" {
		t.Errorf("expected legacyCredentials=%q, got %v", "false", params["legacyCredentials"])
	}
}

func TestCloudantToResourceInstance_gen1_typedFieldWinsOverUserParam(t *testing.T) {
	// User sets legacy_credentials=true via the schema field but also sneaks
	// "legacyCredentials": "false" into the raw parameters map. The typed
	// schema field must win.
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan":               "standard",
		"legacy_credentials": true,
		"capacity":           1,
		"enable_cors":        false,
		"parameters": map[string]interface{}{
			"legacyCredentials": "false",
		},
	})
	cloudantToResourceInstance(d)

	params := d.Get("parameters").(map[string]interface{})
	if params["legacyCredentials"] != "true" {
		t.Errorf("typed legacy_credentials=true should win over user param, got %v", params["legacyCredentials"])
	}
}

// ---------------------------------------------------------------------------
// cloudantToResourceInstance — Gen 2
// ---------------------------------------------------------------------------

func TestCloudantToResourceInstance_gen2_capacityAndCorsEnabled(t *testing.T) {
	d := makeGen2CloudantRD(t)
	cloudantToResourceInstance(d)

	params := unmarshalParamsJSON(t, d)

	if _, ok := params["legacyCredentials"]; ok {
		t.Error("Gen 2 parameters_json should not contain legacyCredentials")
	}

	ds, ok := params["dataservices"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected dataservices map, got %T", params["dataservices"])
	}
	cloudant, ok := ds["cloudant"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected cloudant map, got %T", ds["cloudant"])
	}
	// JSON numbers unmarshal as float64.
	if cloudant["capacity_units"].(float64) != 3 {
		t.Fatalf("expected capacity_units=3, got %v", cloudant["capacity_units"])
	}
	cfg, ok := cloudant["configuration"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected configuration map, got %T", cloudant["configuration"])
	}
	cors, ok := cfg["cors"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected cors map, got %T", cfg["cors"])
	}
	if cors["enabled"] != true {
		t.Fatalf("expected cors.enabled=true, got %v", cors["enabled"])
	}
	if cors["allowCredentials"] != true {
		t.Fatalf("expected cors.allowCredentials=true, got %v", cors["allowCredentials"])
	}
	origins := cors["origins"].([]interface{})
	if len(origins) != 2 || origins[0] != "https://example.com" || origins[1] != "https://other.com" {
		t.Fatalf("unexpected origins: %v", origins)
	}
}

func TestCloudantToResourceInstance_gen2_corsDisabled(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan":        "standard-gen2",
		"capacity":    1,
		"enable_cors": false,
	})
	cloudantToResourceInstance(d)

	params := unmarshalParamsJSON(t, d)
	ds := params["dataservices"].(map[string]interface{})
	cloudant := ds["cloudant"].(map[string]interface{})
	cfg := cloudant["configuration"].(map[string]interface{})
	cors := cfg["cors"].(map[string]interface{})

	if cors["enabled"] != false {
		t.Fatalf("expected cors.enabled=false, got %v", cors["enabled"])
	}
	if _, ok := cors["allowCredentials"]; ok {
		t.Fatal("cors.allowCredentials should not be set when CORS is disabled")
	}
	if _, ok := cors["origins"]; ok {
		t.Fatal("cors.origins should not be set when CORS is disabled")
	}
}

func TestCloudantToResourceInstance_gen2_noGen1Fields(t *testing.T) {
	d := makeGen2CloudantRD(t)
	cloudantToResourceInstance(d)

	params := unmarshalParamsJSON(t, d)
	if _, ok := params["legacyCredentials"]; ok {
		t.Error("Gen 2 parameters_json should not contain legacyCredentials")
	}
	if _, ok := params["environment_crn"]; ok {
		t.Error("Gen 2 parameters_json should not contain environment_crn")
	}
}

func TestCloudantToResourceInstance_gen2_userParamsPassthrough(t *testing.T) {
	// User supplies a custom key and a conflicting typed key via parameters.
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan":        "standard-gen2",
		"capacity":    2,
		"enable_cors": false,
		"parameters": map[string]interface{}{
			"custom_key": "custom_value",
		},
	})
	cloudantToResourceInstance(d)

	params := unmarshalParamsJSON(t, d)
	// Unrelated user key must pass through into the top-level JSON object.
	if params["custom_key"] != "custom_value" {
		t.Errorf("expected user parameter custom_key to pass through, got %v", params["custom_key"])
	}
	// Typed capacity=2 must be present in the nested structure.
	ds := params["dataservices"].(map[string]interface{})
	cloudant := ds["cloudant"].(map[string]interface{})
	if cloudant["capacity_units"].(float64) != 2 {
		t.Errorf("expected capacity_units=2, got %v", cloudant["capacity_units"])
	}
}

// ---------------------------------------------------------------------------
// resourceInstanceToCloudant — shared fields
// ---------------------------------------------------------------------------

func TestResourceInstanceToCloudant_sharedFieldsCopied(t *testing.T) {
	// resourceInstanceToCloudant only mutates Cloudant-specific fields; shared
	// RC fields are already set by the RC read before it is called. Verify that
	// a Gen 1 d with pre-set RC fields is not disturbed.
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"name":     "rc-instance",
		"plan":     "lite",
		"location": "eu-gb",
		"guid":     "guid-789",
		"status":   "active",
		"crn":      "crn:v1:bluemix:public:cloudantnosqldb:eu-gb:::",
		"parameters": map[string]interface{}{
			"legacyCredentials": "false",
		},
	})
	d.SetId("rc-instance-id-002")
	resourceInstanceToCloudant(d)

	if d.Id() != "rc-instance-id-002" {
		t.Errorf("expected ID=%q, got %q", "rc-instance-id-002", d.Id())
	}
	if d.Get("name") != "rc-instance" {
		t.Errorf("expected name=%q, got %v", "rc-instance", d.Get("name"))
	}
	if d.Get("plan") != "lite" {
		t.Errorf("expected plan=%q, got %v", "lite", d.Get("plan"))
	}
	if d.Get("guid") != "guid-789" {
		t.Errorf("expected guid=%q, got %v", "guid-789", d.Get("guid"))
	}
}

func TestResourceInstanceToCloudant_flexFieldsPreserved(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"name":                     "rc-instance",
		"plan":                     "standard",
		"location":                 "us-south",
		flex.ResourceName:          "rc-instance",
		flex.ResourceCRN:           "crn:v1:bluemix:public:::",
		flex.ResourceStatus:        "active",
		flex.ResourceGroupName:     "default",
		flex.ResourceControllerURL: "https://cloud.ibm.com/services/",
		"parameters": map[string]interface{}{
			"legacyCredentials": "false",
		},
	})
	resourceInstanceToCloudant(d)

	if d.Get(flex.ResourceCRN) != "crn:v1:bluemix:public:::" {
		t.Errorf("expected %s preserved, got %v", flex.ResourceCRN, d.Get(flex.ResourceCRN))
	}
	if d.Get(flex.ResourceControllerURL) != "https://cloud.ibm.com/services/" {
		t.Errorf("expected %s preserved, got %v", flex.ResourceControllerURL, d.Get(flex.ResourceControllerURL))
	}
}

// ---------------------------------------------------------------------------
// resourceInstanceToCloudant — Gen 1
// ---------------------------------------------------------------------------

func TestResourceInstanceToCloudant_gen1_legacyCredentials(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan": "standard",
		"parameters": map[string]interface{}{
			"legacyCredentials": "true",
		},
	})
	resourceInstanceToCloudant(d)

	if !d.Get("legacy_credentials").(bool) {
		t.Error("expected legacy_credentials=true")
	}
	// Key must be removed from parameters.
	params := d.Get("parameters")
	if params != nil {
		if m, ok := params.(map[string]interface{}); ok {
			if _, found := m["legacyCredentials"]; found {
				t.Error("legacyCredentials should have been removed from parameters")
			}
		}
	}
}

func TestResourceInstanceToCloudant_gen1_legacyCredentialsFalse(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan": "standard",
		"parameters": map[string]interface{}{
			"legacyCredentials": "false",
		},
	})
	resourceInstanceToCloudant(d)

	if d.Get("legacy_credentials").(bool) {
		t.Error("expected legacy_credentials=false")
	}
}

func TestResourceInstanceToCloudant_gen1_environmentCRN(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan": "dedicated-hardware",
		"parameters": map[string]interface{}{
			"environment_crn": "crn:v1:bluemix:public:cloudantnosqldb:us-south:a/abc:hw-instance::",
		},
	})
	resourceInstanceToCloudant(d)

	if d.Get("environment_crn") != "crn:v1:bluemix:public:cloudantnosqldb:us-south:a/abc:hw-instance::" {
		t.Errorf("unexpected environment_crn: %v", d.Get("environment_crn"))
	}
	// Key must be removed from parameters.
	params := d.Get("parameters")
	if params != nil {
		if m, ok := params.(map[string]interface{}); ok {
			if _, found := m["environment_crn"]; found {
				t.Error("environment_crn should have been removed from parameters")
			}
		}
	}
}

func TestResourceInstanceToCloudant_gen1_userKeysPreservedAfterUnmerge(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan": "standard",
		"parameters": map[string]interface{}{
			"legacyCredentials": "true",
			"custom_key":        "custom_value",
		},
	})
	resourceInstanceToCloudant(d)

	// legacyCredentials extracted and removed; custom_key must survive.
	params := d.Get("parameters").(map[string]interface{})
	if _, found := params["legacyCredentials"]; found {
		t.Error("legacyCredentials should have been removed from parameters")
	}
	if params["custom_key"] != "custom_value" {
		t.Errorf("expected custom_key to survive, got %v", params["custom_key"])
	}
}

func TestResourceInstanceToCloudant_gen1_noParameters(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan": "standard",
	})
	// Must not panic when no parameters are present.
	resourceInstanceToCloudant(d)
}

// ---------------------------------------------------------------------------
// resourceInstanceToCloudant — Gen 2
// ---------------------------------------------------------------------------

func TestResourceInstanceToCloudant_gen2_capacityAndCors(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan": "standard-gen2",
		"extensions": map[string]interface{}{
			"dataservices.cloudant.capacity_units":                      "3",
			"dataservices.cloudant.configuration.cors.enabled":          "true",
			"dataservices.cloudant.configuration.cors.allowCredentials": "true",
			"dataservices.cloudant.configuration.cors.origins.#":        "2",
			"dataservices.cloudant.configuration.cors.origins.0":        "https://example.com",
			"dataservices.cloudant.configuration.cors.origins.1":        "https://other.com",
		},
	})
	resourceInstanceToCloudant(d)

	if got := d.Get("capacity").(int); got != 3 {
		t.Errorf("expected capacity=3, got %d", got)
	}
	if !d.Get("enable_cors").(bool) {
		t.Error("expected enable_cors=true")
	}
	cors := d.Get("cors_config").([]interface{})
	if len(cors) != 1 {
		t.Fatalf("expected 1 cors_config entry, got %d", len(cors))
	}
	entry := cors[0].(map[string]interface{})
	if !entry["allow_credentials"].(bool) {
		t.Error("expected allow_credentials=true")
	}
	origins := entry["origins"].([]interface{})
	if len(origins) != 2 || origins[0] != "https://example.com" || origins[1] != "https://other.com" {
		t.Fatalf("unexpected origins: %v", origins)
	}
}

func TestResourceInstanceToCloudant_gen2_corsDisabled(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan": "standard-gen2",
		"extensions": map[string]interface{}{
			"dataservices.cloudant.capacity_units":               "1",
			"dataservices.cloudant.configuration.cors.enabled":   "false",
			"dataservices.cloudant.configuration.cors.origins.#": "0",
		},
	})
	resourceInstanceToCloudant(d)

	if d.Get("enable_cors").(bool) {
		t.Error("expected enable_cors=false")
	}
	if cors := d.Get("cors_config").([]interface{}); len(cors) != 0 {
		t.Errorf("expected empty cors_config when CORS disabled, got %d entries", len(cors))
	}
}

func TestResourceInstanceToCloudant_gen2_noExtensions(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan": "standard-gen2",
	})
	// Must not panic when extensions are absent.
	resourceInstanceToCloudant(d)
}

func TestResourceInstanceToCloudant_gen2_doesNotClearParametersJSON(t *testing.T) {
	const original = `{"dataservices":{"cloudant":{"capacity_units":3}}}`
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan":            "standard-gen2",
		"parameters_json": original,
	})
	resourceInstanceToCloudant(d)

	if got := d.Get("parameters_json").(string); got != original {
		t.Errorf("expected parameters_json to be left unchanged, got %q", got)
	}
}

func TestResourceInstanceToCloudant_gen2_setsThroughputEmpty(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan": "standard-gen2",
		"extensions": map[string]interface{}{
			"dataservices.cloudant.capacity_units": "1",
		},
	})
	resourceInstanceToCloudant(d)

	throughput := d.Get("throughput").(map[string]interface{})
	if len(throughput) != 0 {
		t.Errorf("expected throughput to be empty map for Gen 2, got %v", throughput)
	}
}

func TestResourceInstanceToCloudant_gen2_gen1FieldsUnchanged(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan": "standard-gen2",
		"extensions": map[string]interface{}{
			"dataservices.cloudant.capacity_units": "2",
		},
		"legacy_credentials": false,
	})
	resourceInstanceToCloudant(d)

	// Gen 2 path must not touch legacy_credentials.
	if d.Get("legacy_credentials").(bool) {
		t.Error("Gen 2 path must not modify legacy_credentials")
	}
}

// ---------------------------------------------------------------------------
// Round-trip — Gen 1
// ---------------------------------------------------------------------------

func TestRoundTrip_gen1_cloudantToRCAndBack(t *testing.T) {
	d := makeGen1CloudantRD(t)

	// Step 1: prepare for RC call — merges Cloudant fields into parameters.
	cloudantToResourceInstance(d)

	// Verify the merged parameters are present before the simulated RC call.
	params := d.Get("parameters").(map[string]interface{})
	if params["legacyCredentials"] != "true" {
		t.Fatalf("expected legacyCredentials=true in parameters before RC call, got %v", params["legacyCredentials"])
	}

	// Step 2: simulate RC call returning (parameters still on d), then
	// resourceInstanceToCloudant extracts and removes Cloudant-specific keys.
	resourceInstanceToCloudant(d)

	// Shared fields must be untouched.
	if d.Get("name") != "my-cloudant" {
		t.Errorf("unexpected name: %v", d.Get("name"))
	}
	if d.Get("plan") != "standard" {
		t.Errorf("unexpected plan: %v", d.Get("plan"))
	}

	// Cloudant-specific fields must be extracted.
	if !d.Get("legacy_credentials").(bool) {
		t.Error("legacy_credentials should be true after round-trip")
	}
	if d.Get("environment_crn") != "crn:v1:bluemix:public:cloudantnosqldb:us-south:a/abc:hw-instance::" {
		t.Errorf("unexpected environment_crn: %v", d.Get("environment_crn"))
	}

	// Cloudant-specific keys must be removed from parameters.
	cleanParams := d.Get("parameters")
	if cleanParams != nil {
		if m, ok := cleanParams.(map[string]interface{}); ok {
			if _, found := m["legacyCredentials"]; found {
				t.Error("legacyCredentials should be removed from parameters after round-trip")
			}
		}
	}
}

// ---------------------------------------------------------------------------
// Round-trip — Gen 2
// ---------------------------------------------------------------------------

func TestRoundTrip_gen2_cloudantToRCAndBack(t *testing.T) {
	d := makeGen2CloudantRD(t)

	// Step 1: prepare for RC call — encodes Cloudant fields into parameters_json.
	cloudantToResourceInstance(d)

	// Verify parameters_json is set.
	if d.Get("parameters_json").(string) == "" {
		t.Fatal("parameters_json should be set after cloudantToResourceInstance")
	}

	// Step 2: simulate RC response — extensions are populated (mirroring what
	// flex.Flatten would produce from the broker's nested response).
	d.Set("extensions", map[string]interface{}{
		"dataservices.cloudant.capacity_units":                      "3",
		"dataservices.cloudant.configuration.cors.enabled":          "true",
		"dataservices.cloudant.configuration.cors.allowCredentials": "true",
		"dataservices.cloudant.configuration.cors.origins.#":        "2",
		"dataservices.cloudant.configuration.cors.origins.0":        "https://example.com",
		"dataservices.cloudant.configuration.cors.origins.1":        "https://other.com",
	})
	resourceInstanceToCloudant(d)

	// Shared fields must survive.
	if d.Get("name") != "my-cloudant-gen2" {
		t.Errorf("unexpected name: %v", d.Get("name"))
	}
	if d.Get("plan") != "standard-gen2" {
		t.Errorf("unexpected plan: %v", d.Get("plan"))
	}

	// Gen 2 Cloudant fields must be populated from extensions.
	if got := d.Get("capacity").(int); got != 3 {
		t.Errorf("expected capacity=3 after Gen 2 round-trip, got %v", got)
	}
	if !d.Get("enable_cors").(bool) {
		t.Error("expected enable_cors=true after Gen 2 round-trip")
	}
	cors := d.Get("cors_config").([]interface{})
	if len(cors) != 1 {
		t.Fatalf("expected 1 cors_config entry after Gen 2 round-trip, got %d", len(cors))
	}
	entry := cors[0].(map[string]interface{})
	origins := entry["origins"].([]interface{})
	if len(origins) != 2 {
		t.Errorf("expected 2 origins after Gen 2 round-trip, got %d", len(origins))
	}
}

// ---------------------------------------------------------------------------
// include_data_events — Gen 2 (cloudantToResourceInstance)
// ---------------------------------------------------------------------------

func TestCloudantToResourceInstance_gen2_includeDataEventsTrue(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan":                "standard-gen2",
		"capacity":            1,
		"enable_cors":         false,
		"include_data_events": true,
	})
	cloudantToResourceInstance(d)

	params := unmarshalParamsJSON(t, d)
	ds := params["dataservices"].(map[string]interface{})
	cloudant := ds["cloudant"].(map[string]interface{})
	cfg := cloudant["configuration"].(map[string]interface{})
	audit, ok := cfg["audit"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected audit map in configuration, got %T", cfg["audit"])
	}
	if audit["data_events"] != true {
		t.Fatalf("expected audit.data_events=true, got %v", audit["data_events"])
	}
}

func TestCloudantToResourceInstance_gen2_includeDataEventsFalse(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan":                "standard-gen2",
		"capacity":            1,
		"enable_cors":         false,
		"include_data_events": false,
	})
	cloudantToResourceInstance(d)

	params := unmarshalParamsJSON(t, d)
	ds := params["dataservices"].(map[string]interface{})
	cloudant := ds["cloudant"].(map[string]interface{})
	cfg := cloudant["configuration"].(map[string]interface{})
	audit, ok := cfg["audit"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected audit map in configuration, got %T", cfg["audit"])
	}
	if audit["data_events"] != false {
		t.Fatalf("expected audit.data_events=false, got %v", audit["data_events"])
	}
}

// ---------------------------------------------------------------------------
// include_data_events — Gen 2 (resourceInstanceToCloudant)
// ---------------------------------------------------------------------------

func TestResourceInstanceToCloudant_gen2_includeDataEventsTrue(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan": "standard-gen2",
		"extensions": map[string]interface{}{
			"dataservices.cloudant.capacity_units":                  "1",
			"dataservices.cloudant.configuration.audit.data_events": "true",
			"dataservices.cloudant.configuration.cors.enabled":      "false",
		},
	})
	resourceInstanceToCloudant(d)

	if !d.Get("include_data_events").(bool) {
		t.Error("expected include_data_events=true")
	}
}

func TestResourceInstanceToCloudant_gen2_includeDataEventsFalse(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan": "standard-gen2",
		"extensions": map[string]interface{}{
			"dataservices.cloudant.capacity_units":                  "1",
			"dataservices.cloudant.configuration.audit.data_events": "false",
			"dataservices.cloudant.configuration.cors.enabled":      "false",
		},
	})
	resourceInstanceToCloudant(d)

	if d.Get("include_data_events").(bool) {
		t.Error("expected include_data_events=false")
	}
}

func TestResourceInstanceToCloudant_gen2_includeDataEventsAbsent(t *testing.T) {
	// When the extension key is absent the attribute should remain at its zero
	// value (false); it must not panic.
	d := schema.TestResourceDataRaw(t, ResourceIBMCloudant().Schema, map[string]interface{}{
		"plan": "standard-gen2",
		"extensions": map[string]interface{}{
			"dataservices.cloudant.capacity_units":             "1",
			"dataservices.cloudant.configuration.cors.enabled": "false",
		},
	})
	resourceInstanceToCloudant(d)

	if d.Get("include_data_events").(bool) {
		t.Error("expected include_data_events=false when extension key is absent")
	}
}
