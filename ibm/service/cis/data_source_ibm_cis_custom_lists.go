// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	CISCustomLists           = "ibm_cis_custom_lists"
	CISCustomListsOutput     = "lists"
	CISCustomListName        = "name"
	CISCustomListDescription = "description"
	CISCustomListKind        = "kind"
	CISCustomListID          = "list_id"
	CISCustomListItemNumbers = "num_items"
	CISCustomListFilters     = "num_referencing_filters"
)

func DataSourceIBMCISCustomLists() *schema.Resource {
	return &schema.Resource{
		Read: resourceIBMCISCustomListsRead,
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
				ValidateFunc: validate.InvokeDataSourceValidator(
					"ibm_cis_ruleset_versions",
					"cis_id"),
			},
			CISCustomListID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Optional:    true,
				ValidateFunc: validate.InvokeDataSourceValidator(
					"ibm_cis_custom_lists",
					"cis_id"),
			},
			CISCustomListsOutput: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Container for response information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						CISCustomListID: {
							Type:        schema.TypeString,
							Description: "CIS instance crn",
							Optional:    true,
							ValidateFunc: validate.InvokeDataSourceValidator(
								"ibm_cis_custom_lists",
								"cis_id"),
						},
						CISCustomListName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the List",
						},
						CISCustomListKind: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Kind of the List",
						},
						CISCustomListDescription: {
							Type:        schema.TypeString,
							Description: "Description of the List",
							Computed:    true,
						},
						CISCustomListItemNumbers: {
							Type:        schema.TypeInt,
							Description: "Number of items in the List",
							Computed:    true,
						},
						CISCustomListFilters: {
							Type:        schema.TypeInt,
							Description: "Number of times the list is used by rule expressions.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func DataSourceIBMCISCustomListsValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true})

	IBMCISCustomListsValidator := validate.ResourceValidator{
		ResourceName: CISCustomLists,
		Schema:       validateSchema}
	return &IBMCISCustomListsValidator
}

func resourceIBMCISCustomListsRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisListsSession()
	if err != nil {
		return err
	}
	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	listId := d.Get(CISCustomListID).(string)
	listsList := make([]map[string]interface{}, 0)
	if listId != "" {
		sess.ListID = core.StringPtr(listId)
		opt := sess.NewGetCustomListOptions()
		result, resp, err := sess.GetCustomList(opt)

		if err != nil {
			log.Printf("[WARN] List Custom List failed: %v\n", resp)
			return err
		}

		listObj := result.Result
		listOutput := map[string]interface{}{}
		listOutput[CISCustomListID] = listObj.ID
		listOutput[CISCustomListDescription] = listObj.Description
		listOutput[CISCustomListKind] = listObj.Kind
		listOutput[CISCustomListName] = listObj.Name
		listOutput[CISCustomListItemNumbers] = listObj.NumItems
		listOutput[CISCustomListFilters] = listObj.NumReferencingFilters
		listsList = append(listsList, listOutput)

	} else {
		opt := sess.NewGetCustomListsOptions()
		result, resp, err := sess.GetCustomLists(opt)

		if err != nil {
			log.Printf("[WARN] List Custom Lists failed: %v\n", resp)
			return err
		}

		for _, listObj := range result.Result {
			listOutput := map[string]interface{}{}
			listOutput[CISCustomListID] = listObj.ID
			listOutput[CISCustomListDescription] = listObj.Description
			listOutput[CISCustomListKind] = listObj.Kind
			listOutput[CISCustomListName] = listObj.Name
			listOutput[CISCustomListItemNumbers] = listObj.NumItems
			listOutput[CISCustomListFilters] = listObj.NumReferencingFilters
			listsList = append(listsList, listOutput)
		}

	}

	d.Set(CISCustomListID, listId)
	d.SetId(resourceSourceCISListsCustomCheckID(d))
	d.Set(CISCustomListsOutput, listsList)
	d.Set(cisID, crn)

	return nil
}

func resourceSourceCISListsCustomCheckID(d *schema.ResourceData) string {
	return "custom_lists" + ":" + d.Get(cisID).(string)
}
