// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/listsapiv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const ()

func ResourceIBMCISCustomListItems() *schema.Resource {
	return &schema.Resource{
		Create: ResourceIBMCISCustomListItemsCreate,
		Update: ResourceIBMCISCustomListItemsUpdate,
		Delete: ResourceIBMCISCustomListItemsDelete,
		Read:   ResourceIBMCISCustomListItemsRead,
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
			CISCustomListItemsOutput: {
				Type:        schema.TypeSet,
				Required:    true,
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
							Optional:    true,
							Description: "IP address of the item",
						},
						CISCustomListItemHostname: {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Hostname of the item",
						},
						CISCustomListItemASN: {
							Type:        schema.TypeInt,
							Description: "ASN of the item",
							Optional:    true,
						},
						CISCustomListItemComment: {
							Type:        schema.TypeString,
							Description: "Item comment",
							Optional:    true,
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

func ResourceIBMCISCustomListItemsValidator() *validate.ResourceValidator {

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

func ResourceIBMCISCustomListItemsCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisListsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while creating the CisListsSession %s", err)
	}

	crn := d.Get(cisID).(string)
	listId := d.Get(CISCustomListID).(string)
	sess.Crn = core.StringPtr(crn)
	sess.ListID = core.StringPtr(listId)

	opt := sess.NewCreateListItemsOptions()

	itemsList := d.Get(CISCustomListItemsOutput)
	itemsListRes := itemsList.(*schema.Set).List()

	itemsReqObj := make([]listsapiv1.CreateListItemsReqItem, 0)

	for _, val := range itemsListRes {

		itemResObj := listsapiv1.CreateListItemsReqItem{}
		itemObj := val.(map[string]interface{})

		ip := itemObj[CISCustomListItemIp].(string)
		hostname := itemObj[CISCustomListItemHostname].(string)
		asn := itemObj[CISCustomListItemASN].(int)
		comment := itemObj[CISCustomListItemComment].(string)

		if ip != "" {
			itemResObj.Ip = &ip
		}
		if hostname != "" {
			itemResObj.Hostname = &hostname
		}
		if asn != 0 {
			fasn := float64(asn)
			itemResObj.Asn = &fasn
		}
		if comment != "" {
			itemResObj.Comment = &comment
		}
		itemsReqObj = append(itemsReqObj, itemResObj)
	}

	opt.SetCreateListItemsReqItem(itemsReqObj)
	result, resp, err := sess.CreateListItems(opt)

	if err != nil || result == nil {
		return fmt.Errorf("[ERROR] Error creating  custom List items : %s %s", err, resp)
	}
	d.SetId(flex.ConvertCisToTfTwoVar(listId, crn))

	return ResourceIBMCISCustomListItemsRead(d, meta)
}

func ResourceIBMCISCustomListItemsUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisListsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while creating the CisListsSession %s", err)
	}

	if d.HasChange(CISCustomListItemsOutput) {

		listId, crn, _ := flex.ConvertTftoCisTwoVar(d.Id())
		sess.Crn = &crn
		sess.ListID = core.StringPtr(listId)

		opt := sess.NewUpdateListItemsOptions()
		itemsList := d.Get(CISCustomListItemsOutput)
		itemsListRes := itemsList.(*schema.Set).List()

		itemsReqObj := make([]listsapiv1.CreateListItemsReqItem, 0)

		for _, val := range itemsListRes {

			itemResObj := listsapiv1.CreateListItemsReqItem{}
			itemObj := val.(map[string]interface{})
			ip := itemObj[CISCustomListItemIp].(string)
			hostname := itemObj[CISCustomListItemHostname].(string)
			asn := itemObj[CISCustomListItemASN].(int)
			comment := itemObj[CISCustomListItemComment].(string)

			if ip != "" {
				itemResObj.Ip = &ip
			}
			if hostname != "" {
				itemResObj.Hostname = &hostname
			}
			if asn != 0 {
				fasn := float64(asn)
				itemResObj.Asn = &fasn
			}
			if comment != "" {
				itemResObj.Comment = &comment
			}
			itemsReqObj = append(itemsReqObj, itemResObj)
		}

		opt.SetCreateListItemsReqItem(itemsReqObj)
		result, resp, err := sess.UpdateListItems(opt)

		if err != nil || result == nil {
			return fmt.Errorf("[ERROR] Error creating  custom List items : %s %s", err, resp)
		}
		d.SetId(flex.ConvertCisToTfTwoVar(listId, crn))
	}
	return ResourceIBMCISCustomListItemsRead(d, meta)
}

func ResourceIBMCISCustomListItemsRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisListsSession()
	if err != nil {
		return err
	}

	listId, crn, _ := flex.ConvertTftoCisTwoVar(d.Id())

	sess.Crn = &crn
	sess.ListID = core.StringPtr(listId)

	listItemList := make([]map[string]interface{}, 0)

	opt := sess.NewGetListItemsOptions()
	result, resp, err := sess.GetListItems(opt)

	if err != nil {
		log.Printf("[WARN] List Custom Lists failed: %v\n", resp)
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

	d.Set(CISCustomListID, listId)
	d.Set(CISCustomListItemsOutput, listItemList)
	d.Set(cisID, crn)

	return nil
}

func ResourceIBMCISCustomListItemsDelete(d *schema.ResourceData, meta interface{}) error {

	return nil
}
