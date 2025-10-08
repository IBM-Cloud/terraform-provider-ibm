// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcemanager

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	rg "github.com/IBM/platform-services-go-sdk/resourcemanagerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMResourceGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMResourceGroupsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Resource group name",
				Type:        schema.TypeString,
				Optional:    true,
				ValidateFunc: validate.InvokeDataSourceValidator("ibm_resource_groups",
					"name"),
			},
			"is_default": {
				Description: "Default Resource group",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"include_deleted": {
				Description: "Include deleted resource groups",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"date": {
				Description: "The date in the format of YYYY-MM which returns resource groups. Deleted resource groups will be excluded before this month.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"resource_groups": {
				Type:        schema.TypeList,
				Description: "List of resource groups",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "The ID of the resource group",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Resource group name",
							Computed:    true,
						},
						"is_default": {
							Type:        schema.TypeBool,
							Description: "Default Resource group",
							Computed:    true,
						},
						"state": {
							Type:        schema.TypeString,
							Description: "State of the resource group",
							Computed:    true,
						},
						"crn": {
							Type:        schema.TypeString,
							Description: "The full CRN associated with the resource group",
							Computed:    true,
						},
						"created_at": {
							Type:        schema.TypeString,
							Description: "The date when the resource group was initially created.",
							Computed:    true,
						},
						"updated_at": {
							Type:        schema.TypeString,
							Description: "The date when the resource group was last updated.",
							Computed:    true,
						},
						"teams_url": {
							Type:        schema.TypeString,
							Description: "The URL to access the team details that associated with the resource group.",
							Computed:    true,
						},
						"payment_methods_url": {
							Type:        schema.TypeString,
							Description: "The URL to access the payment methods details that associated with the resource group.",
							Computed:    true,
						},
						"quota_url": {
							Type:        schema.TypeString,
							Description: "The URL to access the quota details that associated with the resource group.",
							Computed:    true,
						},
						"quota_id": {
							Type:        schema.TypeString,
							Description: "An alpha-numeric value identifying the quota ID associated with the resource group.",
							Computed:    true,
						},
						"resource_linkages": {
							Type:        schema.TypeList,
							Description: "An array of the resources that linked to the resource group",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Computed:    true,
						},
						"account_id": {
							Type:        schema.TypeString,
							Description: "Account ID",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func DataSourceIBMResourceGroupsValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_group",
			CloudDataRange:             []string{"resolved_to:name"},
			Optional:                   true})

	ibmIBMResourceGroupsValidator := validate.ResourceValidator{ResourceName: "ibm_resource_groups", Schema: validateSchema}
	return &ibmIBMResourceGroupsValidator
}

func dataSourceIBMResourceGroupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	rMgtClient, err := meta.(conns.ClientSession).ResourceManagerV2API()
	if err != nil {
		return diag.FromErr(err)
	}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(err)
	}
	accountID := userDetails.UserAccount

	// Build the list options with available filters
	resourceGroupListOptions := rg.ListResourceGroupsOptions{
		AccountID: &accountID,
	}

	// Apply filters
	if name, ok := d.GetOk("name"); ok {
		nameStr := name.(string)
		resourceGroupListOptions.Name = &nameStr
	}

	if isDefault, ok := d.GetOk("is_default"); ok {
		defaultBool := isDefault.(bool)
		resourceGroupListOptions.Default = &defaultBool
	}

	if includeDeleted, ok := d.GetOk("include_deleted"); ok {
		includeDeletedBool := includeDeleted.(bool)
		resourceGroupListOptions.IncludeDeleted = &includeDeletedBool
	}

	if date, ok := d.GetOk("date"); ok {
		dateStr := date.(string)
		resourceGroupListOptions.Date = &dateStr
	}

	// List resource groups
	rgList, response, err := rMgtClient.ListResourceGroupsWithContext(ctx, &resourceGroupListOptions)
	if err != nil {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Error retrieving resource groups",
				Detail:   fmt.Sprintf("Error: %s, Response: %v", err, response),
			},
		}
	}

	if rgList == nil || rgList.Resources == nil {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "No resource groups found",
				Detail:   "No resource groups were returned from the API",
			},
		}
	}

	// Convert resource groups to list of maps
	resourceGroups := make([]map[string]interface{}, 0)
	for _, resourceGroup := range rgList.Resources {
		rgMap := map[string]interface{}{}

		if resourceGroup.ID != nil {
			rgMap["id"] = *resourceGroup.ID
		}
		if resourceGroup.Name != nil {
			rgMap["name"] = *resourceGroup.Name
		}
		if resourceGroup.Default != nil {
			rgMap["is_default"] = *resourceGroup.Default
		}
		if resourceGroup.State != nil {
			rgMap["state"] = *resourceGroup.State
		}
		if resourceGroup.CRN != nil {
			rgMap["crn"] = *resourceGroup.CRN
		}
		if resourceGroup.CreatedAt != nil {
			rgMap["created_at"] = resourceGroup.CreatedAt.String()
		}
		if resourceGroup.UpdatedAt != nil {
			rgMap["updated_at"] = resourceGroup.UpdatedAt.String()
		}
		if resourceGroup.TeamsURL != nil {
			rgMap["teams_url"] = *resourceGroup.TeamsURL
		}
		if resourceGroup.PaymentMethodsURL != nil {
			rgMap["payment_methods_url"] = *resourceGroup.PaymentMethodsURL
		}
		if resourceGroup.QuotaURL != nil {
			rgMap["quota_url"] = *resourceGroup.QuotaURL
		}
		if resourceGroup.QuotaID != nil {
			rgMap["quota_id"] = *resourceGroup.QuotaID
		}
		if resourceGroup.AccountID != nil {
			rgMap["account_id"] = *resourceGroup.AccountID
		}
		if resourceGroup.ResourceLinkages != nil {
			rl := make([]string, 0)
			for _, r := range resourceGroup.ResourceLinkages {
				rl = append(rl, r.(string))
			}
			rgMap["resource_linkages"] = rl
		}

		resourceGroups = append(resourceGroups, rgMap)
	}

	// Set the data source ID and resource groups
	d.SetId(accountID)
	if err := d.Set("resource_groups", resourceGroups); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
