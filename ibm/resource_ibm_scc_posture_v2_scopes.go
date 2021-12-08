// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/scc-go-sdk/posturemanagementv2"
)

func resourceIBMSccPostureV2Scopes() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMSccPostureV2ScopesCreate,
		ReadContext:   resourceIBMSccPostureV2ScopesRead,
		UpdateContext: resourceIBMSccPostureV2ScopesUpdate,
		DeleteContext: resourceIBMSccPostureV2ScopesDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_scc_posture_v2_scopes", "name"),
				Description:  "A unique name for your scope.",
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_scc_posture_v2_scopes", "description"),
				Description:  "A detailed description of the scope.",
			},
			"collector_ids": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "The unique IDs of the collectors that are attached to the scope.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"credential_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_scc_posture_v2_scopes", "credential_id"),
				Description:  "The unique identifier of the credential.",
			},
			"credential_type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_scc_posture_v2_scopes", "credential_type"),
				Description:  "The environment that the scope is targeted to.",
			},
		},
	}
}

func resourceIBMSccPostureV2ScopesValidator() *ResourceValidator {
	validateSchema := make([]ValidateSchema, 1)
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9-\\.,_\\s]*$`,
			MinValueLength:             1,
			MaxValueLength:             50,
		},
		ValidateSchema{
			Identifier:                 "description",
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9-\\.,_\\s]*$`,
			MinValueLength:             1,
			MaxValueLength:             255,
		},
		ValidateSchema{
			Identifier:                 "credential_id",
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9-\\.,_\\s]*$`,
			MinValueLength:             1,
			MaxValueLength:             50,
		},
		ValidateSchema{
			Identifier:                 "credential_type",
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              "aws, azure, gcp, hosted, ibm, on_premise, openstack, services",
		},
	)

	resourceValidator := ResourceValidator{ResourceName: "ibm_scc_posture_v2_scopes", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMSccPostureV2ScopesCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	postureManagementClient, err := meta.(ClientSession).PostureManagementV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createScopeOptions := &posturemanagementv2.CreateScopeOptions{}
	createScopeOptions.SetAccountID(os.Getenv("SCC_POSTURE_V2_ACCOUNT_ID"))

	createScopeOptions.SetName(d.Get("name").(string))
	createScopeOptions.SetDescription(d.Get("description").(string))
	createScopeOptions.SetCollectorIds([]string{"4188"}) //[]string{
	createScopeOptions.SetCredentialID(d.Get("credential_id").(string))
	createScopeOptions.SetCredentialType(d.Get("credential_type").(string))

	scope, response, err := postureManagementClient.CreateScopeWithContext(context, createScopeOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateScopeWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateScopeWithContext failed %s\n%s", err, response))
	}

	d.SetId(*scope.ID)

	return resourceIBMSccPostureV2ScopesRead(context, d, meta)
}

func resourceIBMSccPostureV2ScopesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	postureManagementClient, err := meta.(ClientSession).PostureManagementV2()
	if err != nil {
		return diag.FromErr(err)
	}

	listScopesOptions := &posturemanagementv2.ListScopesOptions{}
	listScopesOptions.SetAccountID(os.Getenv("SCC_POSTURE_V2_ACCOUNT_ID"))

	scopeList, response, err := postureManagementClient.ListScopesWithContext(context, listScopesOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] ListScopesWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ListScopesWithContext failed %s\n%s", err, response))
	}

	log.Printf("scopeList ID %s", *(scopeList.Scopes[0].ID))
	return nil
}

func resourceIBMSccPostureV2ScopesUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	postureManagementClient, err := meta.(ClientSession).PostureManagementV2()
	if err != nil {
		return diag.FromErr(err)
	}

	updateScopeDetailsOptions := &posturemanagementv2.UpdateScopeDetailsOptions{}
	updateScopeDetailsOptions.SetAccountID(os.Getenv("SCC_POSTURE_V2_ACCOUNT_ID"))

	hasChange := false

	updateScopeDetailsOptions.SetID(d.Id())

	if d.HasChange("name") {
		updateScopeDetailsOptions.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("description") {
		updateScopeDetailsOptions.SetDescription(d.Get("description").(string))
		hasChange = true
	}

	if hasChange {
		_, response, err := postureManagementClient.UpdateScopeDetailsWithContext(context, updateScopeDetailsOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateScopeDetailsWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateScopeDetailsWithContext failed %s\n%s", err, response))
		}
	}

	return resourceIBMSccPostureV2ScopesRead(context, d, meta)
}

func resourceIBMSccPostureV2ScopesDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	postureManagementClient, err := meta.(ClientSession).PostureManagementV2()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteScopeOptions := &posturemanagementv2.DeleteScopeOptions{}
	deleteScopeOptions.SetAccountID(os.Getenv("SCC_POSTURE_V2_ACCOUNT_ID"))

	deleteScopeOptions.SetID(d.Id())

	response, err := postureManagementClient.DeleteScopeWithContext(context, deleteScopeOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteScopeWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteScopeWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}
