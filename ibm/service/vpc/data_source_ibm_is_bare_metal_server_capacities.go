// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBareMetalServerCapacities = "capacities"
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
			isBareMetalServerCapacities: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of capacities for each profile. The results will include all profile capacities unless a zone or profile filter are specified.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the profile",
						},
						"zones": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of zones in the region that have capacity for the profile",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISBareMetalServerCapacitiesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	profile := d.Get("profile").(string)
	zone := d.Get("zone").(string)

	sess, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_capacities", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	start := ""
	allCapacities := []vpcv1.BareMetalServerCapacity{}

	for {
		options := &vpcv1.ListBareMetalServerCapacitiesOptions{}
		if profile != "" {
			options.ProfileName = &profile
		}
		if zone != "" {
			options.ZoneName = &zone
		}
		if start != "" {
			options.Start = &start
		}

		bmCapacities, _, err := sess.ListBareMetalServerCapacitiesWithContext(context, options)
		if err != nil || bmCapacities == nil || bmCapacities.Capacities == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListBareMetalServerCapacitiesWithContext failed: %s", err.Error()), "(Data) ibm_is_bare_metal_server_capacities", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		start = flex.GetNext(bmCapacities.Next)
		allCapacities = append(allCapacities, bmCapacities.Capacities...)

		if start == "" {
			break
		}
	}

	capacities := allCapacities

	profileZones := make(map[string]map[string]bool)

	for _, cap := range capacities {
		var profileName, zoneName string
		if cap.Profile != nil && cap.Profile.Name != nil {
			profileName = *cap.Profile.Name
		}
		if cap.Zone != nil && cap.Zone.Name != nil {
			zoneName = *cap.Zone.Name
		}

		if profileName == "" || zoneName == "" {
			continue
		}

		if _, exists := profileZones[profileName]; !exists {
			profileZones[profileName] = make(map[string]bool)
		}
		profileZones[profileName][zoneName] = true
	}

	capacitiesInfo := make([]map[string]interface{}, 0, len(profileZones))

	profileNames := make([]string, 0, len(profileZones))
	for pName := range profileZones {
		profileNames = append(profileNames, pName)
	}
	sort.Strings(profileNames)

	for _, pName := range profileNames {
		zones := profileZones[pName]
		zoneList := make([]string, 0, len(zones))
		for z := range zones {
			zoneList = append(zoneList, z)
		}
		sort.Strings(zoneList)

		capacitiesInfo = append(capacitiesInfo, map[string]interface{}{
			"name":  pName,
			"zones": zoneList,
		})
	}

	d.SetId(dataSourceIBMISBMSCapacitiesID(d))
	if err = d.Set(isBareMetalServerCapacities, capacitiesInfo); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting capacities: %s", err), "(Data) ibm_is_bare_metal_server_capacities", "read", "set-capacities").GetDiag()
	}

	return nil
}

func dataSourceIBMISBMSCapacitiesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
