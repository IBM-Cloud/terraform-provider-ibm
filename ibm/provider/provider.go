// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	"sync"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/drautomationservice"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// RetryDelay
const RetryAPIDelay = 5 * time.Second

// Provider returns a *schema.Provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureFunc: providerConfigure,

		DataSourcesMap: map[string]*schema.Resource{
			// Add new data sources here
			"ibm_pdr_get_deployment_status":   drautomationservice.DataSourceIbmPdrGetDeploymentStatus(),
			"ibm_pdr_get_event":               drautomationservice.DataSourceIbmPdrGetEvent(),
			"ibm_pdr_get_events":              drautomationservice.DataSourceIbmPdrGetEvents(),
			"ibm_pdr_get_machine_types":       drautomationservice.DataSourceIbmPdrGetMachineTypes(),
			"ibm_pdr_get_managed_vm_list":     drautomationservice.DataSourceIbmPdrGetManagedVmList(),
			"ibm_pdr_last_operation":          drautomationservice.DataSourceIbmPdrLastOperation(),
			"ibm_pdr_validate_clustertype":    drautomationservice.DataSourceIbmPdrValidateClustertype(),
			"ibm_pdr_validate_proxyip":        drautomationservice.DataSourceIbmPdrValidateProxyip(),
			"ibm_pdr_validate_workspace":      drautomationservice.DataSourceIbmPdrValidateWorkspace(),
			"ibm_pdr_get_dr_summary_response": drautomationservice.DataSourceIbmPdrGetDrSummaryResponse(),
		},

		ResourcesMap: map[string]*schema.Resource{
			// Add new resources here
			"ibm_pdr_managedr":        drautomationservice.ResourceIbmPdrManagedr(),
			"ibm_pdr_validate_apikey": drautomationservice.ResourceIbmPdrValidateApikey(),
		},

		Schema: map[string]*schema.Schema{
			"ibmcloud_api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IBM Cloud API Key",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_API_KEY", "IBMCLOUD_API_KEY"}, nil),
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IBM cloud Region (for example 'us-south').",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_REGION", "IBMCLOUD_REGION", "BM_REGION", "BLUEMIX_REGION"}, "us-south"),
			},
			"resource_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Resource group id.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_RESOURCE_GROUP", "IBMCLOUD_RESOURCE_GROUP", "BM_RESOURCE_GROUP", "BLUEMIX_RESOURCE_GROUP"}, ""),
			},
			"ibmcloud_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The timeout (in seconds) to set for any IBM Cloud API calls made.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_TIMEOUT", "IBMCLOUD_TIMEOUT"}, 60),
			},
			"max_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The retry count to set for API calls.",
				DefaultFunc: schema.EnvDefaultFunc("MAX_RETRIES", 10),
			},
		},
	}
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		return Provider()
	}
}

var globalValidatorDict validate.ValidatorDict
var initOnce sync.Once

func init() {
	validate.SetValidatorDict(Validator())
}

// Validator return validator
func Validator() validate.ValidatorDict {
	initOnce.Do(func() {
		globalValidatorDict = validate.ValidatorDict{
			ResourceValidatorDictionary: map[string]*validate.ResourceValidator{
				// Add new resource validators here
				"ibm_pdr_managedr": drautomationservice.ResourceIbmPdrManagedrValidator(),
			},
			DataSourceValidatorDictionary: map[string]*validate.ResourceValidator{},
		}
	})
	return globalValidatorDict
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	bluemixAPIKey := d.Get("ibmcloud_api_key").(string)
	region := d.Get("region").(string)
	resourceGrp := d.Get("resource_group").(string)
	bluemixTimeout := d.Get("ibmcloud_timeout").(int)
	retryCount := d.Get("max_retries").(int)

	var zone string
	if v, ok := d.GetOk("zone"); ok {
		zone = v.(string)
	}
	var visibility string
	if v, ok := d.GetOk("visibility"); ok {
		visibility = v.(string)
	}

	config := conns.Config{
		BluemixAPIKey:  bluemixAPIKey,
		Region:         region,
		ResourceGroup:  resourceGrp,
		BluemixTimeout: time.Duration(bluemixTimeout) * time.Second,
		RetryCount:     retryCount,
		RetryDelay:     RetryAPIDelay,
		Zone:           zone,
		Visibility:     visibility,
	}

	return config.ClientSession()
}
