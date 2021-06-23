// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

func dataSourceIbmResourceAliases() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmResourceAliasesRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The human-readable name of the alias.",
			},
			"guid": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Short ID of the alias.",
			},
			"resource_instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Resource instance CRN.",
			},
			"region_instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Short ID of the instance in a specific targeted environment.",
			},
			"resource_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The unique ID of the offering (service name). This value is provided by and stored in the global catalog.",
			},
			"resource_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Short ID of Resource group.",
			},
			"aliases": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of resource aliases.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID associated with the alias.",
						},
						"guid": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "When you create a new alias, a globally unique identifier (GUID) is assigned. This GUID is a unique internal indentifier managed by the resource controller that corresponds to the alias.",
						},
						"url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "When you created a new alias, a relative URL path is created identifying the location of the alias.",
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date when the alias was created.",
						},
						"updated_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date when the alias was last updated.",
						},
						"deleted_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date when the alias was deleted.",
						},
						"created_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subject who created the alias.",
						},
						"updated_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subject who updated the alias.",
						},
						"deleted_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subject who deleted the alias.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The human-readable name of the alias.",
						},
						"resource_instance_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the resource instance that is being aliased.",
						},
						"target_crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN of the target namespace in the specific environment.",
						},
						"account_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An alpha-numeric value identifying the account ID.",
						},
						"resource_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique ID of the offering. This value is provided by and stored in the global catalog.",
						},
						"resource_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the resource group.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN of the alias. For more information about this format, see [Cloud Resource Names](https://cloud.ibm.com/docs/overview?topic=overview-crn).",
						},
						"region_instance_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the instance in the specific target environment, for example, `service_instance_id` in a given IBM Cloud environment.",
						},
						"region_instance_crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN of the instance in the specific target environment.",
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state of the alias.",
						},
						"migrated": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "A boolean that dictates if the alias was migrated from a previous CF instance.",
						},
						"resource_instance_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The relative path to the resource instance.",
						},
						"resource_bindings_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The relative path to the resource bindings for the alias.",
						},
						"resource_keys_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The relative path to the resource keys for the alias.",
						},
					},
				},
			},
		},
	}
}

func dataSourceResourceAliasesGetNext(next *string) (string, error) {
	if reflect.ValueOf(next).IsNil() {
		return "", nil
	}
	u, err := url.Parse(*next)
	if err != nil {
		return "", err
	}
	q := u.Query()
	return q.Get("start"), nil
}

func dataSourceIbmResourceAliasesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceControllerClient, err := meta.(ClientSession).ResourceControllerV2API()
	if err != nil {
		return diag.FromErr(err)
	}

	listResourceAliasesOptions := &resourcecontrollerv2.ListResourceAliasesOptions{}
	if v, ok := d.GetOk("name"); ok && v != nil {
		listResourceAliasesOptions.SetName(v.(string))
	}
	if v, ok := d.GetOk("guid"); ok && v != nil {
		listResourceAliasesOptions.SetGUID(v.(string))
	}
	if v, ok := d.GetOk("resource_instance_id"); ok && v != nil {
		listResourceAliasesOptions.SetResourceInstanceID(v.(string))
	}
	if v, ok := d.GetOk("region_instance_id"); ok && v != nil {
		listResourceAliasesOptions.SetRegionInstanceID(v.(string))
	}
	if v, ok := d.GetOk("resource_id"); ok && v != nil {
		listResourceAliasesOptions.SetResourceID(v.(string))
	}
	if v, ok := d.GetOk("resource_group_id"); ok && v != nil {
		listResourceAliasesOptions.SetResourceGroupID(v.(string))
	}

	start := ""
	var allRecords []resourcecontrollerv2.ResourceAlias
	for {
		if start != "" {
			listResourceAliasesOptions.Start = &start
		}
		resourceAliasesResponse, response, err := resourceControllerClient.ListResourceAliasesWithContext(context, listResourceAliasesOptions)
		if err != nil {
			log.Printf("[DEBUG] ListResourceAliasesWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("ListResourceAliasesWithContext failed %s\n%s", err, response))
		}
		start, err = dataSourceResourceAliasesGetNext(resourceAliasesResponse.NextURL)
		if err != nil {
			log.Printf("[DEBUG] ListResourceAliasesWithContext failed. Error occurred while parsing NextURL: %s", err)
			return diag.FromErr(err)
		}
		allRecords = append(allRecords, resourceAliasesResponse.Resources...)
		if start == "" {
			break
		}
	}

	if allRecords != nil {
		err = d.Set("aliases", dataSourceResourceAliasesListFlattenResources(allRecords))
		if err != nil {
			log.Printf("[DEBUG] Error setting resource aliases %s", err)
			return diag.FromErr(fmt.Errorf("Error setting resource aliases %s", err))
		}
	}

	d.SetId(dataSourceIbmResourceAliasesID(d))
	return nil
}

