// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceIBMPIPlacementGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIPlacementGroupRead,
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_PlacementGroupID: {
				AtLeastOneOf:  []string{Arg_PlacementGroupID, Arg_PlacementGroupName},
				ConflictsWith: []string{Arg_PlacementGroupName},
				Description:   "The placement group ID.",
				Optional:      true,
				Type:          schema.TypeString,
			},
			Arg_PlacementGroupName: {
				AtLeastOneOf:  []string{Arg_PlacementGroupID, Arg_PlacementGroupName},
				ConflictsWith: []string{Arg_PlacementGroupID},
				Deprecated:    "The pi_placement_group_name field is deprecated. Please use pi_placement_group_id instead.",
				Description:   "The name of the placement group.",
				Optional:      true,
				Type:          schema.TypeString,
			},

			// Attribute
			Attr_CRN: {
				Computed:    true,
				Description: "The CRN of this resource.",
				Type:        schema.TypeString,
			},
			Attr_Members: {
				Computed:    true,
				Description: "List of server instances IDs that are members of the placement group.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Type:        schema.TypeList,
			},
			Attr_Name: {
				Computed:    true,
				Description: "The name of the placement group.",
				Type:        schema.TypeString,
			},
			Attr_Policy: {
				Computed:    true,
				Description: "The value of the group's affinity policy. Valid values are affinity and anti-affinity.",
				Type:        schema.TypeString,
			},
			Attr_UserTags: {
				Computed:    true,
				Description: "List of user tags attached to the resource.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Type:        schema.TypeSet,
			},
		},
	}
}

func dataSourceIBMPIPlacementGroupRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "(Data) ibm_pi_placement_group", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	var placementGroupID string
	if v, ok := d.GetOk(Arg_PlacementGroupID); ok {
		placementGroupID = v.(string)
	} else if v, ok := d.GetOk(Arg_PlacementGroupName); ok {
		placementGroupID = v.(string)
	}

	client := instance.NewIBMPIPlacementGroupClient(ctx, sess, cloudInstanceID)
	placementGroup, err := client.Get(placementGroupID)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Get failed: %s", err.Error()), "(Data) ibm_pi_placement_group", "read")
		log.Printf("[DEBUG]  err \n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*placementGroup.ID)
	if placementGroup.Crn != "" {
		d.Set(Attr_CRN, placementGroup.Crn)
		userTags, err := flex.GetGlobalTagsUsingCRN(meta, string(placementGroup.Crn), "", UserTagType)
		if err != nil {
			log.Printf("Error on get of placement group (%s) user_tags: %s", *placementGroup.ID, err)
		}
		d.Set(Attr_UserTags, userTags)
	}
	d.Set(Attr_Members, placementGroup.Members)
	d.Set(Attr_Name, placementGroup.Name)
	d.Set(Attr_Policy, placementGroup.Policy)

	return nil
}
