// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	//"fmt"

	"context"
	"log"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func DataSourceIBMPIDisasterRecoveryLocations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIDisasterRecoveryLocations,
		Schema: map[string]*schema.Schema{

			// Computed Attributes
			"disaster_recovery_locations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"location": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "RegionZone of a site",
						},
						"replication_sites": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_active": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"location": {
										Type:     schema.TypeString,
										Computed: true,
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

func dataSourceIBMPIDisasterRecoveryLocations(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	drClient := instance.NewIBMPIDisasterRecoveryLocationClient(ctx, sess, "")
	drLocationSites, err := drClient.GetAll()
	if err != nil {
		return diag.FromErr(err)
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	d.Set("disaster_recovery_locations", flattenDisasterRecoveryLocations(drLocationSites.DisasterRecoveryLocations))

	return nil
}

func flattenDisasterRecoveryLocations(list []*models.DisasterRecoveryLocation) []map[string]interface{} {
	log.Printf("Calling the flattenDisasterRecoveryLocations call with list %d", len(list))
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		l := map[string]interface{}{
			"location":          i.Location,
			"replication_sites": flattenDisasterRecoveryLocation(i.ReplicationSites),
		}

		result = append(result, l)
	}

	return result
}
