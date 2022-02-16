// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	// Arguments
	PIPlacementGroupName   = "pi_placement_group_name"
	PIPlacementGroupPolicy = "pi_placement_group_policy"

	// Attributes
	PlacementGroups       = "placement_groups"
	PlacementGroupID      = "placement_group_id"
	PlacementGroupName    = "name"
	PlacementGroupMembers = "members"
	PlacementGroupPolicy  = "policy"

	// Attributes need to fix
	PlacementGroupsID = "id"
)

func DataSourceIBMPIPlacementGroup() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPIPlacementGroupRead,
		Schema: map[string]*schema.Schema{
			PIPlacementGroupName: {
				Type:     schema.TypeString,
				Required: true,
			},

			PlacementGroupPolicy: {
				Type:     schema.TypeString,
				Computed: true,
			},

			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			PlacementGroupMembers: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPIPlacementGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	placementGroupName := d.Get(PIPlacementGroupName).(string)
	client := st.NewIBMPIPlacementGroupClient(ctx, sess, cloudInstanceID)

	response, err := client.Get(placementGroupName)
	if err != nil {
		log.Printf("[DEBUG]  err %s", err)
		return diag.FromErr(err)
	}

	d.SetId(*response.ID)
	d.Set(PlacementGroupPolicy, response.Policy)
	d.Set(PlacementGroupMembers, response.Members)

	return nil
}
