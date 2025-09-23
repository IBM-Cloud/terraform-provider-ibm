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
	CISManagedLists           = "ibm_cis_managed_lists"
	CISManagedListsOutput     = "lists"
	CISManagedListName        = "name"
	CISManagedListDescription = "description"
	CISManagedListKind        = "kind"
)

func DataSourceIBMCISManagedLists() *schema.Resource {
	return &schema.Resource{
		Read: dataIBMCISManagedListsRead,
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
				ValidateFunc: validate.InvokeDataSourceValidator(
					"ibm_cis_managed_lists",
					"cis_id"),
			},
			CISManagedListsOutput: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Container for response information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						CISManagedListName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the List",
						},
						CISManagedListDescription: {
							Type:        schema.TypeString,
							Description: "Description of the List",
							Computed:    true,
						},
						CISManagedListKind: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Kind of the List",
						},
					},
				},
			},
		},
	}
}

func DataSourceIBMCISManagedListsValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true})

	IBMCISManagedListsValidator := validate.ResourceValidator{
		ResourceName: CISManagedLists,
		Schema:       validateSchema}
	return &IBMCISManagedListsValidator
}

func dataIBMCISManagedListsRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisListsSession()
	if err != nil {
		return err
	}
	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	opt := sess.NewGetManagedListsOptions()
	result, resp, err := sess.GetManagedLists(opt)

	if err != nil {
		flex.FmtErrorf("[WARN] List Managed List failed: %v\n", resp)
		return err
	}

	listsList := make([]map[string]interface{}, 0)
	for _, listObj := range result.Result {
		listOutput := map[string]interface{}{}
		listOutput[CISManagedListDescription] = *listObj.Description
		listOutput[CISManagedListKind] = *listObj.Kind
		listOutput[CISManagedListName] = *listObj.Name

		listsList = append(listsList, listOutput)

	}

	d.SetId(dataSourceCISListsManagedCheckID(d))
	d.Set(CISManagedListsOutput, listsList)
	d.Set(cisID, crn)

	return nil
}

func dataSourceCISListsManagedCheckID(d *schema.ResourceData) string {
	return "managed_lists" + ":" + d.Get(cisID).(string)
}
