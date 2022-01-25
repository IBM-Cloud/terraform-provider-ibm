// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func DataSourceIBMCmOffering() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMCmOfferingRead,

		Schema: map[string]*schema.Schema{
			"catalog_identifier": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Catalog identifier.",
				ForceNew:    true,
			},
			"offering_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the catalog containing this offering.",
				ForceNew:    true,
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The url for this specific offering.",
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn for this specific offering.",
			},
			"label": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Display Name in the requested language.",
				ForceNew:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The programmatic name of this offering.",
			},
			"offering_icon_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL for an icon associated with this offering.",
			},
			"offering_docs_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL for an additional docs with this offering.",
			},
			"offering_support_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to be displayed in the Consumption UI for getting support on this offering.",
			},
			"short_description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Short description in the requested language.",
			},
			"long_description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Long description in the requested language.",
			},
			"permit_request_ibm_public_publish": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Is it permitted to request publishing to IBM or Public.",
			},
			"ibm_publish_approved": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if this offering has been approved for use by all IBMers.",
			},
			"public_publish_approved": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if this offering has been approved for use by all IBM Cloud users.",
			},
			"public_original_crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The original offering CRN that this publish entry came from.",
			},
			"publish_public_crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the public catalog entry of this offering.",
			},
			"portal_approval_record": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The portal's approval record ID.",
			},
			"portal_ui_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The portal UI URL.",
			},
			"catalog_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the catalog containing this offering.",
			},
			"catalog_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the catalog.",
			},
			"disclaimer": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A disclaimer for this offering.",
			},
			"hidden": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Determine if this offering should be displayed in the Consumption UI.",
			},
			"repo_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Repository info for offerings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"token": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Token for private repos.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Public or enterprise GitHub.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMCmOfferingRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	catalogManagementClient, err := meta.(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{}

	getOfferingOptions.SetCatalogIdentifier(d.Get("catalog_identifier").(string))
	getOfferingOptions.SetOfferingID(d.Get("offering_id").(string))

	offering, response, err := catalogManagementClient.GetOfferingWithContext(context, getOfferingOptions)
	if err != nil {
		log.Printf("[DEBUG] GetOfferingWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(*offering.ID)
	if err = d.Set("url", offering.URL); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting url: %s", err))
	}
	if err = d.Set("crn", offering.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting crn: %s", err))
	}
	if err = d.Set("label", offering.Label); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting label: %s", err))
	}
	if err = d.Set("name", offering.Name); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting name: %s", err))
	}
	if err = d.Set("offering_icon_url", offering.OfferingIconURL); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting offering_icon_url: %s", err))
	}
	if err = d.Set("offering_docs_url", offering.OfferingDocsURL); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting offering_docs_url: %s", err))
	}
	if err = d.Set("offering_support_url", offering.OfferingSupportURL); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting offering_support_url: %s", err))
	}
	if err = d.Set("short_description", offering.ShortDescription); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting short_description: %s", err))
	}
	if err = d.Set("long_description", offering.LongDescription); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting long_description: %s", err))
	}
	if err = d.Set("permit_request_ibm_public_publish", offering.PermitRequestIBMPublicPublish); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting permit_request_ibm_public_publish: %s", err))
	}
	if err = d.Set("ibm_publish_approved", offering.IBMPublishApproved); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting ibm_publish_approved: %s", err))
	}
	if err = d.Set("public_publish_approved", offering.PublicPublishApproved); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting public_publish_approved: %s", err))
	}
	if err = d.Set("public_original_crn", offering.PublicOriginalCRN); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting public_original_crn: %s", err))
	}
	if err = d.Set("publish_public_crn", offering.PublishPublicCRN); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting publish_public_crn: %s", err))
	}
	if err = d.Set("portal_approval_record", offering.PortalApprovalRecord); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting portal_approval_record: %s", err))
	}
	if err = d.Set("portal_ui_url", offering.PortalUIURL); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting portal_ui_url: %s", err))
	}
	if err = d.Set("catalog_id", offering.CatalogID); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting catalog_id: %s", err))
	}
	if err = d.Set("catalog_name", offering.CatalogName); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting catalog_name: %s", err))
	}
	if err = d.Set("disclaimer", offering.Disclaimer); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting disclaimer: %s", err))
	}
	if err = d.Set("hidden", offering.Hidden); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting hidden: %s", err))
	}

	if offering.RepoInfo != nil {
		repoInfoMap := dataSourceOfferingRepoInfoToMap(*offering.RepoInfo)
		if err = d.Set("repo_info", []map[string]interface{}{repoInfoMap}); err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error setting repo_info %s", err))
		}
	}

	return nil
}

func dataSourceOfferingRepoInfoToMap(repoInfoItem catalogmanagementv1.RepoInfo) (repoInfoMap map[string]interface{}) {
	repoInfoMap = map[string]interface{}{}

	if repoInfoItem.Token != nil {
		repoInfoMap["token"] = repoInfoItem.Token
	}
	if repoInfoItem.Type != nil {
		repoInfoMap["type"] = repoInfoItem.Type
	}

	return repoInfoMap
}
