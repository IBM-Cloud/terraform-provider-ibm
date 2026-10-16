// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	cloudMonitoringAPITokenDataSource = "(Data) ibm_cloud_monitoring_api_token"
	// CloudMonitoringEndpointEnv overrides the computed IBM Cloud Monitoring endpoint (for example for test or staging environments).
	CloudMonitoringEndpointEnv = "IBMCLOUD_MONITORING_API_ENDPOINT"
)

func DataSourceIBMCloudMonitoringAPIToken() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMCloudMonitoringAPITokenRead,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GUID or CRN of the IBM Cloud Monitoring instance.",
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The region of the IBM Cloud Monitoring instance. Defaults to the provider region.",
			},
			"endpoint_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "public",
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"public", "private"}),
				Description:  "The type of IBM Cloud Monitoring endpoint to call. Allowed values are public and private.",
			},
			"value": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "The IBM Cloud Monitoring API token of the IBM Cloud Monitoring instance.",
			},
		},
	}
}

type cloudMonitoringAPITokenResponse struct {
	Token *struct {
		Key string `json:"key"`
	} `json:"token"`
}

// CloudMonitoringInstanceGUID returns the instance GUID from a GUID or a resource instance CRN.
func CloudMonitoringInstanceGUID(instanceID string) (string, error) {
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

// CloudMonitoringAPITokenURL builds the URL of the IBM Cloud Monitoring API token endpoint.
func CloudMonitoringAPITokenURL(region, endpointType string) (string, error) {
	if endpoint := strings.TrimRight(strings.TrimSpace(os.Getenv(CloudMonitoringEndpointEnv)), "/"); endpoint != "" {
		return endpoint + "/api/token", nil
	}
	if region == "" {
		return "", fmt.Errorf("region must be set on the data source or the provider")
	}
	if endpointType == "private" {
		return fmt.Sprintf("https://private.%s.monitoring.cloud.ibm.com/api/token", region), nil
	}
	return fmt.Sprintf("https://%s.monitoring.cloud.ibm.com/api/token", region), nil
}

func dataSourceIBMCloudMonitoringAPITokenRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	bmxSess, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), cloudMonitoringAPITokenDataSource, "read", "initialize-client").GetDiag()
	}

	guid, err := CloudMonitoringInstanceGUID(d.Get("instance_id").(string))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), cloudMonitoringAPITokenDataSource, "read", "parse-instance_id").GetDiag()
	}

	region := d.Get("region").(string)
	if region == "" {
		region = bmxSess.Config.Region
	}
	url, err := CloudMonitoringAPITokenURL(region, d.Get("endpoint_type").(string))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), cloudMonitoringAPITokenDataSource, "read", "build-url").GetDiag()
	}

	iamToken := bmxSess.Config.IAMAccessToken
	if iamToken == "" {
		err = fmt.Errorf("IAM access token is not available; configure the provider with an IBM Cloud API key or IAM token")
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), cloudMonitoringAPITokenDataSource, "read", "iam-token").GetDiag()
	}
	if !strings.HasPrefix(strings.ToLower(iamToken), "bearer ") {
		iamToken = "Bearer " + iamToken
	}

	token, err := fetchCloudMonitoringAPIToken(context, url, iamToken, guid)
	if err != nil {
		return flex.TerraformErrorf(err, fmt.Sprintf("Error fetching IBM Cloud Monitoring API token: %s", err), cloudMonitoringAPITokenDataSource, "read").GetDiag()
	}

	d.SetId(guid)
	if err = d.Set("region", region); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting region: %s", err), cloudMonitoringAPITokenDataSource, "read", "set-region").GetDiag()
	}
	if err = d.Set("value", token); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting value: %s", err), cloudMonitoringAPITokenDataSource, "read", "set-value").GetDiag()
	}
	return nil
}

func fetchCloudMonitoringAPIToken(ctx context.Context, url, authorization, guid string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", authorization)
	req.Header.Set("IBMInstanceID", guid)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GET %s returned status %d: %s", url, resp.StatusCode, strings.TrimSpace(string(body)))
	}
	return ParseCloudMonitoringAPITokenResponse(body)
}

// ParseCloudMonitoringAPITokenResponse extracts the token key from the /api/token response body.
func ParseCloudMonitoringAPITokenResponse(body []byte) (string, error) {
	var parsed cloudMonitoringAPITokenResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return "", fmt.Errorf("unable to parse API token response: %s", err)
	}
	if parsed.Token == nil || parsed.Token.Key == "" {
		return "", fmt.Errorf("API token response did not contain token.key")
	}
	return parsed.Token.Key, nil
}
