// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/scc-go-sdk/posturemanagementv1"
)

func dataSourceIBMSccPostureScopes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMSccPostureScopesRead,

		Schema: map[string]*schema.Schema{
			"scope_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An auto-generated unique identifier for the scope.",
			},
			"scopes": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Scopes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A detailed description of the scope.",
						},
						"created_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user who created the scope.",
						},
						"modified_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user who most recently modified the scope.",
						},
						"scope_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An auto-generated unique identifier for the scope.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A unique name for your scope.",
						},
						"enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether scope is enabled/disabled.",
						},
						"environment_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The environment that the scope is targeted to.",
						},
						"created_time": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time that the scope was created in UTC.",
						},
						"modified_time": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time that the scope was last modified in UTC.",
						},
						"last_scan_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last type of scan that was run on the scope.",
						},
						"last_scan_type_description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A description of the last scan type.",
						},
						"last_scan_status_updated_time": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last time that a scan status for a scope was updated in UTC.",
						},
						"collectors_id": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The unique IDs of the collectors that are attached to the scope.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"scans": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A list of the scans that have been run on the scope.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"scan_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "An auto-generated unique identifier for the scan.",
									},
									"discover_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "An auto-generated unique identifier for discovery.",
									},
									"status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the collector as it completes a scan.",
									},
									"status_message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The current status of the collector.",
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

func dataSourceIBMSccPostureScopesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	postureManagementClient, err := meta.(ClientSession).PostureManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	listScopesOptions := &posturemanagementv1.ListScopesOptions{}

	scopesList, response, err := postureManagementClient.ListScopesWithContext(context, listScopesOptions)
	if err != nil {
		log.Printf("[DEBUG] ListScopesWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ListScopesWithContext failed %s\n%s", err, response))
	}

	// Use the provided filter argument and construct a new list with only the requested resource(s)
	var matchScopes []posturemanagementv1.ScopeItem
	var scopeID string
	var suppliedFilter bool

	if v, ok := d.GetOk("scope_id"); ok {
		scopeID = v.(string)
		suppliedFilter = true
		for _, data := range scopesList.Scopes {
			if *data.ScopeID == scopeID {
				matchScopes = append(matchScopes, data)
			}
		}
	} else {
		matchScopes = scopesList.Scopes
	}
	scopesList.Scopes = matchScopes

	if suppliedFilter {
		if len(scopesList.Scopes) == 0 {
			return diag.FromErr(fmt.Errorf("no Scopes found with scopeID %s", scopeID))
		}
		d.SetId(scopeID)
	} else {
		d.SetId(dataSourceIBMSccPostureScopesID(d))
	}

	if scopesList.Scopes != nil {
		err = d.Set("scopes", dataSourceScopesListFlattenScopes(scopesList.Scopes))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting scopes %s", err))
		}
	}

	return nil
}

// dataSourceIBMSccPostureScopesID returns a reasonable ID for the list.
func dataSourceIBMSccPostureScopesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceScopesListFlattenScopes(result []posturemanagementv1.ScopeItem) (scopes []map[string]interface{}) {
	for _, scopesItem := range result {
		scopes = append(scopes, dataSourceScopesListScopesToMap(scopesItem))
	}

	return scopes
}

func dataSourceScopesListScopesToMap(scopesItem posturemanagementv1.ScopeItem) (scopesMap map[string]interface{}) {
	scopesMap = map[string]interface{}{}

	if scopesItem.Description != nil {
		scopesMap["description"] = scopesItem.Description
	}
	if scopesItem.CreatedBy != nil {
		scopesMap["created_by"] = scopesItem.CreatedBy
	}
	if scopesItem.ModifiedBy != nil {
		scopesMap["modified_by"] = scopesItem.ModifiedBy
	}
	if scopesItem.ScopeID != nil {
		scopesMap["scope_id"] = scopesItem.ScopeID
	}
	if scopesItem.Name != nil {
		scopesMap["name"] = scopesItem.Name
	}
	if scopesItem.Enabled != nil {
		scopesMap["enabled"] = scopesItem.Enabled
	}
	if scopesItem.EnvironmentType != nil {
		scopesMap["environment_type"] = scopesItem.EnvironmentType
	}
	if scopesItem.CreatedTime != nil {
		scopesMap["created_time"] = scopesItem.CreatedTime.String()
	}
	if scopesItem.ModifiedTime != nil {
		scopesMap["modified_time"] = scopesItem.ModifiedTime.String()
	}
	if scopesItem.LastScanType != nil {
		scopesMap["last_scan_type"] = scopesItem.LastScanType
	}
	if scopesItem.LastScanTypeDescription != nil {
		scopesMap["last_scan_type_description"] = scopesItem.LastScanTypeDescription
	}
	if scopesItem.LastScanStatusUpdatedTime != nil {
		scopesMap["last_scan_status_updated_time"] = scopesItem.LastScanStatusUpdatedTime.String()
	}
	if scopesItem.CollectorsID != nil {
		scopesMap["collectors_id"] = scopesItem.CollectorsID
	}
	if scopesItem.Scans != nil {
		scansList := []map[string]interface{}{}
		for _, scansItem := range scopesItem.Scans {
			scansList = append(scansList, dataSourceScopesListScopesScansToMap(scansItem))
		}
		scopesMap["scans"] = scansList
	}

	return scopesMap
}

func dataSourceScopesListScopesScansToMap(scansItem posturemanagementv1.Scan) (scansMap map[string]interface{}) {
	scansMap = map[string]interface{}{}

	if scansItem.ScanID != nil {
		scansMap["scan_id"] = scansItem.ScanID
	}
	if scansItem.DiscoverID != nil {
		scansMap["discover_id"] = scansItem.DiscoverID
	}
	if scansItem.Status != nil {
		scansMap["status"] = scansItem.Status
	}
	if scansItem.StatusMessage != nil {
		scansMap["status_message"] = scansItem.StatusMessage
	}

	return scansMap
}
