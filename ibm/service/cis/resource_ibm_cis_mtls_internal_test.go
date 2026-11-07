// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

// Thawte EV RSA CA G2, a public DigiCert intermediate; the fingerprint is what the CIS API returned for it.
const testCisMtlsCertPEM = `-----BEGIN CERTIFICATE-----
MIIFOjCCBCKgAwIBAgIQB8LG0yxvDgqrrA3Q+fzVszANBgkqhkiG9w0BAQsFADBh
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBH
MjAeFw0yMDA3MDIxMjQzMDJaFw0zMDA3MDIxMjQzMDJaMEIxCzAJBgNVBAYTAlVT
MRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxHDAaBgNVBAMTE1RoYXd0ZSBFViBSU0Eg
Q0EgRzIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDZ6jsIHs3bmIoe
7DvnuSvGX375jpWKv25Gf7uz8GZQT3DFVHeRS0NKn5WosqMVNYJKKEImR1Gb2dxl
90/GDypngguJeBVxdTtgGWPpGnuKPXfc6Qy8yKS1SRQsEs2q2EHl0tEU+/vyNro0
DDZcox947My8htb1dLx0Q/y9KjEuagQ5AXeLidtiaiAyKnThZRslD8EK/EcHRvkP
AVMSfGyCVmho3VLBP7LAlLA/RyrAX282OfK8lPZqsASNsOQhmsZPT3IuQhXz7RXc
nCpBM+GfZoFSP6+uO5j8TZMkqLTAuVAsSyGY8zNbzYupA7QmPcIfAwqG1oD9dGah
2GgKoeTnAgMBAAGjggILMIICBzAdBgNVHQ4EFgQUbC7kYbTDub3wyq2mwWh6uNTM
HaAwHwYDVR0jBBgwFoAUTiJUIBiV5uNu5g/6+rkS7QYXjzkwDgYDVR0PAQH/BAQD
AgGGMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjASBgNVHRMBAf8ECDAG
AQH/AgEAMDQGCCsGAQUFBwEBBCgwJjAkBggrBgEFBQcwAYYYaHR0cDovL29jc3Au
ZGlnaWNlcnQuY29tMHsGA1UdHwR0MHIwN6A1oDOGMWh0dHA6Ly9jcmwzLmRpZ2lj
ZXJ0LmNvbS9EaWdpQ2VydEdsb2JhbFJvb3RHMi5jcmwwN6A1oDOGMWh0dHA6Ly9j
cmw0LmRpZ2ljZXJ0LmNvbS9EaWdpQ2VydEdsb2JhbFJvb3RHMi5jcmwwgc4GA1Ud
IASBxjCBwzCBwAYEVR0gADCBtzAoBggrBgEFBQcCARYcaHR0cHM6Ly93d3cuZGln
aWNlcnQuY29tL0NQUzCBigYIKwYBBQUHAgIwfgx8QW55IHVzZSBvZiB0aGlzIENl
cnRpZmljYXRlIGNvbnN0aXR1dGVzIGFjY2VwdGFuY2Ugb2YgdGhlIFJlbHlpbmcg
UGFydHkgQWdyZWVtZW50IGxvY2F0ZWQgYXQgaHR0cHM6Ly93d3cuZGlnaWNlcnQu
Y29tL3JwYS11YTANBgkqhkiG9w0BAQsFAAOCAQEADf6H4Rxu128yUZjIR/vkWEv0
PnYKWUQkGXwbqaioq/zjVC/+LrIwBKwEH2aBQyBO/uvMwF+4PSThZKinw51Pfp8w
PcaBMeuQR1PEIYjQN3+NjRpMIFb2CUtbnMsHybKqKSW3rsTT2r2EIBtDrZ3EfeEg
5bfneV6PKvGqcQRZq5BBNTH11tedhcLJB+7F7GRVz8cKCTz0INpQk8gthjBND1+7
CR+ZHTWPqdXGRqiFJ09NTVmohj4dgL3VyOIArpmPbigODbp80fxzoQJepkHFef6q
mniWnIl7HApTrTIV/UN22gbCcL5dosGv69wnOVQomQsalcOkujetA5BruTiB9A==
-----END CERTIFICATE-----
`

const testCisMtlsCertFingerprint = "MD5 Fingerprint=73:64:72:C4:BF:0C:1C:EF:FC:86:A2:64:4E:5D:5A:EE"

func TestCisMtlsCertFingerprintOf(t *testing.T) {
	if got := cisMtlsCertFingerprintOf(testCisMtlsCertPEM); got != testCisMtlsCertFingerprint {
		t.Errorf("got %q, want %q", got, testCisMtlsCertFingerprint)
	}
	if got := cisMtlsCertFingerprintOf("not a pem"); got != "" {
		t.Errorf("invalid PEM: got %q, want empty", got)
	}
}

func TestCisMtlsCertDiffSuppress(t *testing.T) {
	suppress := ResourceIBMCISMtls().Schema[cisMtlsCert].DiffSuppressFunc
	cases := []struct {
		name, id, fingerprint, old, new string
		want                            bool
	}{
		{"imported, same cert", "id", testCisMtlsCertFingerprint, "", testCisMtlsCertPEM, true},
		{"imported, other cert", "id", "MD5 Fingerprint=00:11", "", testCisMtlsCertPEM, false},
		{"imported, no fingerprint", "id", "", "", "not a pem", false},
		{"managed, changed cert", "id", testCisMtlsCertFingerprint, "old", testCisMtlsCertPEM, false},
		{"create", "", "", "", testCisMtlsCertPEM, false},
	}
	for _, c := range cases {
		d := schema.TestResourceDataRaw(t, ResourceIBMCISMtls().Schema, map[string]interface{}{})
		d.SetId(c.id)
		d.Set(cisMtlsCertFingerprint, c.fingerprint)
		if got := suppress(cisMtlsCert, c.old, c.new, d); got != c.want {
			t.Errorf("%s: got %v, want %v", c.name, got, c.want)
		}
	}
}
