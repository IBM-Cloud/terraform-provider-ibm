// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway

import (
	"github.com/IBM/networking-go-sdk/transitgatewayapisv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
)

func DataSourceIBMTransitGatewayRedundancyGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMTransitGatewayRedundancyGroupRead,
		Schema: map[string]*schema.Schema{
			tgName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the redundancy group",
				ValidateFunc: validate.InvokeValidator("ibm_tg_gateway", tgName),
			},
			tgID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the redundancy group",
			},
			tgCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the redundancy group was created",
			},
			tgUpdatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the redundancy group was last updated",
			},
		},
	}
}

func dataSourceIBMTransitGatewayRedundancyGroupRead(d *schema.ResourceData, meta interface{}) error {
	client, err := transitgatewayClient(meta)
	if err != nil {
		return err
	}

	name := d.Get(tgName).(string)

	listOptions := &transitgatewayapisv1.ListRedundancyGroupsOptions{}
	result, response, err := client.ListRedundancyGroups(listOptions)
	if err != nil {
		return flex.FmtErrorf("[ERROR] Error listing transit gateway redundancy groups %s\n%s", err, response)
	}

	for _, rg := range result.RedundancyGroups {
		if rg.Name != nil && *rg.Name == name {
			d.SetId(*rg.ID)
			d.Set(tgID, *rg.ID)
			d.Set(tgName, *rg.Name)
			if rg.CreatedAt != nil {
				d.Set(tgCreatedAt, rg.CreatedAt.String())
			}
			if rg.UpdatedAt != nil {
				d.Set(tgUpdatedAt, rg.UpdatedAt.String())
			}
			return nil
		}
	}

	return flex.FmtErrorf("[ERROR] Couldn't find any redundancy group with the specified name: (%s)", name)
}
