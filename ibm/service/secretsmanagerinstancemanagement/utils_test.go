// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanagerinstancemanagement

import (
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestReplaceEndpointRegion(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		region   string
		want     string
	}{
		{"public endpoint", "https://us-south.secrets-manager.cloud.ibm.com", "eu-gb", "https://eu-gb.secrets-manager.cloud.ibm.com"},
		{"test environment endpoint", "https://us-south.secrets-manager.test.cloud.ibm.com", "au-syd", "https://au-syd.secrets-manager.test.cloud.ibm.com"},
		{"same region", "https://eu-gb.secrets-manager.cloud.ibm.com", "eu-gb", "https://eu-gb.secrets-manager.cloud.ibm.com"},
		{"path is preserved", "https://us-south.secrets-manager.cloud.ibm.com/v2", "eu-gb", "https://eu-gb.secrets-manager.cloud.ibm.com/v2"},
		{"no scheme is returned unchanged", "us-south.secrets-manager.cloud.ibm.com", "eu-gb", "us-south.secrets-manager.cloud.ibm.com"},
		{"no dot in host is returned unchanged", "https://localhost", "eu-gb", "https://localhost"},
		{"empty endpoint is returned unchanged", "", "eu-gb", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ReplaceEndpointRegion(tt.endpoint, tt.region))
		})
	}
}

func newTestClient(t *testing.T, url string) *secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2 {
	t.Helper()
	client, err := secretsmanagerinstancemanagementv2.NewSecretsManagerInstanceManagementV2(
		&secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2Options{
			URL:           url,
			Authenticator: &core.NoAuthAuthenticator{},
		})
	assert.NoError(t, err)
	assert.NotNil(t, client)
	return client
}

// newTestResourceData returns ResourceData with a region schema; region is set only when it is not empty.
func newTestResourceData(t *testing.T, region string) *schema.ResourceData {
	t.Helper()
	s := map[string]*schema.Schema{
		"region": {Type: schema.TypeString, Optional: true, Computed: true},
	}
	raw := map[string]interface{}{}
	if region != "" {
		raw["region"] = region
	}
	return schema.TestResourceDataRaw(t, s, raw)
}

func TestGetClientWithRegion_NoRegionReturnsOriginalClient(t *testing.T) {
	original := newTestClient(t, "https://us-south.secrets-manager.cloud.ibm.com")

	got := getClientWithRegion(original, newTestResourceData(t, ""))

	assert.Same(t, original, got)
	assert.Equal(t, "https://us-south.secrets-manager.cloud.ibm.com", got.Service.GetServiceURL())
}

func TestGetClientWithRegion_RegionReturnsClonedClient(t *testing.T) {
	original := newTestClient(t, "https://us-south.secrets-manager.cloud.ibm.com")

	got := getClientWithRegion(original, newTestResourceData(t, "eu-gb"))

	assert.NotSame(t, original, got)
	assert.Equal(t, "https://eu-gb.secrets-manager.cloud.ibm.com", got.Service.GetServiceURL())
	// the shared provider client must not be changed, other resources still use it
	assert.Equal(t, "https://us-south.secrets-manager.cloud.ibm.com", original.Service.GetServiceURL())
}

func TestGetClientWithRegion_RegionsAreIndependent(t *testing.T) {
	original := newTestClient(t, "https://us-south.secrets-manager.cloud.ibm.com")

	london := getClientWithRegion(original, newTestResourceData(t, "eu-gb"))
	sydney := getClientWithRegion(original, newTestResourceData(t, "au-syd"))

	assert.Equal(t, "https://eu-gb.secrets-manager.cloud.ibm.com", london.Service.GetServiceURL())
	assert.Equal(t, "https://au-syd.secrets-manager.cloud.ibm.com", sydney.Service.GetServiceURL())
	assert.Equal(t, "https://us-south.secrets-manager.cloud.ibm.com", original.Service.GetServiceURL())
}

func TestGetRegion(t *testing.T) {
	tests := []struct {
		name       string
		clientURL  string
		configured string
		want       string
	}{
		{"configured region wins", "https://us-south.secrets-manager.cloud.ibm.com", "eu-gb", "eu-gb"},
		{"default from provider URL", "https://au-syd.secrets-manager.cloud.ibm.com", "", "au-syd"},
		{"default from test environment URL", "https://eu-de.secrets-manager.test.cloud.ibm.com", "", "eu-de"},
		{"URL without scheme gives empty region", "us-south.secrets-manager.cloud.ibm.com", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := newTestClient(t, tt.clientURL)
			assert.Equal(t, tt.want, getRegion(client, newTestResourceData(t, tt.configured)))
		})
	}
}

// The region that getRegion reports must be the one the client built by getClientWithRegion talks to.
func TestGetRegion_MatchesClientWithRegion(t *testing.T) {
	original := newTestClient(t, "https://us-south.secrets-manager.cloud.ibm.com")

	for _, region := range []string{"", "eu-gb"} {
		d := newTestResourceData(t, region)
		client := getClientWithRegion(original, d)

		want := region
		if want == "" {
			want = "us-south"
		}
		assert.Equal(t, want, getRegion(client, d))
	}
}
