// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamaccessgroup

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMIAMAccessGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIAMAccessGroupCreate,
		ReadContext:   resourceIBMIAMAccessGroupRead,
		UpdateContext: resourceIBMIAMAccessGroupUpdate,
		DeleteContext: resourceIBMIAMAccessGroupDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the access group",
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the access group",
			},

			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of the access group",
			},
		},
	}
}

func resourceIBMIAMAccessGroupCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamAccessGroupsClient, err := meta.(conns.ClientSession).IAMAccessGroupsV2()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_iam_access_group", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(err)
	}

	name := d.Get("name").(string)
	creatAccessGroupOptions := iamAccessGroupsClient.NewCreateAccessGroupOptions(userDetails.UserAccount, name)
	if des, ok := d.GetOk("description"); ok {
		description := des.(string)
		creatAccessGroupOptions.Description = &description
	}
	group, _, err := iamAccessGroupsClient.CreateAccessGroupWithContext(context, creatAccessGroupOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateAccessGroupWithContext failed: %s", err.Error()), "ibm_iam_access_group", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*group.ID)

	return resourceIBMIAMAccessGroupRead(context, d, meta)
}

func resourceIBMIAMAccessGroupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamAccessGroupsClient, err := meta.(conns.ClientSession).IAMAccessGroupsV2()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_iam_access_group", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	agrpID := d.Id()
	getAccessGroupOptions := iamAccessGroupsClient.NewGetAccessGroupOptions(agrpID)
	getAccessGroupOptions.SetShowCRN(true)
	var agrp *iamaccessgroupsv2.Group
	var detailedResponse *core.DetailedResponse
	err = resource.RetryContext(context, 5*time.Second, func() *resource.RetryError {
		agrp, detailedResponse, err = iamAccessGroupsClient.GetAccessGroup(getAccessGroupOptions)
		if err != nil || agrp == nil {
			if detailedResponse != nil && detailedResponse.StatusCode == 404 {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})

	if conns.IsResourceTimeoutError(err) {
		agrp, detailedResponse, err = iamAccessGroupsClient.GetAccessGroupWithContext(context, getAccessGroupOptions)

	}
	if err != nil || agrp == nil {
		if detailedResponse != nil && detailedResponse.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAccessGroupWithContext failed: %s", err.Error()), "ibm_iam_access_group", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	version := detailedResponse.GetHeaders().Get("etag")
	d.Set("name", agrp.Name)
	d.Set("description", agrp.Description)
	d.Set("version", version)
	d.Set("crn", agrp.CRN)
	return nil
}

func resourceIBMIAMAccessGroupUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamAccessGroupsClient, err := meta.(conns.ClientSession).IAMAccessGroupsV2()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_iam_access_group", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	agrpID := d.Id()

	hasChange := false
	version := d.Get("version").(string)
	updateAccessGroupOptions := iamAccessGroupsClient.NewUpdateAccessGroupOptions(agrpID, version)

	if d.HasChange("name") {
		name := d.Get("name").(string)
		updateAccessGroupOptions.Name = &name
		hasChange = true
	}

	if d.HasChange("description") {
		description := d.Get("description").(string)
		updateAccessGroupOptions.Description = &description
		hasChange = true
	}

	if hasChange {
		_, _, err = iamAccessGroupsClient.UpdateAccessGroupWithContext(context, updateAccessGroupOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateAccessGroupWithContext failed: %s", err.Error()), "ibm_iam_access_group", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMIAMAccessGroupRead(context, d, meta)

}

func resourceIBMIAMAccessGroupDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamAccessGroupsClient, err := meta.(conns.ClientSession).IAMAccessGroupsV2()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_iam_access_group", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	agID := d.Id()
	force := true
	deleteAccessGroupOptions := iamAccessGroupsClient.NewDeleteAccessGroupOptions(agID)
	deleteAccessGroupOptions.SetForce(force)

	_, err = iamAccessGroupsClient.DeleteAccessGroupWithContext(context, deleteAccessGroupOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteAccessGroupWithContext failed: %s", err.Error()), "ibm_iam_access_group", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")

	return nil
}
