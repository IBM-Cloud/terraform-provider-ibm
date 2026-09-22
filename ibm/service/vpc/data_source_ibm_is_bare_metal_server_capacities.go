// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBareMetalServerCapacitiesList = "capacities"
)

func DataSourceIBMIsBareMetalServerCapacities() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISBareMetalServerCapacitiesRead,

		Schema: map[string]*schema.Schema{
			"profile": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of a bare metal profile",
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of a zone",
			},
			isBareMetalServerCapacitiesList: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A page of available bare metal server capacities",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"profile": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The profile available in the zone",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this bare metal server profile",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name for this bare metal server profile",
									},
									"resource_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type",
									},
								},
							},
						},
						"zone": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The zone where one or more bare metal servers of the profile are available",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this zone",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique name for this zone",
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

func dataSourceIBMISBareMetalServerCapacitiesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_capacities", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	profileFilter := ""
	if profileFilterOk, ok := d.GetOk("profile"); ok {
		profileFilter = profileFilterOk.(string)
	}
	zoneFilter := ""
	if zoneFilterOk, ok := d.GetOk("zone"); ok {
		zoneFilter = zoneFilterOk.(string)
	}

	start := ""
	allCapacities := []vpcv1.BareMetalServerCapacity{}

	for {
		options := &vpcv1.ListBareMetalServerCapacitiesOptions{}
		if start != "" {
			options.Start = &start
		}
		if profileFilter != "" {
			options.ProfileName = &profileFilter
		}
		if zoneFilter != "" {
			options.ZoneName = &zoneFilter
		}

		bmCapacities, _, err := sess.ListBareMetalServerCapacitiesWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListBareMetalServerCapacitiesWithContext failed: %s", err.Error()), "(Data) ibm_is_bare_metal_server_capacities", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if bmCapacities == nil {
			tfErr := flex.TerraformErrorf(nil, "ListBareMetalServerCapacitiesWithContext returned nil response", "(Data) ibm_is_bare_metal_server_capacities", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		start = flex.GetNext(bmCapacities.Next)
		allCapacities = append(allCapacities, bmCapacities.Capacities...)

		if start == "" {
			break
		}
	}

	capacitiesInfo := make([]map[string]interface{}, 0)

	for _, capacity := range allCapacities {
		capacityMap := make(map[string]interface{})

		if capacity.Profile != nil {
			profileList := make([]map[string]interface{}, 1)
			profileMap := make(map[string]interface{})

			if capacity.Profile.Href != nil {
				profileMap["href"] = *capacity.Profile.Href
			}
			if capacity.Profile.Name != nil {
				profileMap["name"] = *capacity.Profile.Name
			}
			if capacity.Profile.ResourceType != nil {
				profileMap["resource_type"] = *capacity.Profile.ResourceType
			}

			profileList[0] = profileMap
			capacityMap["profile"] = profileList
		}

		if capacity.Zone != nil {
			zoneList := make([]map[string]interface{}, 1)
			zoneMap := make(map[string]interface{})

			if capacity.Zone.Href != nil {
				zoneMap["href"] = *capacity.Zone.Href
			}
			if capacity.Zone.Name != nil {
				zoneMap["name"] = *capacity.Zone.Name
			}

			zoneList[0] = zoneMap
			capacityMap["zone"] = zoneList
		}

		capacitiesInfo = append(capacitiesInfo, capacityMap)
	}

	d.SetId(dataSourceIBMISBMSCapacitiesID(d))
	if err = d.Set(isBareMetalServerCapacitiesList, capacitiesInfo); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting capacities: %s", err), "(Data) ibm_is_bare_metal_server_capacities", "read", "set-capacities").GetDiag()
	}

	return nil
}

func dataSourceIBMISBMSCapacitiesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
