/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibm

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

func dataSourceIbmIsDedicatedHostGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmIsDedicatedHostGroupsRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The unique user-defined name for this dedicated host group. If unspecified, the name will be a hyphenated list of randomly-selected words.",
			},
			"first": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "A link to the first page of resources.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for a page of resources.",
						},
					},
				},
			},
			"groups": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of dedicated host groups.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dedicated host profile class for hosts in this group.",
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that the dedicated host group was created.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this dedicated host group.",
						},
						"dedicated_hosts": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The dedicated hosts that are in this dedicated host group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this dedicated host.",
									},
									"deleted": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"more_info": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Link to documentation about deleted resources.",
												},
											},
										},
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this dedicated host.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this dedicated host.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique user-defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly-selected words.",
									},
									"resource_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of resource referenced.",
									},
								},
							},
						},
						"family": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dedicated host profile family for hosts in this group.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this dedicated host group.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this dedicated host group.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name for this dedicated host group. If unspecified, the name will be a hyphenated list of randomly-selected words.",
						},
						"resource_group": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "The resource group for this dedicated host group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this resource group.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this resource group.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The user-defined name for this resource group.",
									},
								},
							},
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of resource referenced.",
						},
						"supported_instance_profiles": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Array of instance profiles that can be used by instances placed on this dedicated host group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this virtual server instance profile.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique name for this virtual server instance profile.",
									},
								},
							},
						},
						"zone": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "The zone this dedicated host group resides in.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this zone.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique name for this zone.",
									},
								},
							},
						},
					},
				},
			},
			"limit": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum number of resources that can be returned by the request.",
			},
			"next": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "A link to the next page of resources. This property is present for all pagesexcept the last page.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for a page of resources.",
						},
					},
				},
			},
			"total_count": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of resources across all pages.",
			},
		},
	}
}

func dataSourceIbmIsDedicatedHostGroupsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(ClientSession).VpcV1()
	if err != nil {
		return diag.FromErr(err)
	}

	listDedicatedHostGroupsOptions := &vpcv1.ListDedicatedHostGroupsOptions{}

	dedicatedHostGroupCollection, response, err := vpcClient.ListDedicatedHostGroupsWithContext(context, listDedicatedHostGroupsOptions)
	if err != nil {
		log.Printf("[DEBUG] ListDedicatedHostGroupsWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	// Use the provided filter argument and construct a new list with only the requested resource(s)
	var matchGroups []vpcv1.DedicatedHostGroup
	var name string
	var suppliedFilter bool

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
		suppliedFilter = true
		for _, data := range dedicatedHostGroupCollection.Groups {
			if *data.Name == name {
				matchGroups = append(matchGroups, data)
			}
		}
	} else {
		matchGroups = dedicatedHostGroupCollection.Groups
	}
	dedicatedHostGroupCollection.Groups = matchGroups

	if len(dedicatedHostGroupCollection.Groups) == 0 {
		return diag.FromErr(fmt.Errorf("no Groups found with name %s\nIf not specified, please specify more filters", name))
	}

	if suppliedFilter {
		d.SetId(name)
	} else {
		d.SetId(dataSourceIbmIsDedicatedHostGroupsID(d))
	}

	if dedicatedHostGroupCollection.First != nil {
		err = d.Set("first", dataSourceDedicatedHostGroupCollectionFlattenFirst(*dedicatedHostGroupCollection.First))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting first %s", err))
		}
	}

	if dedicatedHostGroupCollection.Groups != nil {
		err = d.Set("groups", dataSourceDedicatedHostGroupCollectionFlattenGroups(dedicatedHostGroupCollection.Groups))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting groups %s", err))
		}
	}
	if err = d.Set("limit", dedicatedHostGroupCollection.Limit); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting limit: %s", err))
	}

	if dedicatedHostGroupCollection.Next != nil {
		err = d.Set("next", dataSourceDedicatedHostGroupCollectionFlattenNext(*dedicatedHostGroupCollection.Next))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting next %s", err))
		}
	}
	if err = d.Set("total_count", dedicatedHostGroupCollection.TotalCount); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting total_count: %s", err))
	}

	return nil
}

// dataSourceIbmIsDedicatedHostGroupsID returns a reasonable ID for the list.
func dataSourceIbmIsDedicatedHostGroupsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceDedicatedHostGroupCollectionFlattenFirst(result vpcv1.DedicatedHostGroupCollectionFirst) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceDedicatedHostGroupCollectionFirstToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceDedicatedHostGroupCollectionFirstToMap(firstItem vpcv1.DedicatedHostGroupCollectionFirst) (firstMap map[string]interface{}) {
	firstMap = map[string]interface{}{}

	if firstItem.Href != nil {
		firstMap["href"] = firstItem.Href
	}

	return firstMap
}


func dataSourceDedicatedHostGroupCollectionFlattenGroups(result []vpcv1.DedicatedHostGroup) (groups []map[string]interface{}) {
	for _, groupsItem := range result {
		groups = append(groups, dataSourceDedicatedHostGroupCollectionGroupsToMap(groupsItem))
	}

	return groups
}

