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

func dataSourceIbmResourceBindings() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmResourceBindingsRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The human-readable name of the binding.",
			},
			"guid": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Short ID of the binding.",
			},
			"resource_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Short ID of resource group.",
			},
			"resource_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The unique ID of the offering (service name). This value is provided by and stored in the global catalog.",
			},
			"region_binding_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Short ID of the instance in a specific targeted environment.",
			},
			"bindings": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of resource bindings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID associated with the binding.",
						},
						"guid": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "When you create a new binding, a globally unique identifier (GUID) is assigned. This GUID is a unique internal identifier managed by the resource controller that corresponds to the binding.",
						},
						"url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "When you provision a new binding, a relative URL path is created identifying the location of the binding.",
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date when the binding was created.",
						},
						"updated_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date when the binding was last updated.",
						},
						"deleted_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date when the binding was deleted.",
						},
						"created_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subject who created the binding.",
						},
						"updated_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subject who updated the binding.",
						},
						"deleted_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subject who deleted the binding.",
						},
						"source_crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN of resource alias associated to the binding.",
						},
						"target_crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN of target resource, for example, application, in a specific environment.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The full Cloud Resource Name (CRN) associated with the binding. For more information about this format, see [Cloud Resource Names](https://cloud.ibm.com/docs/overview?topic=overview-crn).",
						},
						"region_binding_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the binding in the specific target environment, for example, `service_binding_id` in a given IBM Cloud environment.",
						},
						"region_binding_crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN of the binding in the specific target environment.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The human-readable name of the binding.",
						},
						"account_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An alpha-numeric value identifying the account ID.",
						},
						"resource_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the resource group.",
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state of the binding.",
						},
						"credentials": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The credentials for the binding. Additional key-value pairs are passed through from the resource brokers.  For additional details, see the service’s documentation.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"apikey": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The API key for the credentials.",
										Sensitive:   true,
									},
									"iam_apikey_description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The optional description of the API key.",
									},
									"iam_apikey_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the API key.",
									},
									"iam_role_crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Cloud Resource Name for the role of the credentials.",
									},
									"iam_serviceid_crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Cloud Resource Name for the service ID of the credentials.",
									},
								},
							},
						},
						"iam_compatible": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies whether the binding’s credentials support IAM.",
						},
						"resource_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique ID of the offering. This value is provided by and stored in the global catalog.",
						},
						"migrated": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "A boolean that dictates if the alias was migrated from a previous CF instance.",
						},
						"resource_alias_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The relative path to the resource alias that this binding is associated with.",
						},
					},
				},
			},
		},
	}
}

