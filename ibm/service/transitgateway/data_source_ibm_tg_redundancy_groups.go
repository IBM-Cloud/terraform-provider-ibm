// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway

import (
	"time"

	"github.com/IBM/networking-go-sdk/transitgatewayapisv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

const (
	tgRedundancyGroups  = "redundancy_groups"
	tgRedundancyGroupAt = "created_at"
)

func DataSourceIBMTransitGatewayRedundancyGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMTransitGatewayRedundancyGroupsRead,
		Schema: map[string]*schema.Schema{
			tgName: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter redundancy groups by name",
			},
			tgRedundancyGroups: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of redundancy groups in the account",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						tgID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier of the redundancy group",
						},
						tgName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the redundancy group",
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
				},
			},
		},
	}
}

func dataSourceIBMTransitGatewayRedundancyGroupsRead(d *schema.ResourceData, meta interface{}) error {
	client, err := transitgatewayClient(meta)
	if err != nil {
		return err
	}

	listOptions := &transitgatewayapisv1.ListRedundancyGroupsOptions{}
	if name, ok := d.GetOk(tgName); ok {
		nameStr := name.(string)
		listOptions.Name = &nameStr
	}
	result, response, err := client.ListRedundancyGroups(listOptions)
	if err != nil {
		return flex.FmtErrorf("[ERROR] Error listing transit gateway redundancy groups %s\n%s", err, response)
	}

	groups := make([]map[string]interface{}, 0, len(result.RedundancyGroups))
	for _, rg := range result.RedundancyGroups {
		m := map[string]interface{}{}
		if rg.ID != nil {
			m[tgID] = *rg.ID
		}
		if rg.Name != nil {
			m[tgName] = *rg.Name
		}
		if rg.CreatedAt != nil {
			m[tgCreatedAt] = rg.CreatedAt.String()
		}
		if rg.UpdatedAt != nil {
			m[tgUpdatedAt] = rg.UpdatedAt.String()
		}
		groups = append(groups, m)
	}

	d.Set(tgRedundancyGroups, groups)
	d.SetId(time.Now().UTC().String())
	return nil
}