func dataSourceDedicatedHostGroupCollectionGroupsToMap(groupsItem vpcv1.DedicatedHostGroup) (groupsMap map[string]interface{}) {
	groupsMap = map[string]interface{}{}

	if groupsItem.Class != nil {
		groupsMap["class"] = groupsItem.Class
	}
	if groupsItem.CreatedAt != nil {
		groupsMap["created_at"] = groupsItem.CreatedAt.String()
	}
	if groupsItem.Crn != nil {
		groupsMap["crn"] = groupsItem.Crn
	}
	if groupsItem.DedicatedHosts != nil {
		dedicatedHostsList := []map[string]interface{}{}
		for _, dedicatedHostsItem := range groupsItem.DedicatedHosts {
			dedicatedHostsList = append(dedicatedHostsList, dataSourceDedicatedHostGroupCollectionGroupsDedicatedHostsToMap(dedicatedHostsItem))
		}
		groupsMap["dedicated_hosts"] = dedicatedHostsList
	}
	if groupsItem.Family != nil {
		groupsMap["family"] = groupsItem.Family
	}
	if groupsItem.Href != nil {
		groupsMap["href"] = groupsItem.Href
	}
	if groupsItem.ID != nil {
		groupsMap["id"] = groupsItem.ID
	}
	if groupsItem.Name != nil {
		groupsMap["name"] = groupsItem.Name
	}
	if groupsItem.ResourceGroup != nil {
		resourceGroupList := []map[string]interface{}{}
		resourceGroupMap := dataSourceDedicatedHostGroupCollectionGroupsResourceGroupToMap(*groupsItem.ResourceGroup)
		resourceGroupList = append(resourceGroupList, resourceGroupMap)
		groupsMap["resource_group"] = resourceGroupList
	}
	if groupsItem.ResourceType != nil {
		groupsMap["resource_type"] = groupsItem.ResourceType
	}
	if groupsItem.SupportedInstanceProfiles != nil {
		supportedInstanceProfilesList := []map[string]interface{}{}
		for _, supportedInstanceProfilesItem := range groupsItem.SupportedInstanceProfiles {
			supportedInstanceProfilesList = append(supportedInstanceProfilesList, dataSourceDedicatedHostGroupCollectionGroupsSupportedInstanceProfilesToMap(supportedInstanceProfilesItem))
		}
		groupsMap["supported_instance_profiles"] = supportedInstanceProfilesList
	}
	if groupsItem.Zone != nil {
		zoneList := []map[string]interface{}{}
		zoneMap := dataSourceDedicatedHostGroupCollectionGroupsZoneToMap(*groupsItem.Zone)
		zoneList = append(zoneList, zoneMap)
		groupsMap["zone"] = zoneList
	}

	return groupsMap
}

func dataSourceDedicatedHostGroupCollectionGroupsDedicatedHostsToMap(dedicatedHostsItem vpcv1.DedicatedHostReference) (dedicatedHostsMap map[string]interface{}) {
	dedicatedHostsMap = map[string]interface{}{}

	if dedicatedHostsItem.Crn != nil {
		dedicatedHostsMap["crn"] = dedicatedHostsItem.Crn
	}
	if dedicatedHostsItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceDedicatedHostGroupCollectionDedicatedHostsDeletedToMap(*dedicatedHostsItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		dedicatedHostsMap["deleted"] = deletedList
	}
	if dedicatedHostsItem.Href != nil {
		dedicatedHostsMap["href"] = dedicatedHostsItem.Href
	}
	if dedicatedHostsItem.ID != nil {
		dedicatedHostsMap["id"] = dedicatedHostsItem.ID
	}
	if dedicatedHostsItem.Name != nil {
		dedicatedHostsMap["name"] = dedicatedHostsItem.Name
	}
	if dedicatedHostsItem.ResourceType != nil {
		dedicatedHostsMap["resource_type"] = dedicatedHostsItem.ResourceType
	}

	return dedicatedHostsMap
}


func dataSourceDedicatedHostGroupCollectionGroupsResourceGroupToMap(resourceGroupItem vpcv1.ResourceGroupReference) (resourceGroupMap map[string]interface{}) {
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


func dataSourceDedicatedHostGroupCollectionGroupsSupportedInstanceProfilesToMap(supportedInstanceProfilesItem vpcv1.InstanceProfileReference) (supportedInstanceProfilesMap map[string]interface{}) {
	supportedInstanceProfilesMap = map[string]interface{}{}

	if supportedInstanceProfilesItem.Href != nil {
		supportedInstanceProfilesMap["href"] = supportedInstanceProfilesItem.Href
	}
	if supportedInstanceProfilesItem.Name != nil {
		supportedInstanceProfilesMap["name"] = supportedInstanceProfilesItem.Name
	}

	return supportedInstanceProfilesMap
}


func dataSourceDedicatedHostGroupCollectionGroupsZoneToMap(zoneItem vpcv1.ZoneReference) (zoneMap map[string]interface{}) {
	zoneMap = map[string]interface{}{}

	if zoneItem.Href != nil {
		zoneMap["href"] = zoneItem.Href
	}
	if zoneItem.Name != nil {
		zoneMap["name"] = zoneItem.Name
	}

	return zoneMap
}



func dataSourceDedicatedHostGroupCollectionFlattenNext(result vpcv1.DedicatedHostGroupCollectionNext) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceDedicatedHostGroupCollectionNextToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceDedicatedHostGroupCollectionNextToMap(nextItem vpcv1.DedicatedHostGroupCollectionNext) (nextMap map[string]interface{}) {
	nextMap = map[string]interface{}{}

	if nextItem.Href != nil {
		nextMap["href"] = nextItem.Href
	}

	return nextMap
}

