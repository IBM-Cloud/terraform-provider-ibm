// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	CISCustomListItems          = "ibm_cis_custom_list_items"
	CISCustomListItemID         = "item_id"
	CISCustomListItemsOutput    = "items"
	CISCustomListItemIp         = "ip"
	CISCustomListItemHostname   = "hostname"
	CISCustomListItemASN        = "asn"
	CISCustomListItemComment    = "comment"
	CISCustomListItemCreatedOn  = "created_on"
	CISCustomListItemModifiedOn = "modified_on"
)

func DataSourceIBMCISCustomListItems() *schema.Resource {
	return &schema.Resource{
		Read: DataSourceIBMCISCustomListItemsRead,
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
				Description: "Custom List ID",
				Required:    true,
			},
			CISCustomListItemID: {
				Type:        schema.TypeString,
				Description: "Custom List Item ID",
				Optional:    true,
			},
			CISCustomListItemsOutput: {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Container for response information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						CISCustomListItemID: {
							Type:        schema.TypeString,
							Description: "Custom List Item ID",
							Computed:    true,
						},
						CISCustomListItemIp: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP address of the item",
						},
						CISCustomListItemHostname: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Hostname of the item",
						},
						CISCustomListItemASN: {
							Type:        schema.TypeInt,
							Description: "ASN of the item",
							Computed:    true,
						},
						CISCustomListItemComment: {
							Type:        schema.TypeString,
							Description: "Item comment",
							Computed:    true,
						},
						CISCustomListItemCreatedOn: {
							Type:        schema.TypeString,
							Description: "Item Create date",
							Computed:    true,
						},
						CISCustomListItemModifiedOn: {
							Type:        schema.TypeString,
							Description: "Item last modified date",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func DataSourceIBMCISCustomListItemsValidator() *validate.ResourceValidator {

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
		ResourceName: CISCustomListItems,
		Schema:       validateSchema}
	return &IBMCISCustomListsValidator
}

func DataSourceIBMCISCustomListItemsRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisListsSession()
	if err != nil {
		return err
	}
	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	listId := d.Get(CISCustomListID).(string)
	sess.ListID = core.StringPtr(listId)

	itemId := d.Get(CISCustomListItemID).(string)

	listItemList := make([]map[string]interface{}, 0)
	if itemId != "" {
		sess.ItemID = core.StringPtr(itemId)
		opt := sess.NewGetListItemOptions()
		result, resp, err := sess.GetListItem(opt)

		if err != nil {
			flex.FmtErrorf("[WARN] Get Custom List Item failed: %v\n", resp)
			return err
		}

		itemObj := result.Result
		itemOutput := map[string]interface{}{}
		itemOutput[CISCustomListItemID] = itemObj.ID
		itemOutput[CISCustomListItemIp] = itemObj.Ip
		itemOutput[CISCustomListItemHostname] = itemObj.Hostname
		itemOutput[CISCustomListItemASN] = itemObj.Asn
		itemOutput[CISCustomListItemComment] = itemObj.Comment
		itemOutput[CISCustomListItemCreatedOn] = itemObj.CreatedOn
		itemOutput[CISCustomListItemModifiedOn] = itemObj.ModifiedOn
		listItemList = append(listItemList, itemOutput)

	} else {
		opt := sess.NewGetListItemsOptions()
		result, resp, err := sess.GetListItems(opt)

		if err != nil {
			flex.FmtErrorf("[WARN] List Custom Lists failed: %v\n", resp)
			return err
		}

		for _, itemObj := range result.Result {
			itemOutput := map[string]interface{}{}
			itemOutput[CISCustomListItemID] = itemObj.ID
			itemOutput[CISCustomListItemIp] = itemObj.Ip
			itemOutput[CISCustomListItemHostname] = itemObj.Hostname
			itemOutput[CISCustomListItemASN] = itemObj.Asn
			itemOutput[CISCustomListItemComment] = itemObj.Comment
			itemOutput[CISCustomListItemCreatedOn] = itemObj.CreatedOn
			itemOutput[CISCustomListItemModifiedOn] = itemObj.ModifiedOn
			listItemList = append(listItemList, itemOutput)
		}

	}

	d.Set(CISCustomListID, listId)
	d.Set(CISCustomListItemID, itemId)
	d.SetId(dataSourceCISCustomListItemsCheckID(d))
	d.Set(CISCustomListItemsOutput, listItemList)
	d.Set(cisID, crn)

	return nil
}

func dataSourceCISCustomListItemsCheckID(d *schema.ResourceData) string {
	return "custom_list_item" + ":" + d.Get(CISCustomListID).(string) + ":" + d.Get(cisID).(string)
}
