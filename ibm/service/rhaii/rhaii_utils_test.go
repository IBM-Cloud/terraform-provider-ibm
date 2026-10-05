// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"testing"
)

func TestRhaiiProjectEndpoint(t *testing.T) {
	got := rhaiiProjectEndpoint("us-east", "917bc95a-fef0-4039-b936-e0b6fb17b721")
	want := "https://us-east.rhai.ibm.com/v1/projects/917bc95a-fef0-4039-b936-e0b6fb17b721"
	if got != want {
		t.Fatalf("rhaiiProjectEndpoint() = %q, want %q", got, want)
	}
}

func TestResourceIBMRhaiiProjectSchema(t *testing.T) {
	r := ResourceIBMRhaiiProject()
	if err := r.InternalValidate(nil, true); err != nil {
		t.Fatalf("resource schema is not valid: %s", err)
	}

	s := r.Schema
	if s["service"].Required || s["service"].Optional || !s["service"].Computed {
		t.Errorf("service must be computed only")
	}
	if s["plan"].Default != rhaiiDefaultPlan {
		t.Errorf("plan default = %v, want %q", s["plan"].Default, rhaiiDefaultPlan)
	}
	if s["location"].Default != rhaiiDefaultLocation || !s["location"].ForceNew {
		t.Errorf("location must default to %q and force a new resource", rhaiiDefaultLocation)
	}
	if !s["name"].Required {
		t.Errorf("name must be required")
	}
	for _, k := range []string{"project_id", "endpoint"} {
		if s[k] == nil || !s[k].Computed {
			t.Errorf("%s must be a computed attribute", k)
		}
	}
}

func TestDataSourceIBMRhaiiProjectSchema(t *testing.T) {
	r := DataSourceIBMRhaiiProject()
	if err := r.InternalValidate(nil, false); err != nil {
		t.Fatalf("data source schema is not valid: %s", err)
	}

	s := r.Schema
	if s["service"].Optional || !s["service"].Computed {
		t.Errorf("service must be computed only")
	}
	for _, k := range s["identifier"].ConflictsWith {
		if k == "service" {
			t.Errorf("identifier must not conflict with the computed service attribute")
		}
	}
	for _, k := range []string{"project_id", "endpoint"} {
		if s[k] == nil || !s[k].Computed {
			t.Errorf("%s must be a computed attribute", k)
		}
	}
}
