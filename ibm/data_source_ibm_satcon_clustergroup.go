// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMSatconClusterGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMSatconClusterGroupRead,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Description: "ID of the clustergroup",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name or id of the clustergroup",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"created": {
				Description: "Creation time of the clustergroup",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": {
							Description: "ID of the cluster",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Name of the cluster",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMSatconClusterGroupRead(d *schema.ResourceData, meta interface{}) error {
	satconClient, err := meta.(ClientSession).SatellitConfigClientSession()
	if err != nil {
		return err
	}
	satconGroupAPI := satconClient.Groups

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	var name string

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}

	if name == "" {
		return fmt.Errorf("provided satellite clustergroup name is empty")
	}

	group, err := satconGroupAPI.GroupByName(userDetails.userAccount, name)
	if err != nil {
		return fmt.Errorf("error retrieving satellite clustergroup: %s", err)
	}

	clusters := make([]map[string]interface{}, 0)
	for _, c := range group.Clusters {
		cluster := map[string]interface{}{
			"id":   c.ID,
			"name": c.Name,
		}
		clusters = append(clusters, cluster)
	}

	d.SetId(group.Name)
	d.Set("uuid", group.UUID)
	d.Set("name", group.Name)
	d.Set("created", group.Created)
	d.Set("clusters", clusters)

	return nil
}
