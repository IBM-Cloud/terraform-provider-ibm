// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMISSSHKey() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISSSHKeyRead,

		Schema: map[string]*schema.Schema{
			"resource_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Resource group ID",
			},
			"id": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"name", "id"},
				Description:  "SSH key ID",
			},

			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "User Tags for the ssh",
			},

			isKeyName: {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"name", "id"},
				Description:  "The name of the ssh key",
			},
			// missing schema added
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the key was created.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this key.",
			},
			isKeyType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ssh key type",
			},

			isKeyFingerprint: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ssh key Fingerprint",
			},

			isKeyPublicKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SSH Public key data",
			},

			isKeyLength: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ssh key length",
			},

			flex.ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},

			flex.ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			flex.ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			IsKeyCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},

			isKeyAccessTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access tags",
			},
		},
	}
}

func dataSourceIBMISSSHKeyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := ""
	if nameOk, ok := d.GetOk(isKeyName); ok {
		name = nameOk.(string)
	}
	id := ""
	if idOk, ok := d.GetOk("id"); ok {
		id = idOk.(string)
	}

	diag := keyGetByNameOrId(context, d, meta, name, id)
	if diag != nil {
		return diag
	}
	return nil
}

func keyGetByNameOrId(context context.Context, d *schema.ResourceData, meta interface{}, name, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_ssh_key", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	var key vpcv1.Key

	if id != "" {
		getKeyOptions := &vpcv1.GetKeyOptions{
			ID: &id,
		}
		keyintf, _, err := sess.GetKeyWithContext(context, getKeyOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetKeyWithContext failed: %s", err.Error()), "(Data) ibm_is_ssh_key", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		key = *keyintf

	} else {
		listKeysOptions := &vpcv1.ListKeysOptions{}

		start := ""
		allrecs := []vpcv1.Key{}
		for {
			if start != "" {
				listKeysOptions.Start = &start
			}

			keys, _, err := sess.ListKeysWithContext(context, listKeysOptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListKeysWithContext failed: %s", err.Error()), "(Data) ibm_is_ssh_key", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			start = flex.GetNext(keys.Next)
			allrecs = append(allrecs, keys.Keys...)
			if start == "" {
				break
			}
		}
		found := false
		for _, keyintf := range allrecs {
			if *keyintf.Name == name {
				key = keyintf
				found = true
			}
		}
		if !found {
			err = fmt.Errorf("[ERROR] No SSH Key found with name %s", name)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Not found: %s", err.Error()), "(Data) ibm_is_ssh_key", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	d.SetId(*key.ID)
	if err = d.Set("name", key.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_ssh_key", "read", "set-name").GetDiag()
	}
	if err = d.Set("type", key.Type); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_is_ssh_key", "read", "set-type").GetDiag()
	}
	if err = d.Set("fingerprint", key.Fingerprint); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting fingerprint: %s", err), "(Data) ibm_is_ssh_key", "read", "set-fingerprint").GetDiag()
	}
	if err = d.Set("length", flex.IntValue(key.Length)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting length: %s", err), "(Data) ibm_is_ssh_key", "read", "set-length").GetDiag()
	}
	controller, err := flex.GetBaseController(meta)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_controller_url: %s", err), "(Data) ibm_is_ssh_key", "read", "set-resource_controller_url").GetDiag()
	}
	d.Set(flex.ResourceControllerURL, controller+"/vpc/compute/sshKeys")
	if err = d.Set("created_at", flex.DateTimeToString(key.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_ssh_key", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("href", key.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_ssh_key", "read", "set-href").GetDiag()
	}
	if err = d.Set(flex.ResourceName, key.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_name: %s", err), "(Data) ibm_is_ssh_key", "read", "set-resource_name").GetDiag()
	}
	if err = d.Set(flex.ResourceCRN, key.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_crn: %s", err), "(Data) ibm_is_ssh_key", "read", "set-resource_crn").GetDiag()
	}
	if err = d.Set("crn", key.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_ssh_key", "read", "set-crn").GetDiag()
	}
	if key.ResourceGroup != nil {
		d.Set(flex.ResourceGroupName, *key.ResourceGroup.ID)
	}
	if err = d.Set("public_key", key.PublicKey); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting public_key: %s", err), "(Data) ibm_is_ssh_key", "read", "set-public_key").GetDiag()
	}
	tags, err := flex.GetGlobalTagsUsingCRN(meta, *key.CRN, "", isUserTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource vpc ssh key (%s) tags: %s", d.Id(), err)
	}
	d.Set("tags", tags)
	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *key.CRN, "", isKeyAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource SSH Key (%s) access tags: %s", d.Id(), err)
	}
	d.Set(isKeyAccessTags, accesstags)
	return nil
}