// dataSourceIbmResourceAliasesID returns a reasonable ID for the list.
func dataSourceIbmResourceAliasesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceResourceAliasesListFlattenResources(result []resourcecontrollerv2.ResourceAlias) (resources []map[string]interface{}) {
	for _, resourcesItem := range result {
		resources = append(resources, dataSourceResourceAliasesListResourcesToMap(resourcesItem))
	}

	return resources
}

func dataSourceResourceAliasesListResourcesToMap(resourcesItem resourcecontrollerv2.ResourceAlias) (resourcesMap map[string]interface{}) {
	resourcesMap = map[string]interface{}{}

	if resourcesItem.ID != nil {
		resourcesMap["id"] = resourcesItem.ID
	}
	if resourcesItem.GUID != nil {
		resourcesMap["guid"] = resourcesItem.GUID
	}
	if resourcesItem.URL != nil {
		resourcesMap["url"] = resourcesItem.URL
	}
	if resourcesItem.CreatedAt != nil {
		resourcesMap["created_at"] = resourcesItem.CreatedAt.String()
	}
	if resourcesItem.UpdatedAt != nil {
		resourcesMap["updated_at"] = resourcesItem.UpdatedAt.String()
	}
	if resourcesItem.DeletedAt != nil {
		resourcesMap["deleted_at"] = resourcesItem.DeletedAt.String()
	}
	if resourcesItem.CreatedBy != nil {
		resourcesMap["created_by"] = resourcesItem.CreatedBy
	}
	if resourcesItem.UpdatedBy != nil {
		resourcesMap["updated_by"] = resourcesItem.UpdatedBy
	}
	if resourcesItem.DeletedBy != nil {
		resourcesMap["deleted_by"] = resourcesItem.DeletedBy
	}
	if resourcesItem.Name != nil {
		resourcesMap["name"] = resourcesItem.Name
	}
	if resourcesItem.ResourceInstanceID != nil {
		resourcesMap["resource_instance_id"] = resourcesItem.ResourceInstanceID
	}
	if resourcesItem.TargetCRN != nil {
		resourcesMap["target_crn"] = resourcesItem.TargetCRN
	}
	if resourcesItem.AccountID != nil {
		resourcesMap["account_id"] = resourcesItem.AccountID
	}
	if resourcesItem.ResourceID != nil {
		resourcesMap["resource_id"] = resourcesItem.ResourceID
	}
	if resourcesItem.ResourceGroupID != nil {
		resourcesMap["resource_group_id"] = resourcesItem.ResourceGroupID
	}
	if resourcesItem.CRN != nil {
		resourcesMap["crn"] = resourcesItem.CRN
	}
	if resourcesItem.RegionInstanceID != nil {
		resourcesMap["region_instance_id"] = resourcesItem.RegionInstanceID
	}
	if resourcesItem.RegionInstanceCRN != nil {
		resourcesMap["region_instance_crn"] = resourcesItem.RegionInstanceCRN
	}
	if resourcesItem.State != nil {
		resourcesMap["state"] = resourcesItem.State
	}
	if resourcesItem.Migrated != nil {
		resourcesMap["migrated"] = resourcesItem.Migrated
	}
	if resourcesItem.ResourceInstanceURL != nil {
		resourcesMap["resource_instance_url"] = resourcesItem.ResourceInstanceURL
	}
	if resourcesItem.ResourceBindingsURL != nil {
		resourcesMap["resource_bindings_url"] = resourcesItem.ResourceBindingsURL
	}
	if resourcesItem.ResourceKeysURL != nil {
		resourcesMap["resource_keys_url"] = resourcesItem.ResourceKeysURL
	}

	return resourcesMap
}
