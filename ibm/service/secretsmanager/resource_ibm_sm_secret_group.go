// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/secrets-manager-go-sdk/secretsmanagerv2"
)

func ResourceIbmSmSecretGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSmSecretGroupCreate,
		ReadContext:   resourceIbmSmSecretGroupRead,
		UpdateContext: resourceIbmSmSecretGroupUpdate,
		DeleteContext: resourceIbmSmSecretGroupDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_secret_group", "name"),
				Description:  "The name of your secret group.",
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_secret_group", "description"),
				Description:  "An extended description of your secret group.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date that a resource was created. The date format follows RFC 3339.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date that a resource was recently modified. The date format follows RFC 3339.",
			},
		},
	}
}

func ResourceIbmSmSecretGroupValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[A-Za-z][A-Za-z0-9]*(?:_?-?.?[A-Za-z0-9]+)*$`,
			MinValueLength:             2,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "description",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `(.*?)`,
			MinValueLength:             0,
			MaxValueLength:             1024,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_sm_secret_group", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmSmSecretGroupCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return diag.FromErr(err)
	}

	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, d)

	createSecretGroupOptions := &secretsmanagerv2.CreateSecretGroupOptions{}

	createSecretGroupOptions.SetName(d.Get("name").(string))
	if _, ok := d.GetOk("description"); ok {
		createSecretGroupOptions.SetDescription(d.Get("description").(string))
	}

	secretGroup, response, err := secretsManagerClient.CreateSecretGroupWithContext(context, createSecretGroupOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateSecretGroupWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateSecretGroupWithContext failed %s\n%s", err, response))
	}

	d.SetId(*secretGroup.ID)

	return resourceIbmSmSecretGroupRead(context, d, meta)
}

func resourceIbmSmSecretGroupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return diag.FromErr(err)
	}

	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, d)

	getSecretGroupOptions := &secretsmanagerv2.GetSecretGroupOptions{}

	getSecretGroupOptions.SetID(d.Id())

	secretGroup, response, err := secretsManagerClient.GetSecretGroupWithContext(context, getSecretGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetSecretGroupWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetSecretGroupWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("name", secretGroup.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("description", secretGroup.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}
	if err = d.Set("created_at", flex.DateTimeToString(secretGroup.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}
	if err = d.Set("updated_at", flex.DateTimeToString(secretGroup.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}

	return nil
}

func resourceIbmSmSecretGroupUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return diag.FromErr(err)
	}

	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, d)

	updateSecretGroupOptions := &secretsmanagerv2.UpdateSecretGroupOptions{}

	updateSecretGroupOptions.SetID(d.Id())

	hasChange := false

	patchVals := &secretsmanagerv2.SecretGroupPatch{}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		hasChange = true
	}
	if d.HasChange("description") {
		newDescription := d.Get("description").(string)
		patchVals.Description = &newDescription
		hasChange = true
	}

	if hasChange {
		updateSecretGroupOptions.SecretGroupPatch, _ = patchVals.AsPatch()
		_, response, err := secretsManagerClient.UpdateSecretGroupWithContext(context, updateSecretGroupOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateSecretGroupWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateSecretGroupWithContext failed %s\n%s", err, response))
		}
	}

	return resourceIbmSmSecretGroupRead(context, d, meta)
}

func resourceIbmSmSecretGroupDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return diag.FromErr(err)
	}

	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, d)

	deleteSecretGroupOptions := &secretsmanagerv2.DeleteSecretGroupOptions{}

	deleteSecretGroupOptions.SetID(d.Id())

	response, err := secretsManagerClient.DeleteSecretGroupWithContext(context, deleteSecretGroupOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteSecretGroupWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteSecretGroupWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}
