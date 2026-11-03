// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"reflect"
	"testing"
)

var testCisMtlsOwner = cisMtlsHostOwnerLookup(map[string]string{
	"za":  "zone-a.example.com",
	"zb":  "zone-b.example.com",
	"zbs": "sub.zone-b.example.com",
})

func TestCisMtlsSchemaIsValid(t *testing.T) {
	if err := ResourceIBMCISMtls().InternalValidate(nil, true); err != nil {
		t.Fatal(err)
	}
}

func TestCisMtlsHostOwnerLookup(t *testing.T) {
	cases := map[string]string{
		"api.zone-a.example.com":   "za",
		"zone-a.example.com":       "za",
		"API.Zone-B.example.com.":  "zb",
		"x.sub.zone-b.example.com": "zbs",
		"evilzone-a.example.com":   "",
		"api.zone-c.example.com":   "",
	}
	for host, want := range cases {
		if got := testCisMtlsOwner(host); got != want {
			t.Errorf("%s: got %q, want %q", host, got, want)
		}
	}
}

func TestPlanCisMtlsZoneUpdates(t *testing.T) {
	cases := []struct {
		name    string
		current []string
		zones   []cisMtlsZoneHosts
		want    []cisMtlsZoneHosts
	}{
		{
			name:    "create through first zone, add second",
			current: []string{"api.zone-a.example.com"},
			zones: []cisMtlsZoneHosts{
				{zoneID: "za", hostnames: []string{"api.zone-a.example.com"}},
				{zoneID: "zb", hostnames: []string{"api.zone-b.example.com"}},
			},
			want: []cisMtlsZoneHosts{
				{zoneID: "zb", hostnames: []string{"api.zone-a.example.com", "api.zone-b.example.com"}},
			},
		},
		{
			name:    "no change",
			current: []string{"api.zone-a.example.com", "api.zone-b.example.com"},
			zones: []cisMtlsZoneHosts{
				{zoneID: "za", hostnames: []string{"api.zone-a.example.com"}},
				{zoneID: "zb", hostnames: []string{"api.zone-b.example.com"}},
			},
			want: []cisMtlsZoneHosts{},
		},
		{
			name:    "add and remove per zone, keep unmanaged host",
			current: []string{"api.zone-a.example.com", "old.zone-b.example.com", "api.zone-c.example.com"},
			zones: []cisMtlsZoneHosts{
				{zoneID: "za", hostnames: []string{"api.zone-a.example.com", "www.zone-a.example.com"}},
				{zoneID: "zb", hostnames: []string{"api.zone-b.example.com"}},
			},
			want: []cisMtlsZoneHosts{
				{zoneID: "za", hostnames: []string{"api.zone-a.example.com", "old.zone-b.example.com", "api.zone-c.example.com", "www.zone-a.example.com"}},
				{zoneID: "zb", hostnames: []string{"api.zone-a.example.com", "api.zone-c.example.com", "www.zone-a.example.com", "api.zone-b.example.com"}},
			},
		},
		{
			name:    "zone removed from configuration",
			current: []string{"api.zone-a.example.com", "api.zone-b.example.com"},
			zones: []cisMtlsZoneHosts{
				{zoneID: "za", hostnames: []string{"api.zone-a.example.com"}},
				{zoneID: "zb"},
			},
			want: []cisMtlsZoneHosts{
				{zoneID: "zb", hostnames: []string{"api.zone-a.example.com"}},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := planCisMtlsZoneUpdates(c.current, c.zones, testCisMtlsOwner)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got %+v, want %+v", got, c.want)
			}
		})
	}
}

func TestOrderCisMtlsHostnames(t *testing.T) {
	got := orderCisMtlsHostnames([]string{"c", "a", "d"}, []string{"a", "b", "c"})
	if want := []string{"a", "c", "d"}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
