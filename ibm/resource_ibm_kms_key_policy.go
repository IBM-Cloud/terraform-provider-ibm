// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	kp "github.com/IBM/keyprotect-go-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIBMKmskeyPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMKmsKeyPolicyCreate,
		ReadContext:   resourceIBMKmsKeyPolicyRead,
		UpdateContext: resourceIBMKmsKeyPolicyUpdate,
		DeleteContext: resourceIBMKmsKeyPolicyDelete,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Key protect or hpcs instance GUID",
			},
			"key_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Key ID",
			},
			"endpoint_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"public", "private"}),
				Description:  "public or private",
				ForceNew:     true,
				Default:      "public",
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
			ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},
			ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},
			ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},
		},
	}
}
func resourceIBMKmsKeyPolicyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	kpAPI, err := meta.(ClientSession).keyManagementAPI()
	if err != nil {
		return diag.FromErr(err)
	}

	rContollerClient, err := meta.(ClientSession).ResourceControllerAPIV2()
	if err != nil {
		return diag.FromErr(err)
	}

	log.Println("Entering the Policies++++++++>>>>>>>>")

	instanceID := d.Get("instance_id").(string)
	endpointType := d.Get("endpoint_type").(string)
	key_id := d.Get("key_id").(string)

	rContollerApi := rContollerClient.ResourceServiceInstanceV2()

	instanceData, err := rContollerApi.GetInstance(instanceID)
	if err != nil {
		return diag.FromErr(err)
	}
	instanceCRN := instanceData.Crn.String()
	crnData := strings.Split(instanceCRN, ":")
	var hpcsEndpointURL string

	if crnData[4] == "hs-crypto" {
		hpcsEndpointAPI, err := meta.(ClientSession).HpcsEndpointAPI()
		if err != nil {
			return diag.FromErr(err)
		}

		resp, err := hpcsEndpointAPI.Endpoint().GetAPIEndpoint(instanceID)
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
		kpAPI.URL = u
	} else if crnData[4] == "kms" {
		if endpointType == "private" {
			if !strings.Contains(kpAPI.Config.BaseURL, "private") {
				kmsEndpURL := strings.SplitAfter(kpAPI.Config.BaseURL, "https://")
				if len(kmsEndpURL) == 2 {
					kmsEndpointURL := kmsEndpURL[0] + "private." + kmsEndpURL[1]
					u, err := url.Parse(kmsEndpointURL)
					if err != nil {
						return diag.Errorf("Error Parsing kms EndpointURL")
					}
					kpAPI.URL = u
				} else {
					return diag.Errorf("Error in Kms EndPoint URL ")
				}
			}
		}
	} else {
		return diag.Errorf("Invalid or unsupported service Instance")
	}
	kpAPI.Config.InstanceID = instanceID

	var expiration *time.Time
	if es, ok := d.GetOk("expiration_date"); ok {
		expiration_string := es.(string)
		// parse string to required time format
		expiration_time, err := time.Parse(time.RFC3339, expiration_string)
		if err != nil {
			return diag.Errorf("Invalid time format (the date format follows RFC 3339): %s", err)
		}
		expiration = &expiration_time
	} else {
		expiration = nil
	}
	log.Println("expiration", expiration)
	key, err := kpAPI.GetKey(context, key_id)
	if err != nil {
		return diag.Errorf("Get Key failed with error: %s", err)
	}
	err = resourceHandlePolicies(context, d, kpAPI, meta, key_id)
	if err != nil {
		return diag.Errorf("Could not create policies: %s", err)
	}
	d.SetId(key.CRN)
	return resourceIBMKmsKeyPolicyUpdate(context, d, meta)
}

func resourceIBMKmsKeyPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	kpAPI, err := meta.(ClientSession).keyManagementAPI()
	if err != nil {
		return diag.FromErr(err)
	}
	crn := d.Id()
	crnData := strings.Split(crn, ":")
	endpointType := crnData[3]
	instanceID := crnData[len(crnData)-3]
	keyid := crnData[len(crnData)-1]

	var instanceType string
	var hpcsEndpointURL string
	log.Println("instancetype", instanceType)
	if crnData[4] == "hs-crypto" {
		instanceType = "hs-crypto"
		hpcsEndpointAPI, err := meta.(ClientSession).HpcsEndpointAPI()
		if err != nil {
			return diag.FromErr(err)
		}

		resp, err := hpcsEndpointAPI.Endpoint().GetAPIEndpoint(instanceID)
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
		kpAPI.URL = u
	} else if crnData[4] == "kms" {
		instanceType = "kms"
		if endpointType == "private" {
			if !strings.Contains(kpAPI.Config.BaseURL, "private") {
				kmsEndpURL := strings.SplitAfter(kpAPI.Config.BaseURL, "https://")
				if len(kmsEndpURL) == 2 {
					kmsEndpointURL := kmsEndpURL[0] + "private." + kmsEndpURL[1]
					u, err := url.Parse(kmsEndpointURL)
					if err != nil {
						return diag.Errorf("Error Parsing kms EndpointURL")
					}
					kpAPI.URL = u
				} else {
					return diag.Errorf("Error in Kms EndPoint URL ")
				}
			}
		}
	} else {
		return diag.Errorf("Invalid or unsupported service Instance")
	}

	kpAPI.Config.InstanceID = instanceID
	// keyid := d.Id()
	key, err := kpAPI.GetKey(context, keyid)
	if err != nil {
		return diag.Errorf("Get Key failed with error: %s", err)
	}

	policies, err := kpAPI.GetPolicies(context, keyid)

	if err != nil {
		return diag.Errorf("Failed to read policies: %s", err)
	}
	if len(policies) == 0 {
		log.Printf("No Policy Configurations read\n")
	} else {
		d.Set("policies", flattenKeyPolicies(policies))
	}
	d.Set("instance_id", instanceID)
	d.Set("key_id", keyid)
	d.Set("endpoint_type", endpointType)
	if key.Expiration != nil {
		expiration := key.Expiration
		d.Set("expiration_date", expiration.Format(time.RFC3339))
	} else {
		d.Set("expiration_date", "")
	}
	d.Set(ResourceName, key.Name)
	d.Set(ResourceCRN, key.CRN)
	state := key.State
	d.Set(ResourceStatus, strconv.Itoa(state))
	rcontroller, err := getBaseController(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	id := key.ID
	crn1 := strings.TrimSuffix(key.CRN, ":key:"+id)

	d.Set(ResourceControllerURL, rcontroller+"/services/kms/"+url.QueryEscape(crn1)+"%3A%3A")

	return nil

}

func resourceIBMKmsKeyPolicyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	if d.HasChange("policies") {

		kpAPI, err := meta.(ClientSession).keyManagementAPI()
		if err != nil {
			return diag.FromErr(err)
		}

		rContollerClient, err := meta.(ClientSession).ResourceControllerAPIV2()
		if err != nil {
			return diag.FromErr(err)
		}

		instanceID := d.Get("instance_id").(string)
		endpointType := d.Get("endpoint_type").(string)

		rContollerApi := rContollerClient.ResourceServiceInstanceV2()

		instanceData, err := rContollerApi.GetInstance(instanceID)
		if err != nil {
			return diag.FromErr(err)
		}
		instanceCRN := instanceData.Crn.String()
		crnData := strings.Split(instanceCRN, ":")

		var hpcsEndpointURL string

		if crnData[4] == "hs-crypto" {
			hpcsEndpointAPI, err := meta.(ClientSession).HpcsEndpointAPI()
			if err != nil {
				return diag.FromErr(err)
			}

			resp, err := hpcsEndpointAPI.Endpoint().GetAPIEndpoint(instanceID)
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
			kpAPI.URL = u
		} else if crnData[4] == "kms" {
			if endpointType == "private" {
				if !strings.Contains(kpAPI.Config.BaseURL, "private") {
					kmsEndpURL := strings.SplitAfter(kpAPI.Config.BaseURL, "https://")
					if len(kmsEndpURL) == 2 {
						kmsEndpointURL := kmsEndpURL[0] + "private." + kmsEndpURL[1]
						u, err := url.Parse(kmsEndpointURL)
						if err != nil {
							return diag.Errorf("Error Parsing kms EndpointURL")
						}
						kpAPI.URL = u
					} else {
						return diag.Errorf("Error in Kms EndPoint URL ")
					}
				}
			}
		} else {
			return diag.Errorf("Invalid or unsupported service Instance")
		}

		kpAPI.Config.InstanceID = instanceID

		crn := d.Id()
		crnData = strings.Split(crn, ":")
		key_id := crnData[len(crnData)-1]

		err = resourceHandlePolicies(context, d, kpAPI, meta, key_id)
		if err != nil {
			resourceIBMKmsKeyRead(d, meta)
			return diag.Errorf("Could not create policies: %s", err)
		}
	}
	return resourceIBMKmsKeyPolicyRead(context, d, meta)

}

func resourceIBMKmsKeyPolicyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	//Do not support delete Policies
	log.Println("Warning:  `terraform destroy` does not remove the policies of the Key but only clears the state file. Key Policies get deleted when the associated key resource is destroyed.")
	d.SetId("")
	return nil

}

func resourceHandlePolicies(context context.Context, d *schema.ResourceData, kpAPI *kp.Client, meta interface{}, key_id string) error {
	var setRotation, setDualAuthDelete, dualAuthEnable bool
	var rotationInterval int

	if policyInfo, ok := d.GetOk("policies"); ok {

		policyDataList := policyInfo.([]interface{})
		policyData := policyDataList[0].(map[string]interface{})

		if rpd, ok := policyData["rotation"]; ok {
			rpdList := rpd.([]interface{})
			if len(rpdList) != 0 {
				rotationInterval = rpdList[0].(map[string]interface{})["interval_month"].(int)
				setRotation = true
			}
		}
		if dadp, ok := policyData["dual_auth_delete"]; ok {
			dadpList := dadp.([]interface{})
			if len(dadpList) != 0 {
				dualAuthEnable = dadpList[0].(map[string]interface{})["enabled"].(bool)
				setDualAuthDelete = true
			}
		}

		_, err := kpAPI.SetPolicies(context, key_id, setRotation, rotationInterval, setDualAuthDelete, dualAuthEnable)
		if err != nil {
			return fmt.Errorf("Error while creating policies: %s", err)
		}
	}
	return nil
}
