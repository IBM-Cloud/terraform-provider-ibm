// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func DataSourceIBMCmOfferingInstance() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMCmOfferingInstanceRead,

		Schema: map[string]*schema.Schema{
			"instance_identifier": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Version Instance identifier.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "provisioned instance ID (part of the CRN).",
			},
			"rev": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloudant revision.",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "url reference to this object.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "platform CRN for this instance.",
			},
			"label": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "the label for this instance.",
			},
			"catalog_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Catalog ID this instance was created from.",
			},
			"offering_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Offering ID this instance was created from.",
			},
			"kind_format": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "the format this instance has (helm, operator, ova...).",
			},
			"version": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version this instance was installed from (semver - not version id).",
			},
			"version_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version id this instance was installed from (version id - not semver).",
			},
			"cluster_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster ID.",
			},
			"cluster_region": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster region (e.g., us-south).",
			},
			"cluster_namespaces": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of target namespaces to install into.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cluster_all_namespaces": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "designate to install into all namespaces.",
			},
			"schematics_workspace_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Id of the schematics workspace, for offering instances provisioned through schematics.",
			},
			"install_plan": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of install plan (also known as approval strategy) for operator subscriptions. Can be either automatic, which automatically upgrades operators to the latest in a channel, or manual, which requires approval on the cluster.",
			},
			"channel": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Channel to pin the operator subscription to.",
			},
			"created": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "date and time create.",
			},
			"updated": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "date and time updated.",
			},
			"metadata": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Map of metadata values for this offering instance.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"resource_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Id of the resource group to provision the offering instance into.",
			},
			"location": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "String location of OfferingInstance deployment.",
			},
			"disabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if Resource Controller has disabled this instance.",
			},
			"account": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account this instance is owned by.",
			},
			"last_operation": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "the last operation performed and status.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"operation": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "last operation performed.",
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "state after the last operation performed.",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "additional information about the last operation.",
						},
						"transaction_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "transaction id from the last operation.",
						},
						"updated": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time last updated.",
						},
						"code": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Error code from the last operation, if applicable.",
						},
					},
				},
			},
			"kind_target": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The target kind for the installed software version.",
			},
			"sha": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The digest value of the installed software version.",
			},
		},
	}
}

func dataSourceIBMCmOfferingInstanceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	catalogManagementClient, err := meta.(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getOfferingInstanceOptions := &catalogmanagementv1.GetOfferingInstanceOptions{}

	getOfferingInstanceOptions.SetInstanceIdentifier(d.Get("instance_identifier").(string))

	offeringInstance, response, err := catalogManagementClient.GetOfferingInstanceWithContext(context, getOfferingInstanceOptions)
	if err != nil {
		log.Printf("[DEBUG] GetOfferingInstanceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetOfferingInstanceWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s", *getOfferingInstanceOptions.InstanceIdentifier))

	if err = d.Set("id", offeringInstance.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting id: %s", err))
	}

	if err = d.Set("rev", offeringInstance.Rev); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting rev: %s", err))
	}

	if err = d.Set("url", offeringInstance.URL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting url: %s", err))
	}

	if err = d.Set("crn", offeringInstance.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}

	if err = d.Set("label", offeringInstance.Label); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting label: %s", err))
	}

	if err = d.Set("catalog_id", offeringInstance.CatalogID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting catalog_id: %s", err))
	}

	if err = d.Set("offering_id", offeringInstance.OfferingID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting offering_id: %s", err))
	}

	if err = d.Set("kind_format", offeringInstance.KindFormat); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting kind_format: %s", err))
	}

	if err = d.Set("version", offeringInstance.Version); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version: %s", err))
	}

	if err = d.Set("version_id", offeringInstance.VersionID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version_id: %s", err))
	}

	if err = d.Set("cluster_id", offeringInstance.ClusterID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting cluster_id: %s", err))
	}

	if err = d.Set("cluster_region", offeringInstance.ClusterRegion); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting cluster_region: %s", err))
	}

	if err = d.Set("cluster_all_namespaces", offeringInstance.ClusterAllNamespaces); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting cluster_all_namespaces: %s", err))
	}

	if err = d.Set("schematics_workspace_id", offeringInstance.SchematicsWorkspaceID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting schematics_workspace_id: %s", err))
	}

	if err = d.Set("install_plan", offeringInstance.InstallPlan); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting install_plan: %s", err))
	}

	if err = d.Set("channel", offeringInstance.Channel); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting channel: %s", err))
	}

	if err = d.Set("created", flex.DateTimeToString(offeringInstance.Created)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created: %s", err))
	}

	if err = d.Set("updated", flex.DateTimeToString(offeringInstance.Updated)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated: %s", err))
	}

	if offeringInstance.Metadata != nil {
		convertedMap := make(map[string]interface{}, len(offeringInstance.Metadata))
		for k, v := range offeringInstance.Metadata {
			convertedMap[k] = v
		}

		if err = d.Set("metadata", flex.Flatten(convertedMap)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting metadata: %s", err))
		}
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting metadata %s", err))
		}
	}

	if err = d.Set("resource_group_id", offeringInstance.ResourceGroupID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group_id: %s", err))
	}

	if err = d.Set("location", offeringInstance.Location); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting location: %s", err))
	}

	if err = d.Set("disabled", offeringInstance.Disabled); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting disabled: %s", err))
	}

	if err = d.Set("account", offeringInstance.Account); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting account: %s", err))
	}

	lastOperation := []map[string]interface{}{}
	if offeringInstance.LastOperation != nil {
		modelMap, err := dataSourceIBMCmOfferingInstanceOfferingInstanceLastOperationToMap(offeringInstance.LastOperation)
		if err != nil {
			return diag.FromErr(err)
		}
		lastOperation = append(lastOperation, modelMap)
	}
	if err = d.Set("last_operation", lastOperation); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_operation %s", err))
	}

	if err = d.Set("kind_target", offeringInstance.KindTarget); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting kind_target: %s", err))
	}

	if err = d.Set("sha", offeringInstance.Sha); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting sha: %s", err))
	}

	return nil
}

func dataSourceIBMCmOfferingInstanceOfferingInstanceLastOperationToMap(model *catalogmanagementv1.OfferingInstanceLastOperation) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Operation != nil {
		modelMap["operation"] = *model.Operation
	}
	if model.State != nil {
		modelMap["state"] = *model.State
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	if model.TransactionID != nil {
		modelMap["transaction_id"] = *model.TransactionID
	}
	if model.Updated != nil {
		modelMap["updated"] = model.Updated.String()
	}
	if model.Code != nil {
		modelMap["code"] = *model.Code
	}
	return modelMap, nil
}
