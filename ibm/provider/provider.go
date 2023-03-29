// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/project"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
)

// RetryDelay
const RetryAPIDelay = 5 * time.Second

// Provider returns a *schema.Provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureFunc: providerConfigure,

		DataSourcesMap: map[string]*schema.Resource{
			// Add new data sources here
			"ibm_project": project.DataSourceIbmProject(),
			"ibm_event_notification": project.DataSourceIbmEventNotification(),
		},

		ResourcesMap: map[string]*schema.Resource{
			// Add new resources here
			"ibm_project": project.ResourceIbmProject(),
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
				"ibm_project": project.ResourceIbmProjectValidator(),
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
