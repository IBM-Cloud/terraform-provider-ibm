// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	sccWorkloadProtectionAPITokenDataSource = "(Data) ibm_scc_workload_protection_api_token"
	// SccWorkloadProtectionEndpointEnv overrides the Workload Protection endpoint built in conns/config.go.
	SccWorkloadProtectionEndpointEnv = "IBMCLOUD_SCC_WORKLOAD_PROTECTION_API_ENDPOINT"
)

func DataSourceIbmSccWorkloadProtectionAPIToken() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSccWorkloadProtectionAPITokenRead,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GUID or CRN of the Security and Compliance Center Workload Protection instance.",
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The region of the Security and Compliance Center Workload Protection instance. Defaults to the provider region.",
			},
			"endpoint_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"public", "private"}),
				Description:  "The type of Security and Compliance Center Workload Protection endpoint to call. Allowed values are public and private. Defaults to the provider visibility setting.",
			},
			"value": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "The API token of the Security and Compliance Center Workload Protection instance.",
			},
		},
	}
}

type sccWorkloadProtectionAPITokenResponse struct {
	Token *struct {
		Key string `json:"key"`
	} `json:"token"`
}

// SccWorkloadProtectionInstanceGUID returns the instance GUID from a GUID or a resource instance CRN.
func SccWorkloadProtectionInstanceGUID(instanceID string) (string, error) {
	instanceID = strings.TrimSpace(instanceID)
	if !strings.HasPrefix(instanceID, "crn:") {
		if instanceID == "" {
			return "", fmt.Errorf("instance_id must not be empty")
		}
		return instanceID, nil
	}
	// crn:version:cname:ctype:service-name:location:scope:service-instance:resource-type:resource
	parts := strings.Split(instanceID, ":")
	if len(parts) < 8 || parts[7] == "" {
		return "", fmt.Errorf("instance_id %q is not a valid service instance CRN", instanceID)
	}
	return parts[7], nil
}

// SccWorkloadProtectionServiceURL builds the base URL of the Workload Protection instance API for a region and endpoint type.
func SccWorkloadProtectionServiceURL(region, endpointType string) (string, error) {
	if region == "" {
		return "", fmt.Errorf("region must be set on the data source or the provider")
	}
	if endpointType == "private" {
		return conns.ContructEndpoint(fmt.Sprintf("private.%s.security-compliance-secure", region), "cloud.ibm.com"), nil
	}
	return conns.ContructEndpoint(fmt.Sprintf("%s.security-compliance-secure", region), "cloud.ibm.com"), nil
}

func dataSourceIbmSccWorkloadProtectionAPITokenRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	workloadProtectionClient, err := meta.(conns.ClientSession).SccWorkloadProtectionV1()
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), sccWorkloadProtectionAPITokenDataSource, "read", "initialize-client").GetDiag()
	}
	bmxSess, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), sccWorkloadProtectionAPITokenDataSource, "read", "initialize-client").GetDiag()
	}

	guid, err := SccWorkloadProtectionInstanceGUID(d.Get("instance_id").(string))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), sccWorkloadProtectionAPITokenDataSource, "read", "parse-instance_id").GetDiag()
	}

	region := d.Get("region").(string)
	if region == "" {
		region = bmxSess.Config.Region
	}

	// Work on a copy so the shared session client keeps the URL built in conns/config.go.
	service := workloadProtectionClient.Clone()
	endpointType, endpointTypeSet := d.GetOk("endpoint_type")
	_, regionSet := d.GetOk("region")
	if (regionSet || endpointTypeSet) && os.Getenv(SccWorkloadProtectionEndpointEnv) == "" {
		if !endpointTypeSet {
			endpointType = "public"
			if strings.Contains(service.GetServiceURL(), "://private.") {
				endpointType = "private"
			}
		}
		serviceURL, err := SccWorkloadProtectionServiceURL(region, endpointType.(string))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), sccWorkloadProtectionAPITokenDataSource, "read", "build-url").GetDiag()
		}
		if err = service.SetServiceURL(serviceURL); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), sccWorkloadProtectionAPITokenDataSource, "read", "set-service-url").GetDiag()
		}
	}

	token, err := FetchSccWorkloadProtectionAPIToken(context, service, guid)
	if err != nil {
		return flex.TerraformErrorf(err, fmt.Sprintf("Error fetching Security and Compliance Center Workload Protection API token: %s", err), sccWorkloadProtectionAPITokenDataSource, "read").GetDiag()
	}

	d.SetId(guid)
	if err = d.Set("region", region); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting region: %s", err), sccWorkloadProtectionAPITokenDataSource, "read", "set-region").GetDiag()
	}
	if err = d.Set("value", token); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting value: %s", err), sccWorkloadProtectionAPITokenDataSource, "read", "set-value").GetDiag()
	}
	return nil
}

// FetchSccWorkloadProtectionAPIToken calls GET /api/token on the Workload Protection instance API.
func FetchSccWorkloadProtectionAPIToken(ctx context.Context, service *core.BaseService, guid string) (string, error) {
	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = service.GetEnableGzipCompression()
	if _, err := builder.ResolveRequestURL(service.GetServiceURL(), "/api/token", nil); err != nil {
		return "", err
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("IBMInstanceID", guid)

	request, err := builder.Build()
	if err != nil {
		return "", err
	}

	var rawResponse json.RawMessage
	response, err := service.Request(request, &rawResponse)
	if err != nil {
		if response != nil {
			return "", fmt.Errorf("GET %s/api/token returned status %d: %s", service.GetServiceURL(), response.StatusCode, err)
		}
		return "", err
	}
	return ParseSccWorkloadProtectionAPITokenResponse(rawResponse)
}

// ParseSccWorkloadProtectionAPITokenResponse extracts the token key from the /api/token response body.
func ParseSccWorkloadProtectionAPITokenResponse(body []byte) (string, error) {
	var parsed sccWorkloadProtectionAPITokenResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return "", fmt.Errorf("unable to parse API token response: %s", err)
	}
	if parsed.Token == nil || parsed.Token.Key == "" {
		return "", fmt.Errorf("API token response did not contain token.key")
	}
	return parsed.Token.Key, nil
}
