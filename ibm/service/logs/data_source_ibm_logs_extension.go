// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func DataSourceIbmLogsExtension() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmLogsExtensionRead,

		Schema: map[string]*schema.Schema{
			"logs_extension_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of the extension.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the Extension.",
			},
			"revisions": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of all revisions of the Extension, each representing a versioned snapshot of the Extension's functionality and appearance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version identifier for this revision of the Extension.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The detailed description of what this revision includes, changes made, and any important information users should be aware of.",
						},
						"excerpt": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The brief summary or excerpt of the Extension's description for quick reference.",
						},
						"labels": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of labels or tags associated with the Extension for front-end categorization and filtering.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"items": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The Extension items included in this revision.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the Extension item.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the Extension item.",
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The detailed description of the Extension item.",
									},
									"target_domain": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The domain of the Extension item.",
									},
									"is_mandatory": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "A flag to indicate if the Extension item is mandatory or not. Mandatory items must be specified when deploying the Extension.",
									},
								},
							},
						},
					},
				},
			},
			"keywords": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of keywords to enhance search capabilities on the front-end side.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"changelog": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The of changelog entries made in each version of the Extension.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of the Extension this changelog entry refers to.",
						},
						"description_md": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the changes made in this version, formatted in Markdown for rich text presentation.",
						},
					},
				},
			},
			"deployment": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Deployment details of an Extension scoped by extension ID in the path.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of the Extension revision to deploy.",
						},
						"item_ids": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of Extension item IDs to deploy.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"applications": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Applications that the Extension is deployed for. When this is empty, it is applied to all applications.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"subsystems": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Subsystems that the Extension is deployed. When this is empty, it is applied to all subsystems.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier of the extension.",
						},
					},
				},
			},
			"deprecation": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Deprecation details of the Extension.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reason": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reason why the element (e.g., an Extension or a version of it) is being deprecated.",
						},
						"replacement_extensions": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of Extension IDs that serve as replacements for the deprecated element.",
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

func dataSourceIbmLogsExtensionRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_extension", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getLogsInstanceRegion(logsClient, d)
	instanceId := d.Get("instance_id").(string)
	logsClient, err = getClientWithLogsInstanceEndpoint(logsClient, meta, instanceId, region, getLogsInstanceEndpointType(logsClient, d))
	if err != nil {
		return diag.FromErr(fmt.Errorf("Unable to get updated logs instance client"))
	}

	getExtensionOptions := &logsv0.GetExtensionOptions{}

	getExtensionOptions.SetID(d.Get("logs_extension_id").(string))

	extension, response, err := logsClient.GetExtensionWithContext(context, getExtensionOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetExtensionWithContext failed: %s, %s", err.Error(), response), "(Data) ibm_logs_extension", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*extension.ID)

	if err = d.Set("name", extension.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_logs_extension", "read", "set-name").GetDiag()
	}

	if !core.IsNil(extension.Revisions) {
		revisions := []map[string]interface{}{}
		for _, revisionsItem := range extension.Revisions {
			revisionsItemMap, err := DataSourceIbmLogsExtensionExtensionsV1ExtensionRevisionToMap(&revisionsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_extension", "read", "revisions-to-map").GetDiag()
			}
			revisions = append(revisions, revisionsItemMap)
		}
		if err = d.Set("revisions", revisions); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting revisions: %s", err), "(Data) ibm_logs_extension", "read", "set-revisions").GetDiag()
		}
	}

	if !core.IsNil(extension.Keywords) {
		keywords := []interface{}{}
		for _, keywordsItem := range extension.Keywords {
			keywords = append(keywords, keywordsItem)
		}
		if err = d.Set("keywords", keywords); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting keywords: %s", err), "(Data) ibm_logs_extension", "read", "set-keywords").GetDiag()
		}
	}

	if !core.IsNil(extension.Changelog) {
		changelog := []map[string]interface{}{}
		for _, changelogItem := range extension.Changelog {
			changelogItemMap, err := DataSourceIbmLogsExtensionExtensionsV1ChangelogEntryToMap(&changelogItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_extension", "read", "changelog-to-map").GetDiag()
			}
			changelog = append(changelog, changelogItemMap)
		}
		if err = d.Set("changelog", changelog); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting changelog: %s", err), "(Data) ibm_logs_extension", "read", "set-changelog").GetDiag()
		}
	}

	if !core.IsNil(extension.Deployment) {
		deployment := []map[string]interface{}{}
		deploymentMap, err := DataSourceIbmLogsExtensionExtensionDeploymentToMap(extension.Deployment)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_extension", "read", "deployment-to-map").GetDiag()
		}
		deployment = append(deployment, deploymentMap)
		if err = d.Set("deployment", deployment); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting deployment: %s", err), "(Data) ibm_logs_extension", "read", "set-deployment").GetDiag()
		}
	}

	if !core.IsNil(extension.Deprecation) {
		deprecation := []map[string]interface{}{}
		deprecationMap, err := DataSourceIbmLogsExtensionExtensionsV1DeprecationToMap(extension.Deprecation)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_extension", "read", "deprecation-to-map").GetDiag()
		}
		deprecation = append(deprecation, deprecationMap)
		if err = d.Set("deprecation", deprecation); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting deprecation: %s", err), "(Data) ibm_logs_extension", "read", "set-deprecation").GetDiag()
		}
	}

	return nil
}

func DataSourceIbmLogsExtensionExtensionsV1ExtensionRevisionToMap(model *logsv0.ExtensionsV1ExtensionRevision) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["version"] = *model.Version
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Excerpt != nil {
		modelMap["excerpt"] = *model.Excerpt
	}
	if model.Labels != nil {
		modelMap["labels"] = model.Labels
	}
	if model.Items != nil {
		items := []map[string]interface{}{}
		for _, itemsItem := range model.Items {
			itemsItemMap, err := DataSourceIbmLogsExtensionExtensionsV1ExtensionItemToMap(&itemsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			items = append(items, itemsItemMap)
		}
		modelMap["items"] = items
	}
	return modelMap, nil
}

func DataSourceIbmLogsExtensionExtensionsV1ExtensionItemToMap(model *logsv0.ExtensionsV1ExtensionItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	modelMap["target_domain"] = *model.TargetDomain
	if model.IsMandatory != nil {
		modelMap["is_mandatory"] = *model.IsMandatory
	}
	return modelMap, nil
}

func DataSourceIbmLogsExtensionExtensionsV1ChangelogEntryToMap(model *logsv0.ExtensionsV1ChangelogEntry) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["version"] = *model.Version
	modelMap["description_md"] = *model.DescriptionMd
	return modelMap, nil
}

func DataSourceIbmLogsExtensionExtensionDeploymentToMap(model *logsv0.ExtensionDeployment) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["version"] = *model.Version
	modelMap["item_ids"] = model.ItemIds
	if model.Applications != nil {
		modelMap["applications"] = model.Applications
	}
	if model.Subsystems != nil {
		modelMap["subsystems"] = model.Subsystems
	}
	// Sub-resource: no ID field on model
	return modelMap, nil
}

func DataSourceIbmLogsExtensionExtensionsV1DeprecationToMap(model *logsv0.ExtensionsV1Deprecation) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["reason"] = *model.Reason
	if model.ReplacementExtensions != nil {
		modelMap["replacement_extensions"] = model.ReplacementExtensions
	}
	return modelMap, nil
}
