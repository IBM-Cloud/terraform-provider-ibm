// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package watsonx

import "testing"

func TestAPIBase(t *testing.T) {
	cases := []struct {
		region, endpoint, want string
		wantErr                bool
	}{
		{"us-south", "", "https://api.dataplatform.cloud.ibm.com", false},
		{"ca-tor", "", "https://api.ca-tor.dai.cloud.ibm.com", false},
		{"us-south", "//private.example.com/", "https://private.example.com", false},
		{"mars-1", "", "", true},
	}
	for _, c := range cases {
		got, err := apiBase(c.region, c.endpoint)
		if (err != nil) != c.wantErr || got != c.want {
			t.Errorf("apiBase(%q, %q) = %q, %v; want %q, err=%v", c.region, c.endpoint, got, err, c.want, c.wantErr)
		}
	}
	if got := uiBase("eu-de"); got != "https://eu-de.dataplatform.cloud.ibm.com" {
		t.Errorf("uiBase(eu-de) = %q", got)
	}
}

func TestSplitImportID(t *testing.T) {
	if r, id, ok := splitImportID("us-south/abc-123"); !ok || r != "us-south" || id != "abc-123" {
		t.Errorf("got %q %q %v", r, id, ok)
	}
	for _, bad := range []string{"", "us-south", "/abc", "us-south/"} {
		if _, _, ok := splitImportID(bad); ok {
			t.Errorf("splitImportID(%q) should fail", bad)
		}
	}
}

func TestIsNotFound(t *testing.T) {
	if !isNotFound(&apiError{StatusCode: 404}) || isNotFound(&apiError{StatusCode: 500}) {
		t.Error("isNotFound misclassified status codes")
	}
}
