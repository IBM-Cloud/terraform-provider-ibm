// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMKmsInstancePolicies() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceIBMKmsInstancePolicyRead,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Key protect or hpcs instance GUID or CRN",
				DiffSuppressFunc: suppressKMSInstanceIDDiff,
			},
			"dual_auth_delete": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Data associated with the dual auth delete policy for instance",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If set to true, Key Protect enables a dual authorization policy for the instance.",
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
						"last_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Updates when the policy is replaced or modified. The date format follows RFC 3339.",
						},
					},
				},
			},
			"rotation": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Data associated with the rotation policy for instance",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If set to true, Key Protect enables a rotation policy for the instance",
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
						"last_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Updates when the policy is replaced or modified. The date format follows RFC 3339.",
						},
						"interval_month": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the rotation time interval in months for the instance",
						},
					},
				},
			},
			"key_create_import_access": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Data associated with the key create import access policy for the instance",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If set to true, Key Protect enables a KCIA policy for the instance.",
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
						"last_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Updates when the policy is replaced or modified. The date format follows RFC 3339.",
						},
						"create_root_key": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If set to true, Key Protect allows you or any authorized users to create root keys in the instance.",
						},
						"create_standard_key": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If set to true, Key Protect allows you or any authorized users to create standard keys in the instance.",
						},
						"import_root_key": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If set to true, Key Protect allows you or any authorized users to import root keys into the instance.",
						},
						"import_standard_key": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If set to true, Key Protect allows you or any authorized users to import standard keys into the instance.",
						},
						"enforce_token": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If set to true, the service prevents you or any authorized users from importing key material into the specified service instance without using an import token.",
						},
					},
				},
			},
			"metrics": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Data associated with the metric policy for the instance",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "If set to true, Key Protect enables a metrics policy on the instance.",
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
						"last_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Updates when the policy is replaced or modified. The date format follows RFC 3339.",
						},
					},
				},
			},
		},
	}
}

func resourceIBMKmsInstancePolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instanceID := getInstanceIDFromCRN(d.Get("instance_id").(string))
	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return diag.FromErr(err)
	}
	instancePolicies, err := kpAPI.GetInstancePolicies(context)
	if err != nil {
		return diag.Errorf("[ERROR] Error retrieving instance policies: %s", err)
	}
	d.Set("instance_id", instanceID)
	d.SetId(instanceID)
	d.Set("dual_auth_delete", flex.FlattenInstancePolicy("dual_auth_delete", instancePolicies))
	d.Set("rotation", flex.FlattenInstancePolicy("rotation", instancePolicies))
	d.Set("metrics", flex.FlattenInstancePolicy("metrics", instancePolicies))
	d.Set("key_create_import_access", flex.FlattenInstancePolicy("key_create_import_access", instancePolicies))
	return nil

}
