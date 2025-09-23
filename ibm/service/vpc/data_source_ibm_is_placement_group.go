// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIbmIsPlacementGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmIsPlacementGroupRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The placement group name.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the placement group was created.",
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this placement group.",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this placement group.",
			},
			"lifecycle_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the placement group.",
			},
			"resource_group": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The resource group for this placement group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this resource group.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this resource group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this resource group.",
						},
					},
				},
			},
			"resource_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"strategy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The strategy for this placement group- `host_spread`: place on different compute hosts- `power_spread`: place on compute hosts that use different power sourcesThe enumerated values for this property may expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the placement group on which the unexpected strategy was encountered.",
			},
			isPlacementGroupTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of tags",
			},
			isPlacementGroupAccessTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
			},
		},
	}
}

func dataSourceIbmIsPlacementGroupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_placement_group", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	pgname := d.Get("name").(string)
	listPlacementGroupsOptions := &vpcv1.ListPlacementGroupsOptions{}
	start := ""
	allrecs := []vpcv1.PlacementGroup{}
	for {
		if start != "" {
			listPlacementGroupsOptions.Start = &start
		}
		placementGroupCollection, _, err := vpcClient.ListPlacementGroupsWithContext(context, listPlacementGroupsOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListPlacementGroupsWithContext failed: %s", err.Error()), "(Data) ibm_is_placement_group", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		start = flex.GetNext(placementGroupCollection.Next)
		allrecs = append(allrecs, placementGroupCollection.PlacementGroups...)
		if start == "" {
			break
		}
	}
	for _, placementGroup := range allrecs {
		if *placementGroup.Name == pgname {

			d.SetId(*placementGroup.ID)
			if err = d.Set("created_at", flex.DateTimeToString(placementGroup.CreatedAt)); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_placement_group", "read", "set-created_at").GetDiag()
			}
			if err = d.Set("crn", placementGroup.CRN); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_placement_group", "read", "set-crn").GetDiag()
			}
			if err = d.Set("href", placementGroup.Href); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_placement_group", "read", "set-href").GetDiag()
			}
			if err = d.Set("lifecycle_state", placementGroup.LifecycleState); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "(Data) ibm_is_placement_group", "read", "set-lifecycle_state").GetDiag()
			}
			if err = d.Set("name", placementGroup.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_placement_group", "read", "set-name").GetDiag()
			}

			if placementGroup.ResourceGroup != nil {
				err = d.Set("resource_group", dataSourcePlacementGroupFlattenResourceGroup(*placementGroup.ResourceGroup))
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_placement_group", "read", "set-resource_group").GetDiag()
				}
			}
			if err = d.Set("resource_type", placementGroup.ResourceType); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_placement_group", "read", "set-resource_type").GetDiag()
			}
			if err = d.Set("strategy", placementGroup.Strategy); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting strategy: %s", err), "(Data) ibm_is_placement_group", "read", "set-strategy").GetDiag()
			}
			tags, err := flex.GetGlobalTagsUsingCRN(meta, *placementGroup.CRN, "", isUserTagType)
			if err != nil {
				log.Printf(
					"An error getting placement group (%s) tags : %s", d.Id(), err)
			}

			accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *placementGroup.CRN, "", isAccessTagType)
			if err != nil {
				log.Printf(
					"Error getting placement group (%s) access tags: %s", d.Id(), err)
			}

			d.Set(isPlacementGroupTags, tags)
			if err = d.Set("tags", tags); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tags: %s", err), "(Data) ibm_is_placement_group", "read", "set-tags").GetDiag()
			}
			d.Set(isPlacementGroupAccessTags, accesstags)
			if err = d.Set("access_tags", accesstags); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_tags: %s", err), "(Data) ibm_is_placement_group", "read", "set-access_tags").GetDiag()
			}
			return nil
		}
	}
	return diag.FromErr(fmt.Errorf("[ERROR] No placement group found with name %s", pgname))
}

func dataSourcePlacementGroupFlattenResourceGroup(result vpcv1.ResourceGroupReference) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourcePlacementGroupResourceGroupToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourcePlacementGroupResourceGroupToMap(resourceGroupItem vpcv1.ResourceGroupReference) (resourceGroupMap map[string]interface{}) {
	resourceGroupMap = map[string]interface{}{}

	if resourceGroupItem.Href != nil {
		resourceGroupMap["href"] = resourceGroupItem.Href
	}
	if resourceGroupItem.ID != nil {
		resourceGroupMap["id"] = resourceGroupItem.ID
	}
	if resourceGroupItem.Name != nil {
		resourceGroupMap["name"] = resourceGroupItem.Name
	}

	return resourceGroupMap
}
