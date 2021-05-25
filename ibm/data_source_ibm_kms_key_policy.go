// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"log"
	"net/url"
	"strings"

	//kp "github.com/IBM/keyprotect-go-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMKMSkeyPolicies() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMKMSKeyPoliciesRead,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Key protect or hpcs instance GUID",
			},
			"endpoint_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"public", "private"}),
				Description:  "public or private",
				Default:      "public",
			},
			"key_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Key ID of the Key",
			},
			"policies": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Creates or updates one or more policies for the specified key",
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rotation": {
							Type:         schema.TypeList,
							Optional:     true,
							Computed:     true,
							AtLeastOneOf: []string{"policies.0.rotation", "policies.0.dual_auth_delete"},
							Description:  "Specifies the key rotation time interval in months, with a minimum of 1, and a maximum of 12",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The v4 UUID used to uniquely identify the policy resource, as specified by RFC 4122.",
									},
									"crn": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cloud Resource Name (CRN) that uniquely identifies your cloud resources.",
									},
									"created_by": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for the resource that created the policy.",
									},
									"creation_date": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The date the policy was created. The date format follows RFC 3339.",
									},
									"updated_by": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for the resource that updated the policy.",
									},
									"last_update_date": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Updates when the policy is replaced or modified. The date format follows RFC 3339.",
									},
									"interval_month": {
										Type:         schema.TypeInt,
										Required:     true,
										ValidateFunc: validateAllowedRangeInt(1, 12),
										Description:  "Specifies the key rotation time interval in months",
									},
								},
							},
						},
						"dual_auth_delete": {
							Type:         schema.TypeList,
							Optional:     true,
							Computed:     true,
							AtLeastOneOf: []string{"policies.0.rotation", "policies.0.dual_auth_delete"},
							Description:  "Data associated with the dual authorization delete policy.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The v4 UUID used to uniquely identify the policy resource, as specified by RFC 4122.",
									},
									"crn": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cloud Resource Name (CRN) that uniquely identifies your cloud resources.",
									},
									"created_by": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for the resource that created the policy.",
									},
									"creation_date": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The date the policy was created. The date format follows RFC 3339.",
									},
									"updated_by": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for the resource that updated the policy.",
									},
									"last_update_date": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Updates when the policy is replaced or modified. The date format follows RFC 3339.",
									},
									"enabled": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "If set to true, Key Protect enables a dual authorization policy on a single key.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMKMSKeyPoliciesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	api, err := meta.(ClientSession).keyManagementAPI()
	if err != nil {
		return diag.FromErr(err)
	}

	rContollerClient, err := meta.(ClientSession).ResourceControllerAPIV2()
	if err != nil {
		return diag.FromErr(err)
	}

	instanceID := d.Get("instance_id").(string)
	endpointType := d.Get("endpoint_type").(string)
	key_id := d.Get("key_id").(string)

	rContollerApi := rContollerClient.ResourceServiceInstanceV2()

	instanceData, err := rContollerApi.GetInstance(instanceID)
	if err != nil {
		return diag.FromErr(err)
	}
	instanceCRN := instanceData.Crn.String()

	var hpcsEndpointURL string
	crnData := strings.Split(instanceCRN, ":")

	if crnData[4] == "hs-crypto" {

		hpcsEndpointApi, err := meta.(ClientSession).HpcsEndpointAPI()
		if err != nil {
			return diag.FromErr(err)
		}
		resp, err := hpcsEndpointApi.Endpoint().GetAPIEndpoint(instanceID)
		if err != nil {
			return diag.FromErr(err)
		}

		if endpointType == "public" {
			hpcsEndpointURL = "https://" + resp.Kms.Public + "/api/v2/keys"
		} else {
			hpcsEndpointURL = "https://" + resp.Kms.Private + "/api/v2/keys"
		}

		u, err := url.Parse(hpcsEndpointURL)
		if err != nil {
			return diag.Errorf("Error Parsing hpcs EndpointURL")
		}
		api.URL = u
	} else if crnData[4] == "kms" {
		if endpointType == "private" {
			if !strings.HasPrefix(api.Config.BaseURL, "private") {
				api.Config.BaseURL = "private." + api.Config.BaseURL
			}
		}
	} else {
		return diag.Errorf("Invalid or unsupported service Instance")
	}

	api.Config.InstanceID = instanceID

	policies, err := api.GetPolicies(context, key_id)
	if err != nil {
		return diag.Errorf("Failed to read policies: %s", err)
	}

	if len(policies) == 0 {
		log.Printf("No Policy Configurations read\n")
	} else {
		d.Set("policies", flattenKeyPolicies(policies))
	}

	d.SetId(instanceID)
	d.Set("key_id", key_id)
	d.Set("instance_id", instanceID)
	d.Set("endpoint_type", endpointType)

	return nil

}
