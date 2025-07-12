// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	CISCustomList     = "ibm_cis_custom_list"
	AllowedKindValues = "ip, asn, hostname"
)

func ResourceIBMCISCustomList() *schema.Resource {
	return &schema.Resource{
		Read:     ResourceIBMCISCustomListRead,
		Create:   ResourceIBMCISCustomListCreate,
		Update:   ResourceIBMCISCustomListUpdate,
		Delete:   ResourceIBMCustomListDelete,
		Importer: &schema.ResourceImporter{},
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
				Computed:    true,
			},
			CISCustomListName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the List",
			},
			CISCustomListKind: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Kind of the List",
				ValidateFunc: validate.InvokeValidator(CISCustomList, CISCustomListKind),
			},
			CISCustomListDescription: {
				Type:        schema.TypeString,
				Description: "Description of the List",
				Optional:    true,
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
	}
}

func ResourceIBMCISCustomListValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 CISCustomListKind,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			AllowedValues:              AllowedKindValues,
			Required:                   true})

	IBMCISCustomListsValidator := validate.ResourceValidator{
		ResourceName: CISCustomList,
		Schema:       validateSchema}
	return &IBMCISCustomListsValidator
}

func ResourceIBMCISCustomListCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisListsSession()
	if err != nil {
		return flex.FmtErrorf("[ERROR] Error while creating the CisListsSession %s", err)
	}

	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)
	opt := sess.NewCreateCustomListsOptions()

	if n, ok := d.GetOk(CISCustomListName); ok {
		name := n.(string)
		opt.Name = &name
	}
	if k, ok := d.GetOk(CISCustomListKind); ok {
		kind := k.(string)
		opt.Kind = &kind
	}
	if des, ok := d.GetOk(CISCustomListDescription); ok {
		description := des.(string)
		opt.Description = &description
	}

	result, resp, err := sess.CreateCustomLists(opt)

	if err != nil || result == nil {
		return flex.FmtErrorf("[ERROR] Error creating  custom List : %s %s", err, resp)
	}
	d.SetId(flex.ConvertCisToTfTwoVar(*result.Result.ID, crn))

	return ResourceIBMCISCustomListRead(d, meta)
}

func ResourceIBMCISCustomListUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisListsSession()
	if err != nil {
		return flex.FmtErrorf("[ERROR] Error while creating the CisListsSession %s", err)
	}

	listId, crn, _ := flex.ConvertTftoCisTwoVar(d.Id())

	sess.Crn = &crn
	sess.ListID = core.StringPtr(listId)
	if d.HasChange(CISCustomListName) {
		return flex.FmtErrorf("List's name can not be changed")
	}
	if d.HasChange(CISCustomListKind) {
		return flex.FmtErrorf("List's kind can not be changed")
	}
	if d.HasChange(CISCustomListDescription) {
		opt := sess.NewUpdateCustomListOptions()
		if des, ok := d.GetOk(CISCustomListDescription); ok {
			description := des.(string)
			opt.Description = &description
		}
		result, resp, err := sess.UpdateCustomList(opt)
		if err != nil || result == nil {
			return flex.FmtErrorf("[ERROR] Error updating  custom List : %s %s", err, resp)
		}
	}

	return ResourceIBMCISCustomListRead(d, meta)
}

func ResourceIBMCISCustomListRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisListsSession()
	if err != nil {
		return err
	}

	listId, crn, _ := flex.ConvertTftoCisTwoVar(d.Id())

	sess.Crn = &crn
	sess.ListID = core.StringPtr(listId)

	opt := sess.NewGetCustomListOptions()
	result, _, err := sess.GetCustomList(opt)

	if err != nil {
		flex.FmtErrorf("[ERROR] Get Custom List failed: %v\n", err)
		return err
	}

	d.Set(CISCustomListID, listId)
	d.Set(cisID, crn)
	d.Set(CISCustomListName, result.Result.Name)
	d.Set(CISCustomListKind, result.Result.Kind)
	d.Set(CISCustomListDescription, result.Result.Description)
	d.Set(CISCustomListFilters, result.Result.NumReferencingFilters)
	d.Set(CISCustomListItemNumbers, result.Result.NumItems)

	return nil
}

func ResourceIBMCustomListDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisListsSession()
	if err != nil {
		return err
	}

	listID, crn, _ := flex.ConvertTftoCisTwoVar(d.Id())
	sess.Crn = core.StringPtr(crn)
	sess.ListID = core.StringPtr(listID)
	opt := sess.NewDeleteCustomListOptions()
	_, response, err := sess.DeleteCustomList(opt)
	if err != nil {
		flex.FmtErrorf("Delete list failed: %v", response)
		return err
	}
	return nil
}