func dataSourceResourceBindingsGetNext(next *string) (string, error) {
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

func dataSourceIbmResourceBindingsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceControllerClient, err := meta.(ClientSession).ResourceControllerV2API()
	if err != nil {
		return diag.FromErr(err)
	}

	listResourceBindingsOptions := &resourcecontrollerv2.ListResourceBindingsOptions{}
	if v, ok := d.GetOk("name"); ok && v != nil {
		listResourceBindingsOptions.SetName(v.(string))
	}
	if v, ok := d.GetOk("guid"); ok && v != nil {
		listResourceBindingsOptions.SetGUID(v.(string))
	}
	if v, ok := d.GetOk("resource_group_id"); ok && v != nil {
		listResourceBindingsOptions.SetResourceGroupID(v.(string))
	}
	if v, ok := d.GetOk("resource_id"); ok && v != nil {
		listResourceBindingsOptions.SetResourceID(v.(string))
	}
	if v, ok := d.GetOk("region_binding_id"); ok && v != nil {
		listResourceBindingsOptions.SetRegionBindingID(v.(string))
	}

	start := ""
	var allRecords []resourcecontrollerv2.ResourceBinding
	for {
		if start != "" {
			listResourceBindingsOptions.Start = &start
		}
		resourceBindingsListResponse, response, err := resourceControllerClient.ListResourceBindingsWithContext(context, listResourceBindingsOptions)
		if err != nil {
			log.Printf("[DEBUG] ListResourceBindingsWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("ListResourceBindingsWithContext failed %s\n%s", err, response))
		}
		start, err = dataSourceResourceBindingsGetNext(resourceBindingsListResponse.NextURL)
		if err != nil {
			log.Printf("[DEBUG] ListResourceBindingsWithContext failed. Error occurred while parsing NextURL: %s", err)
			return diag.FromErr(err)
		}
		allRecords = append(allRecords, resourceBindingsListResponse.Resources...)
		if start == "" {
			break
		}
	}

	if allRecords != nil {
		err = d.Set("bindings", dataSourceResourceBindingsListFlattenResources(allRecords))
		if err != nil {
			log.Printf("[DEBUG] Error setting resource bindings %s", err)
			return diag.FromErr(fmt.Errorf("Error setting resource bindings %s", err))
		}
	}

	d.SetId(dataSourceIbmResourceBindingsID(d))
	return nil
}

// dataSourceIbmResourceBindingsID returns a reasonable ID for the list.
func dataSourceIbmResourceBindingsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceResourceBindingsListFlattenResources(result []resourcecontrollerv2.ResourceBinding) (resources []map[string]interface{}) {
	for _, resourcesItem := range result {
		resources = append(resources, dataSourceResourceBindingsListResourcesToMap(resourcesItem))
	}

	return resources
}

func dataSourceResourceBindingsListResourcesToMap(resourcesItem resourcecontrollerv2.ResourceBinding) (resourcesMap map[string]interface{}) {
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
	if resourcesItem.SourceCRN != nil {
		resourcesMap["source_crn"] = resourcesItem.SourceCRN
	}
	if resourcesItem.TargetCRN != nil {
		resourcesMap["target_crn"] = resourcesItem.TargetCRN
	}
	if resourcesItem.CRN != nil {
		resourcesMap["crn"] = resourcesItem.CRN
	}
	if resourcesItem.RegionBindingID != nil {
		resourcesMap["region_binding_id"] = resourcesItem.RegionBindingID
	}
	if resourcesItem.RegionBindingCRN != nil {
		resourcesMap["region_binding_crn"] = resourcesItem.RegionBindingCRN
	}
	if resourcesItem.Name != nil {
		resourcesMap["name"] = resourcesItem.Name
	}
	if resourcesItem.AccountID != nil {
		resourcesMap["account_id"] = resourcesItem.AccountID
	}
	if resourcesItem.ResourceGroupID != nil {
		resourcesMap["resource_group_id"] = resourcesItem.ResourceGroupID
	}
	if resourcesItem.State != nil {
		resourcesMap["state"] = resourcesItem.State
	}
	if resourcesItem.Credentials != nil {
		credentialsList := []map[string]interface{}{}
		credentialsMap := dataSourceResourceBindingsListResourcesCredentialsToMap(*resourcesItem.Credentials)
		credentialsList = append(credentialsList, credentialsMap)
		resourcesMap["credentials"] = credentialsList
	}
	if resourcesItem.IamCompatible != nil {
		resourcesMap["iam_compatible"] = resourcesItem.IamCompatible
	}
	if resourcesItem.ResourceID != nil {
		resourcesMap["resource_id"] = resourcesItem.ResourceID
	}
	if resourcesItem.Migrated != nil {
		resourcesMap["migrated"] = resourcesItem.Migrated
	}
	if resourcesItem.ResourceAliasURL != nil {
		resourcesMap["resource_alias_url"] = resourcesItem.ResourceAliasURL
	}

	return resourcesMap
}

func dataSourceResourceBindingsListResourcesCredentialsToMap(credentialsItem resourcecontrollerv2.Credentials) (credentialsMap map[string]interface{}) {
	credentialsMap = map[string]interface{}{}

	if credentialsItem.Apikey != nil {
		credentialsMap["apikey"] = credentialsItem.Apikey
	}
	if credentialsItem.IamApikeyDescription != nil {
		credentialsMap["iam_apikey_description"] = credentialsItem.IamApikeyDescription
	}
	if credentialsItem.IamApikeyName != nil {
		credentialsMap["iam_apikey_name"] = credentialsItem.IamApikeyName
	}
	if credentialsItem.IamRoleCRN != nil {
		credentialsMap["iam_role_crn"] = credentialsItem.IamRoleCRN
	}
	if credentialsItem.IamServiceidCRN != nil {
		credentialsMap["iam_serviceid_crn"] = credentialsItem.IamServiceidCRN
	}

	return credentialsMap
}
